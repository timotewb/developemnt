package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	l "systemsim/lib"
	m "systemsim/models"
)

func main(){
	file, _ := os.ReadFile("db.json")
	data := m.SimType{}

	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println("--- CalcLinks() ---")
	l.CalcLinks(data)

	fmt.Println("--- Pretty Print ---")
    // Pretty print josn with MarshalIndent
    dataJSON, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("%s\n", string(dataJSON))
}

// add | around ids in equation