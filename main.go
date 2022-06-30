package main

import (
	"SystemYaml/racqueteers"
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {

	sysFile, err := ioutil.ReadFile("system.yaml")

	if err != nil {

		log.Fatal(err)
	}

	data := &racqueteers.TeamD{}

	err2 := yaml.Unmarshal(sysFile, data)

	if err2 != nil {

		log.Fatal(err2)
	}

	fmt.Printf("%+v \n", data)

}
