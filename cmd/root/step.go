package root

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

// Run implements the command
func (o *options) Run() error {

	homedir, err  := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "cannot get users HOME directory")
	}
	token, err := o.getToken(filepath.Join(homedir, ".git-credentials"))
	if err != nil {
		return errors.Wrap(err, "failed to get token")
	}
	fmt.Println(token)
	return nil
}

func (o *options) getToken(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("cannot open file %s", filename))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains( scanner.Text(), "disabled") {
			continue
		}
		if len(scanner.Text()) > 0 {
			lines = append(lines, scanner.Text())

		}
	}

	if len(lines) > 1 {
		return "", fmt.Errorf("multiple credentials found in %s", filename)
	}

	for _, line := range lines {
		parts := strings.Split(line, ":")
		token := strings.Split(parts[2], "@")

		return token[0], nil
	}
	if err := scanner.Err(); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("failed to scan file %s", filename))
	}
	return "", fmt.Errorf("no token found in file %s", filename)
}