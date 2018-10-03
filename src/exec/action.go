package exec

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func action(c *cli.Context) error {
	cmd := exec.Command("sh", "-c", c.String("command"))
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
