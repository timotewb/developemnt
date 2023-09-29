package lib

import (
	m "systemsim/models"
)

func Input() (m.InputType){
	var r m.InputType
	r.Name = "This is an input"
	return r
}