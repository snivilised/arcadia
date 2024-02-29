package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func Path(parent, relative string) string {
	segments := strings.Split(relative, "/")
	return filepath.Join(append([]string{parent}, segments...)...)
}

func Normalise(p string) string {
	return strings.ReplaceAll(p, "/", string(filepath.Separator))
}

func Reason(name string) string {
	return fmt.Sprintf("❌ for item named: '%v'", name)
}

func JoinCwd(segments ...string) string {
	if current, err := os.Getwd(); err == nil {
		parent, _ := filepath.Split(current)
		grand := filepath.Dir(parent)
		great := filepath.Dir(grand)
		all := append([]string{great}, segments...)

		return filepath.Join(all...)
	}

	panic("could not get root path")
}

func Root() string {
	if current, err := os.Getwd(); err == nil {
		return current
	}

	panic("could not get root path")
}

func Repo(relative string) string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	bytes, err := cmd.Output()

	if err != nil {
		panic(errors.Wrap(err, "couldn't get repo root"))
	}

	segments := strings.Split(relative, "/")
	output := strings.TrimSuffix(string(bytes), "\n")
	path := []string{output}
	path = append(path, segments...)

	return filepath.Join(path...)
}

func Log() string {
	if current, err := os.Getwd(); err == nil {
		parent, _ := filepath.Split(current)
		grand := filepath.Dir(parent)
		great := filepath.Dir(grand)

		return filepath.Join(great, "Test", "test.log")
	}

	panic("could not get root path")
}
