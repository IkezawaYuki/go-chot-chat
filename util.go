package main

import (
	"encoding/json"
	"os"
)

var config Configuration

type Configuration struct {
	Address string
	ReadTimeout int64
	WriteTimeout int64
	Static string
}

func init(){
	loadConfig()
}


func loadConfig(){
	files, err := os.Open("config.json")
	if err != nil{
		panic(err)
	}

	decorder := json.NewDecoder(files)
	config = Configuration{}
	err = decorder.Decode(&config)
	if err != nil{
		panic(err)
	}
}