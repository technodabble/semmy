package app

import (
	"os"

	"github.com/technodabble/semmy/commands"
	"github.com/urfave/cli"
)

type Semmy struct {
	a *cli.App
}

func Create() *Semmy {
	sem := &Semmy{
		a: cli.NewApp(),
	}
	sem.a.Usage = "A Command-Line Semantic Versioning Toolbox"
  sem.a.UsageText = "semmy range [version]"
  sem.a.ArgsUsage = "range has a specific format see: https://github.com/blang/semver#ranges"
	sem.a.Version = "0.1.0"
	sem.a.Action = commands.Check
	return sem
}

func (s Semmy) Run() {
	s.a.Run(os.Args)
}
