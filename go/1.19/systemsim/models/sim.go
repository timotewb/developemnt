package models

//----------------------------------------------------------------------------------------
// Sim
//----------------------------------------------------------------------------------------
type SimType struct{
	Inputs []InputType `json:"inputs"`
	Link []LinkType `json:"links"`
}
