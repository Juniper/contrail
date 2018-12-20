package agent

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/flosch/pongo2"
	"github.com/joho/godotenv"
	shellwords "github.com/mattn/go-shellwords"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type handler map[string]interface{}
type taskHandler func(handler handler, task *task, resource map[string]interface{}) (interface{}, error)

var globalTaskHandler = map[string]taskHandler{}

func registerHandler(id string, f taskHandler) {
	globalTaskHandler[id] = f
}

func init() {
	registerHandler("command", commandHandler)
	registerHandler("save", saveHandler)
	registerHandler("remove", removeHandler)
	registerHandler("template", templateHandler)
	registerHandler("vars", varsHandler)
	registerHandler("env_file", envHandler)
	registerHandler("debug", debugHandler)
}

// nolint: gocyclo
func commandHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	c, err := getCommand(handler, context)
	if err != nil {
		return nil, err
	}

	args, err := shellwords.Parse(c)
	if err != nil {
		return nil, err
	}

	var cmd *exec.Cmd
	switch len(args) {
	case 0:
		return nil, fmt.Errorf("command format error")
	case 1:
		cmd = exec.Command(args[0])
	default:
		cmd = exec.Command(args[0], args[1:]...)
	}

	chdir := ""
	commandArgs, ok := handler["args"].(map[interface{}]interface{})
	if ok {
		chdir, _ = applyTemplate(commandArgs["chdir"], context)
	}

	if chdir != "" {
		cmd.Dir = chdir
	}

	// Pass Env stored in the context to the Command
	var env []string
	if context["env"] != nil {
		for k, v := range context["env"].(map[string]string) {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
	}
	cmd.Env = env

	var output bytes.Buffer
	stdout, _ := cmd.StdoutPipe() // nolint: errcheck
	stderr, _ := cmd.StderrPipe() // nolint: errcheck
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	stdoutScanner := bufio.NewScanner(stdout)
	for stdoutScanner.Scan() {
		m := stdoutScanner.Text()
		output.WriteString(m)
		log.Debug(m)
	}

	stderrScanner := bufio.NewScanner(stderr)
	for stderrScanner.Scan() {
		m := stderrScanner.Text()
		output.WriteString(m)
		log.Error(m)
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func getCommand(h handler, context map[string]interface{}) (string, error) {
	command, err := applyTemplate(h["command"], context)
	if err != nil {
		return "", err
	}
	if command == "" {
		return "", fmt.Errorf("empty command")
	}

	return command, nil
}

func saveHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	saveConf, ok := handler["save"].(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("template format issue")
	}
	outputPath, err := applyTemplate(saveConf["dest"], context)
	if err != nil {
		return nil, err
	}
	format, err := applyTemplate(saveConf["format"], context)
	if err != nil {
		return nil, err
	}
	var resourceBytes []byte
	resource := context["resource"]
	resourceMap, _ := resource.(map[string]interface{})
	outputData := map[string]interface{}{
		resourceMap["schema_id"].(string): map[string]interface{}{
			resourceMap["uuid"].(string): resourceMap,
		},
	}
	if format == "json" {
		resourceBytes, _ = json.Marshal(outputData)
	} else {
		resourceBytes, _ = yaml.Marshal(outputData)
	}
	return nil, task.agent.backend.write(outputPath, resourceBytes)
}

func removeHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	outputPath, err := applyTemplate(handler["remove"], context)
	if err != nil {
		return nil, err
	}
	return nil, task.agent.backend.remove(outputPath)
}

func templateHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	templateConf, ok := handler["template"].(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("template format issue")
	}
	templateSrc, err := applyTemplate(templateConf["src"], context)
	if err != nil {
		return nil, nil
	}
	outputPath, err := applyTemplate(templateConf["dest"], context)
	if err != nil {
		return nil, nil
	}
	template, err := pongo2.FromFile(templateSrc)
	if err != nil {
		return nil, nil
	}
	output, err := template.ExecuteBytes(context)
	if err != nil {
		return nil, nil
	}
	return nil, task.agent.backend.write(outputPath, output)
}

func applyTemplateObject(template interface{}, context map[string]interface{}) (interface{}, error) {
	var err error
	switch t := template.(type) {
	case string:
		return applyTemplate(t, context)
	case []interface{}:
		result := []interface{}{}
		for _, item := range t {
			output, aErr := applyTemplateObject(item, context)
			if aErr != nil {
				return nil, aErr
			}
			result = append(result, output)
		}
		return result, nil
	case map[interface{}]interface{}:
		result := map[string]interface{}{}
		for key, value := range t {
			result[key.(string)], err = applyTemplateObject(value, context)
			if err != nil {
				return nil, err
			}
		}
		return result, nil
	}
	return nil, nil
}

func varsHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	vars, ok := handler["vars"].(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid vars format")
	}

	_, err := applyTemplateObject(vars, context)
	if err != nil {
		return nil, err
	}

	for key, value := range vars {
		context[key.(string)], _ = applyTemplateObject(value, context)
	}
	return nil, nil
}

func envHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	envFile, err := applyTemplate(handler["env_file"], context)
	if err != nil {
		return nil, err
	}

	readEnv, err := godotenv.Read(envFile)
	if err != nil {
		return nil, fmt.Errorf("cannot Read Env File:%s", envFile)
	}

	context["env"] = readEnv
	return nil, nil
}

func debugHandler(handler handler, task *task, context map[string]interface{}) (interface{}, error) {
	debugLog, err := applyTemplate(handler["debug"], context)
	if err != nil {
		return nil, err
	}
	log.Debug(debugLog)
	return nil, nil
}
