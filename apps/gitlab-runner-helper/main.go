package main

import (
	"os"
	"path"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/gitlab-org/gitlab-runner/common"
	"gitlab.com/gitlab-org/gitlab-runner/helpers/cli"
	"gitlab.com/gitlab-org/gitlab-runner/helpers/formatter"

	_ "gitlab.com/gitlab-org/gitlab-runner/commands/helpers"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			// log panics forces exit
			if _, ok := r.(*logrus.Entry); ok {
				os.Exit(1)
			}
			panic(r)
		}
	}()

	formatter.SetRunnerFormatter()
	cli_helpers.AddSecretsCleanupLogHook()

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "an Alloy Runner Helper"
	app.Version = common.AppVersion.ShortLine()
	cli.VersionPrinter = common.AppVersion.Printer
	app.Authors = []cli.Author{
		{
			Name:  "Patricio Cano",
			Email: "admin@alloy-ci.com",
		},
	}
	cli_helpers.SetupLogLevelOptions(app)
	app.Commands = common.GetCommands()
	app.CommandNotFound = func(context *cli.Context, command string) {
		logrus.Fatalln("Command", command, "not found")
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
