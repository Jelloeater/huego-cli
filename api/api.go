package api

import (
	"../settings"
	"encoding/json"
	"github.com/Jeffail/gabs"
	"github.com/jedib0t/go-pretty/table"
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Lights struct {
	Light
}

type Light struct {
	id    int
	name  string
	state bool
}

func (l *Light) Id() int {
	return l.id
}

func (l *Light) Name() string {
	return l.name
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
	log.Info(singleLightJson.String())
	if strings.Contains(singleLightJson.String(), "not available") { // Check to make sure light actually exists
		log.Fatal("Light not found ID= " + strconv.Itoa(Id_in))
	}

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

func (l *Lights) GetAllLightObjects() []Light {
	var LightObjList []Light
	jsonParsed, _ := gabs.ParseJSON(ApiHelpers{}.GetApiJSON()) // Pulls in API JSON
	lightListMap, _ := jsonParsed.Search("lights").ChildrenMap()
	// Searches JSON tree for lights object list and maps them to "string":object pairs

	log.Info(lightListMap)
	for JsonObjectName, singleLightObject := range lightListMap {
		log.Println(singleLightObject.String())
		nameJson := singleLightObject.Search("name").String()
		idJson, _ := strconv.Atoi(JsonObjectName)
		stateJson, _ := strconv.ParseBool(singleLightObject.Search("state").Search("on").String())

		singleLightObj := l.NewLight(idJson, nameJson, stateJson)
		LightObjList = append(LightObjList, singleLightObj)
	}
	return LightObjList
}

func (l *Lights) GenerateSortedLightList() []Light {
	light_list := l.GetAllLightObjects()
	sort.Slice(light_list, func(i, j int) bool { return light_list[i].id < light_list[j].id })
	return light_list
}

func (l *Lights) generateLightTable() table.Writer {
	light_list := l.GenerateSortedLightList()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "state"})

	for _, i := range light_list {
		t.AppendRow(table.Row{i.id, i.name, i.state})
	}
	return t
}

func (l *Lights) GenerateLightTableMarkdown() string {
	return l.generateLightTable().RenderMarkdown()
}

func (l *Lights) GenerateLightTableHTML() string {
	return l.generateLightTable().RenderHTML()
}

func (Lights) PrintLightTable() {
	lightObj := Lights{}
	lightObj.generateLightTable().Render()
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
