package models

//----------------------------------------------------------------------------------------
// Input
//----------------------------------------------------------------------------------------
type InputType struct{
	Name string `json:"name"`
	Attributes []InputAttribute `json:"attributes"`
}

type InputAttribute struct{
	ID int	`json:"id"`
	Value float32 `json:"value"`
	Description string `json:"srting"`
}