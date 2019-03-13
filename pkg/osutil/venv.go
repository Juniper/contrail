package osutil

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pkg/errors"
)

// Getenv finds the value of the latest occurence of the environment variable
func Getenv(env []string, key string) (string, bool) {
	prefix := key + "="
	for i := len(env) - 1; i >= 0; i-- {
		kv := env[i]
		if strings.HasPrefix(kv, prefix) {
			return kv[len(prefix):], true
		}
	}

	return "", false
}

// Unset returns environment with the variable unset
func Unset(env []string, key string) []string {
	res := make([]string, 0, len(env))
	prefix := key + "="
	for i := 0; i < len(env); i++ {
		kv := env[i]
		if !strings.HasPrefix(kv, prefix) {
			res = append(res, kv)
		}
	}

	return res
}

// Which resolves executable name using PATH variable in environment
func Which(cmdName string, env []string) (string, error) {
	if strings.Contains(cmdName, string(os.PathSeparator)) {
		return cmdName, nil
	}

	pathVal, ok := Getenv(env, "PATH")

	if !ok {
		return "", errors.Errorf("no PATH in env")
	}

	dirsToCheck := strings.Split(pathVal, string(os.PathListSeparator))
	for _, dir := range dirsToCheck {
		pathToCheck := path.Join(dir, cmdName)

		info, err := os.Stat(pathToCheck)

		if err != nil && !os.IsNotExist(err) {
			return "", errors.WithStack(err)
		}

		if err == nil {
			if info.Mode().IsRegular() && info.Mode().Perm()&0111 != 0 {
				return pathToCheck, nil
			}
		}
	}

	return "", errors.Errorf("unable to find executable %s in PATH", cmdName)
}

// Venv sets up the environment for the command to run it in the python virtual environment
// (akin to what `source venvdir/bin/activate` would do)
func Venv(cmd *exec.Cmd, venvDir string) (*exec.Cmd, error) {
	newEnv := Unset(cmd.Env, "PYTHONHOME")
	newEnv = append(newEnv, fmt.Sprintf("VIRTUAL_ENV=%s", venvDir))

	cmdPath, okPath := Getenv(cmd.Env, "PATH")
	var newPath string
	if okPath {
		newPath = fmt.Sprintf("%s/bin:%s", venvDir, cmdPath)
	} else {
		newPath = fmt.Sprintf("%s/bin", venvDir)
	}

	newEnv = append(newEnv, fmt.Sprintf("PATH=%s", newPath))

	whichCmdPath, err := Which(cmd.Path, newEnv)

	if err != nil {
		return nil, err
	}

	newCmd := &exec.Cmd{
		Path:        whichCmdPath,
		Args:        cmd.Args,
		Env:         newEnv,
		Dir:         cmd.Dir,
		Stdin:       cmd.Stdin,
		Stdout:      cmd.Stdout,
		Stderr:      cmd.Stderr,
		ExtraFiles:  cmd.ExtraFiles,
		SysProcAttr: cmd.SysProcAttr,
	}

	return newCmd, nil
}

// VenvCommand creates a command with the environment to run it in the python virtual environment
// (akin to what `source venvdir/bin/activate` would do)
func VenvCommand(venvDir, path string, args ...string) (*exec.Cmd, error) {
	cmd := &exec.Cmd{
		Path: path,
		Args: append([]string{path}, args...),
		Env:  os.Environ(),
	}

	return Venv(cmd, venvDir)
}
