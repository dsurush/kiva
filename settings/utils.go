package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// AppSettings app settnigs r.
	ReqURL RequestsURL
)

func ReadSettings(filepath string) RequestsURL {
	var ReqURL RequestsURL
	doc, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Printf("err is = %e\n", err)
	}

	err = json.Unmarshal(doc, &ReqURL)
	if err != nil {
		//		panic(err)
		fmt.Printf("err is = %e\n", err)
	}

	return ReqURL
}