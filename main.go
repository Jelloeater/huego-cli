package main

import (
	"os"
	"strconv"

	"./api"
	"./web"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	log.SetReportCaller(true)

	app := cli.NewApp()
	app.Name = "HueGo"
	app.Usage = ""
	app.HideVersion = true
	app.HideHelp = false
	app.EnableBashCompletion = true

	// Commands to be run go here, after parsing variables
	app.Commands = []*cli.Command{
		{
			UseShortOptionHandling: true,
			Name:                   "web",
			Aliases:                []string{"w"},
			Usage:                  "start web GUI",
			Action: func(c *cli.Context) error {
				web.StartServer()
				return nil
			},
		},
		{
			UseShortOptionHandling: true,
			Name:                   "list",
			Aliases:                []string{"l"},
			Usage:                  "print list of lights",
			Action: func(c *cli.Context) error {
				api.Lights{}.PrintLightTable()
				return nil
			},
		},
		{
			UseShortOptionHandling: true,
			Name:                   "turn_on",
			Aliases:                []string{"on"},
			Usage:                  "turn on a light",
			Action: func(c *cli.Context) error {
				arg, err := strconv.Atoi(c.Args().First()) // Converts first arg from string to int
				if err != nil {
					log.Error("Invalid Input")
					log.Fatal(err)
				}
				l := new(api.Light).GetLight(arg) // Create new light object
				l.TurnOn()
				return nil
			},
		},
		{
			UseShortOptionHandling: true,
			Name:                   "turn_off",
			Aliases:                []string{"off"},
			Usage:                  "turn off a light",
			Action: func(c *cli.Context) error {
				arg, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					log.Error("Invalid Input")
					log.Fatal(err)
				}
				x := api.Light{}
				x = x.GetLight(arg)
				x.TurnOff()
				return nil
			},
		},
	}

	// Parse Commands and flags here, order of commands matters "-d l" != "l -d"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("EOP")
}
