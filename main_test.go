package main

import (
	"./api"
	"testing"
)

func TestGetLights(t *testing.T) {

	l := api.Lights{}
	x := l.GetAllLightObjects()
	if x == nil {
		t.Error()
	}
}

func TestPrintTable(t *testing.T) {
	l := api.Lights{}
	l.PrintLightTable()
}

func TestTurnOn(t *testing.T) {
	new(api.Lights).PrintLightTable()
	newLight := api.Light{}
	newLight = newLight.GetLight(1)
	newLight.TurnOn()
	new(api.Lights).PrintLightTable()

}

func TestTurnOff(t *testing.T) {
	new(api.Lights).PrintLightTable()
	newLight := api.Light{}
	newLight = newLight.GetLight(1)
	newLight.TurnOff()
	new(api.Lights).PrintLightTable()

}
