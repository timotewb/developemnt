package main

import (
	"encoding/json"
	"fmt"
	"log"
	m "systemsim/models"
)

func main(){
	data := m.SimType{
		Inputs: []m.InputType{
			{
				Name: "Resource",
				Attributes: []m.InputAttribute{
					{
						ID: 1,
						Description: "Initial Resource count",
						Value: 10.0,
					},
					{
						ID: 2,
						Description: "Max utilisation",
						Value: 80.0,
					},
				},
			},
			{
				Name: "Engagements",
				Attributes: []m.InputAttribute{
					{
						ID: 1,
						Description: "Initial Engagement count",
						Value: 5.0,
					},
					{
						ID: 2,
						Description: "Resource per engagement",
						Value: 2.5,
					},
					{
						ID: 3,
						Description: "Engagement value",
						Value: 100.0,
					},
				},
			},
		},
	}

    // Pretty print josn with MarshalIndent
    dataJSON, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("%s\n", string(dataJSON))
}