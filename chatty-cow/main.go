package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "chatty-cow"
	app.Version = VERSION
	app.Usage = "Talks to other chatty cows and reports the health of its conversations."
	app.Action = func(c *cli.Context) error {
		logrus.Info("I'm a turkey")
		return nil
	}

	app.Run(os.Args)
}
