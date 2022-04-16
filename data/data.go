package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadData() *DungeonData {
	// Open our jsonFile
	jsonFile, err := os.Open("./data/dungeon.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dungeon DungeonData

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &dungeon)
	return &dungeon
}
