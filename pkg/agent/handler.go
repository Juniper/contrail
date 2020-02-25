package agent

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"

	"github.com/flosch/pongo2"
	"github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// TODO(dji): determin the output type
func commandHandler(command string) (string, error) {
	args, err := shellwords.Parse(command)
	if err != nil {
		return "", errors.Wrap(err, "parse command arguments failed")
	}

	var cmd *exec.Cmd
	switch len(args) {
	case 0:
		return "", fmt.Errorf("command format error")
	case 1:
		cmd = exec.Command(args[0])
	default:
		cmd = exec.Command(args[0], args[1:]...)
	}

	var output bytes.Buffer
	stdout, _ := cmd.StdoutPipe() // nolint: errcheck
	stderr, _ := cmd.StderrPipe() // nolint: errcheck
	err = cmd.Start()
	if err != nil {
		return "", errors.Wrap(err, "cmd start failed")
	}
	stdoutScanner := bufio.NewScanner(stdout)
	for stdoutScanner.Scan() {
		m := stdoutScanner.Text()
		output.WriteString(m)
		logrus.Debug(m)
	}

	stderrScanner := bufio.NewScanner(stderr)
	for stderrScanner.Scan() {
		m := stderrScanner.Text()
		output.WriteString(m)
		logrus.Error(m)
	}

	err = cmd.Wait()
	if err != nil {
		return "", errors.Wrap(err, "wait for the command to exit failed")
	}
	return output.String(), nil
}

func templateHandler(templateSrc, outputPath string, context map[string]interface{}, a *Agent) error {
	template, err := pongo2.FromFile(templateSrc)
	if err != nil {
		return errors.Wrap(err, "pongo2 template generation failed")
	}

	output, err := template.ExecuteBytes(context)
	if err != nil {
		return errors.Wrap(err, "execute template and render it as []byte failed")
	}

	return a.backend.write(outputPath, output)
}
