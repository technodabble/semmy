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
  sem.a.UsageText = "semmy [global-options] command [options] [version]"
	sem.a.Version = "0.2.0"
  sem.a.Commands = []cli.Command{
    commands.Check(),
  }
	return sem
}

func (s Semmy) Run() {
	s.a.Run(os.Args)
}
