package api

import (
	"../settings"
	"encoding/json"
	"github.com/Jeffail/gabs"
	"github.com/jedib0t/go-pretty/table"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Lights struct {
	Light
}

type Light struct {
	id    int
	name  string
	state bool
}

//NewLight Constructor for new light objects
func (l *Light) NewLight(Id_In int, Name_In string, State_In bool) Light {
	newLight := new(Light)
	newLight.id = Id_In
	newLight.name = Name_In
	newLight.state = State_In
	return *newLight
}

//GetLight Loads a light object with data
func (l *Light) GetLight(Id_in int) Light {
	singleLightJson, _ := gabs.ParseJSON(ApiHelpers{}.GetApiSingleLightJSON(Id_in)) // Pulls in API JSON
	log.Println(singleLightJson.String())

	newLight := new(Light)
	newLight.id = Id_in
	newLight.name = singleLightJson.Search("name").String()
	newLight.state, _ = strconv.ParseBool(singleLightJson.Search("state").Search("on").String())

	return *newLight
}

func (l *Light) TurnOn() *grequests.Response {
	in, _ := json.Marshal(map[string]bool{"on": true}) // Create a map for the body submission
	ro := grequests.RequestOptions{JSON: in}
	reqURL := settings.Base_url + "/lights/" + strconv.Itoa(l.id) + "/state"
	resp, err := grequests.Put(reqURL, &ro)
	// API JSON body parser *IS* case sensitive, you cannot marshal structs, as
	// their values need to be uppercase to be public, and this causes an issue with the Hue
	// API requests ... -_-

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Info(resp.String())
	return resp
}

func (l *Light) TurnOff() *grequests.Response {
	in, _ := json.Marshal(map[string]bool{"on": false})
	ro := grequests.RequestOptions{JSON: in}
	reqURL := settings.Base_url + "/lights/" + strconv.Itoa(l.id) + "/state"
	resp, err := grequests.Put(reqURL, &ro)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	log.Info(resp.String())
	return resp
}

func (l *Lights) GetListOfLights() []Light {
	var LightObjList []Light
	jsonParsed, _ := gabs.ParseJSON(ApiHelpers{}.GetApiJSON()) // Pulls in API JSON
	lightListJSON, _ := jsonParsed.Search("lights").Children() // Searches JSON tree for lights array

	for index, single_light := range lightListJSON {
		log.Println(single_light.String())
		nameJson := single_light.Search("name").String()
		idJson := index + 1
		stateJson, _ := strconv.ParseBool(single_light.Search("state").Search("on").String())

		singleLightObj := l.NewLight(idJson, nameJson, stateJson)
		LightObjList = append(LightObjList, singleLightObj)
	}
	return LightObjList
}

func (l *Lights) generateLightTable() {

	light_list := l.GetListOfLights()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "state"})

	for _, i := range light_list {
		t.AppendRow(table.Row{i.id, i.name, i.state})
	}
	t.Render()
}

func (Lights) PrintLightTable() {
	lightObj := Lights{}
	lightObj.generateLightTable()
}

type ApiHelpers struct {
}

func (ApiHelpers) GetApiJSON() []byte {
	resp, err := grequests.Get(settings.Base_url+"/", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp.Bytes()
}

func (ApiHelpers) GetApiSingleLightJSON(id int) []byte {
	reqURL := settings.Base_url + "/lights/" + strconv.Itoa(id)
	resp, err := grequests.Get(reqURL, nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp.Bytes()
}
