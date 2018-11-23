package main

import (
	"./api"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"os"
)

func main() {
	log.SetReportCaller(true)

	app := cli.NewApp()
	app.Name = "HueGo"
	app.Usage = ""
	app.HideVersion = true
	app.HideHelp = false
	app.EnableBashCompletion = true

	// Setup flags here
	var DebugMode bool
	flags := []cli.Flag{
		cli.BoolFlag{

			Name:        "debug, d",
			Usage:       "enable debug mode",
			Destination: &DebugMode,
		},
	}

	// Commands to be run go here, after parsing variables
	app.Commands = []cli.Command{
		{
			UseShortOptionHandling: true,
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "print list of lights",
			Action: func(c *cli.Context) error {
				api.Lights{}.PrintLightTable()
				return nil
			},
		},
	}

	app.Flags = flags // Assign flags via parse right before we start work
	app.Before = func(c *cli.Context) error {
		// Actions to run before running parsed commands
		if DebugMode {
			log.SetLevel(5)
			log.Info("Debug Mode")
		} else {
			log.SetLevel(3)
			log.Warn("Normal Mode")
		}
		return nil
	}
	// Parse Commands and flags here, order of commands matters "-d l" != "l -d"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("EOP")
}
