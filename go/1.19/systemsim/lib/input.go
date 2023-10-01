package lib

import (
	m "systemsim/models"
)

func Input() (m.InputType){
	_ = m.SimType{
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
	var r m.InputType
	r.Name = "This is an input"
	return r
}