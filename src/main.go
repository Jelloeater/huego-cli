package main



import (
	"github.com/levigross/grequests"
	"log"
	"./settings"
	"github.com/Jeffail/gabs"
)


type LightHelpers struct {
}

func (LightHelpers)GetLightsRawJSON_bytes() []byte{
	resp, err := grequests.Get(settings.Base_url+"/", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp.Bytes()
}

func (LightHelpers)GetLightsRawJSON_String()string{
	return string(LightHelpers{}.GetLightsRawJSON_bytes())
}

func main() {
	log.Println(LightHelpers{}.GetLightsRawJSON_String())

	jsonParsed, _ := gabs.ParseJSON(LightHelpers{}.GetLightsRawJSON_bytes())

	//log.Println(jsonParsed.Path("lights.1.state.on").String())
	//log.Println(jsonParsed.Path("lights.1.name").String())

	//x, err := jsonParsed.Path("lights").ArrayCount()
	//_ = err
	//log.Println(x)
	//log.Println(jsonParsed.Path("lights").String())

	//lights2 := jsonParsed.Search("lights").Data()
	//
	//println(lights2)

	// S is shorthand for Search
	light_list, _ := jsonParsed.Search("lights").Children()
	for _, i := range light_list {
		log.Println(i.String())
		x:= i.Search("name")
		y := i.Search("state").Search("on")


		log.Println(x)
		log.Println(y)

	}


}