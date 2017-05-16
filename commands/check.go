package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/urfave/cli"
)

func Check() cli.Command {
  return cli.Command {
    Name: "check",
    Aliases: []string{"c"},
    ArgsUsage: "range [version]",
    Usage: "Checks if the version is in the given range",
    Description: `Checks if the version is in the given range.
      range - As described here: https://github.com/blang/semver#ranges
      version - The version string. If omitted, it will be read from STDIN.
      returns zero if true and non-zero otherwise.`,
    Action: check,
  }
}

func check(c *cli.Context) error {
  var version string
  if c.NArg() > 2 || c.NArg() < 1 {
    cli.ShowAppHelp(c)
  }

  target := c.Args().Get(0)
	r, err := semver.ParseRange(target)
	if err != nil {
    return cli.NewExitError(fmt.Sprintf("Error Parsing Range: %v\n%v\n", target, err), 1)
	}

  if c.NArg() > 1 {
    version = c.Args().Get(1)
  } else {
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
      version = scanner.Text()
    }
    if err := scanner.Err(); err != nil {
      return cli.NewExitError(fmt.Sprintf("Error Reading From Stdin! %v", err), 1)
    }
  }

	v, err := semver.Parse(version)
	if err != nil {
    return cli.NewExitError(fmt.Sprintf("Error Parsing Version: %v\n%v\n", version, err), 1)
	}

	if r(v) {
    fmt.Printf("%v is in the range: %v\n", v, target)
    return nil
	} else {
    fmt.Printf("%v is NOT in the range: %v\n", v, target)
    return cli.NewExitError("", 1)
  }
}
