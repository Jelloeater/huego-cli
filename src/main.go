package main

import (
	"./settings"
	"github.com/Jeffail/gabs"
	"github.com/jedib0t/go-pretty/table"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"os"
	"strconv"
)

type Lights struct {
	Light
}

type Light struct {
	Name  string
	State bool
}

//NewLight Constructor for new light objects
func (l *Light) NewLight(NameIn string, State_In bool) Light {
	m := new(Light)
	m.Name = NameIn
	m.State = State_In
	return *m
}

func (l *Lights) GetListOfLights() []Light {
	var LightObjList []Light
	jsonParsed, _ := gabs.ParseJSON(ApiHelpers{}.GetLightsRawJSON_bytes()) // Pulls in API JSON
	lightListJSON, _ := jsonParsed.Search("lights").Children()             // Searches JSON tree for lights array

	for _, single_light := range lightListJSON {
		log.Println(single_light.String())
		nameJson := single_light.Search("name").String()
		stateJson, _ := strconv.ParseBool(single_light.Search("state").Search("on").String())

		singleLightObj := l.NewLight(nameJson, stateJson)
		LightObjList = append(LightObjList, singleLightObj)
	}
	return LightObjList
}

func (l *Lights) PrintListOfLights() {

	light_list := l.GetListOfLights()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "State"})

	for _, i := range light_list {
		t.AppendRow(table.Row{i.Name, i.State})
	}
	t.Render()
}
func (Lights) PrintLightLight() {
	lightObj := Lights{}
	lightObj.PrintListOfLights()
}

type ApiHelpers struct {
}

func (ApiHelpers) GetLightsRawJSON_bytes() []byte {
	resp, err := grequests.Get(settings.Base_url+"/", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp.Bytes()
}

func (a *ApiHelpers) GetLightsRawJSON_String() string {
	return string(a.GetLightsRawJSON_bytes())
}

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
				Lights{}.PrintLightLight()
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
