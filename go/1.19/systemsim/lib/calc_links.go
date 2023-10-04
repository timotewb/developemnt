package lib

import (
	"go/token"
	"go/types"
	"strconv"
	"strings"
	m "systemsim/models"
)

func CalcLinks(s m.SimType){

	// for each link
	var e string
	var j int
	for i := 0; i < len(s.Link); i++ {

		// create executable equation
		e = s.Link[i].Equation
		// for each item go find it and replace the id in the equation with value
		for j = 0; j <len(s.Link[i].Items); j++{
			e = strings.Replace(e, "|"+strconv.Itoa(s.Link[i].Items[j].ID)+"|", strconv.FormatFloat(getInputAttributeValue(s, s.Link[i].Items[j].ItemID, s.Link[i].Items[j].ValueID), 'g', 5, 64), 1)
		}
		s.Link[i].EquationEvaluated = e

		// evaluate equation
		fs := token.NewFileSet()
		tv, err := types.Eval(fs, nil, token.NoPos, e)
		if err != nil {
			panic(err)
		}

		// udpate output
		s.Link[i].Output = tv.Value.String()
	}
}

func getInputAttributeValue(s m.SimType ,inputID int, attributeID int,) (float64){

	var r float64

	// loop over each of the inputs
	for i := 0; i < len(s.Inputs); i++ {
		if s.Inputs[i].ID == inputID{
			// loop over each of the attributes
			for j := 0; j < len(s.Inputs); j++ {
				if s.Inputs[i].Attributes[j].ID == attributeID{
					r = float64(s.Inputs[i].Attributes[j].Value)
				}
			}
		}
	}
	return r
}