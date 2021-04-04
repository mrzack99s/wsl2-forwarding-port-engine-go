package runtimes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func Parse() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(usr.HomeDir + "\\.wfp-engine\\.wfp-routines.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &RuleTables)
}

func WriteToFile() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	file, _ := json.MarshalIndent(RuleTables, "", " ")
	_ = ioutil.WriteFile(usr.HomeDir+"\\.wfp-engine\\.wfp-routines.json", file, 0655)
}
