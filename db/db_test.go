package db

import (
	"testing"
)

func TestFindAndSave(t *testing.T){
	result:= FindTiddlyByName("adamatti")
	if result != nil {
		t.Errorf("First time tiddler adamatti shouldn't exist")
		return
	}

	SaveTiddly("adamatti", []byte("something"))
	result = FindTiddlyByName("adamatti")
	if result == nil {
		t.Errorf("Second time tiddler adamatti should exist")
	}
}
