package config

import (
	"encoding/json"
	"io/ioutil"
	"fmt"

	simpleJson "github.com/bitly/go-simplejson"
)

func Load(path string, cfg interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(cfg)
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)
	return json.Unmarshal(content, cfg)
}

func JsonMapLoad(fileName string) *simpleJson.Json {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	js, err := simpleJson.NewJson(content)
	if err != nil {
		panic(err)
	}
	return js
}
