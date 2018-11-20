package main



import (
	"github.com/levigross/grequests"
	"log"
	"./settings"
	"github.com/Jeffail/gabs"
	"fmt"
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

	//jsonParsed, _ := gabs.ParseJSON([]byte(`{"array":[ {"value":1}, {"value":2}, {"value":3} ]}`))
	fmt.Println(jsonParsed.Path("lights.1.state.on").String())
	fmt.Println(jsonParsed.Path("lights.1.name").String())

	x, err := jsonParsed.Path("lights").ArrayCount()
	_ = err
	fmt.Println(x)
	fmt.Println(jsonParsed.Path("lights").String())

	lights2 := jsonParsed.Search("lights").Data()

	println(lights2)
	//children, _ := jsonParsed.S("lights").Children()

	//	for _, child := range children {
	//		fmt.Println(child.Data())
	//		child.ArrayCount()
	//	}
	//}


}