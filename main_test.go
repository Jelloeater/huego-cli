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
