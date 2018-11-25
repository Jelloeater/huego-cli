# HueGo
**A CLI application to interact with Philips Hue API**

![](docs/huego-logo.jpg)

## Requirments
* Go Ver = 10.x

**Libraries needed**

* github.com/levigross/grequests
* github.com/Jeffail/gabs
* github.com/jedib0t/go-pretty/table
* github.com/levigross/grequests
* github.com/sirupsen/logrujess

**Settings**

You will need to define your settings in `/settings/setting.go`

Just rename the `settings-example.go` file and fill in your info

There are a few guides online on how to generate a local API key

Ex https://developers.meethue.com/documentation/getting-started

## Usage
* To list lights
  * `hugocli.exe l`
* To turn light #1 on
  * `hugocli.exe on 1`
* To turn light #2 off
  * `hugocli.exe off 2`