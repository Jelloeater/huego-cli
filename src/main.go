package main



import (
	"github.com/levigross/grequests"
	log "github.com/sirupsen/logrus"
	"./settings"
	"github.com/Jeffail/gabs"
	"github.com/jedib0t/go-pretty/table"
	"strconv"
	"os"
)

type Light struct {
	Name string
	State bool
}

//NewLight Constructor for new light objects
func (Light) NewLight(NameIn string, State_In bool) Light {
	m := new(Light)
	m.Name = NameIn
	m.State = State_In
	return *m
}

func (Light) GetListOfLights()[]Light{
	var LightObjList []Light
	jsonParsed, _ := gabs.ParseJSON(ApiHelpers{}.GetLightsRawJSON_bytes()) // Pulls in API JSON
	lightListJSON, _ := jsonParsed.Search("lights").Children()             // Searches JSON tree for lights array

	for _, single_light := range lightListJSON {
		log.Println(single_light.String())
		nameJson := single_light.Search("name").String()
		stateJson, _ := strconv.ParseBool(single_light.Search("state").Search("on").String())

		singleLightObj := Light{}.NewLight(nameJson, stateJson)
		LightObjList = append(LightObjList, singleLightObj)
	}
	return LightObjList
}

func (Light) PrintListOfLights(){

	light_list := Light{}.GetListOfLights()
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "State"})

	for _, i := range light_list{
		t.AppendRow(table.Row{i.Name,i.State})
	}
	t.Render()
}



type ApiHelpers struct {

}


func (ApiHelpers)GetLightsRawJSON_bytes() []byte{
	resp, err := grequests.Get(settings.Base_url+"/", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp.Bytes()
}

func (ApiHelpers)GetLightsRawJSON_String()string{
	return string(ApiHelpers{}.GetLightsRawJSON_bytes())
}

func main() {
	log.SetReportCaller(true)

	Light{}.PrintListOfLights()
}