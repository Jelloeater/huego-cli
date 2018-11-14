package src



import (
	"github.com/levigross/grequests"
	"log"
	"./settings"
)


type Light struct {

}

func (Light)GetLightsRawJSON()string{
	resp, err := grequests.Get(settings.Base_url+"/lights", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}
	return resp.String()
}
func main() {
	out := Light{}.GetLightsRawJSON()
	println(out)
}
