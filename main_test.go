package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	key := Request{"khchan"}
	result, err := Handle(key, nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	wanted := "Hello khchan!"
	if result != wanted {
		t.Errorf("wanted %s, actual %s", wanted, result)
	}

}
