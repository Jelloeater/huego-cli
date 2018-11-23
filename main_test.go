package main

import (
	"./api"
	"testing"
)

func TestTable(t *testing.T) {

	l := api.Lights{}
	x := l.GetListOfLights()
	if x == nil {
		t.Error()
	}
}

func TestPrintTable(t *testing.T) {
	l := api.Lights{}
	l.PrintLightTable()
}

func TestTurnOn(t *testing.T) {
	newLight := api.Light{}
	newLight = newLight.GetLight(1)
	newLight.TurnOn()

}

func TestTurnOff(t *testing.T) {
	newLight := api.Light{}
	newLight = newLight.GetLight(1)
	newLight.TurnOff()

}
