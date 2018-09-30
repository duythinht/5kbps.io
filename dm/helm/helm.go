package helm

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"go.5kbps.io/dm/vault"
)

const command = "helm"

func run(stdin io.Reader, args ...string) error {
	fmt.Println(args)
	cmd := exec.Command(command, args...)
	if stdin != nil {
		cmd.Stdin = stdin
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func DeployEncryptedValues(fp string, k string, args ...string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	values, err := vault.DecryptFile(f, k)
	if err != nil {
		return err
	}
	return run(strings.NewReader(values), append(args, "-f", "-")...)
}

func Run(args ...string) error {
	return run(nil, args...)
}
