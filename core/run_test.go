package core_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMashall(t *testing.T) {

	mp := map[string]string{
		"RootPath": "./storage-1",
	}

	res, err := json.Marshal(mp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(res))
}
