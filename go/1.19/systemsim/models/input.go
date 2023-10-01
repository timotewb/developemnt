package models

//----------------------------------------------------------------------------------------
// Input
//----------------------------------------------------------------------------------------
type InputType struct{
	ID int	`json:"id"`
	Name string `json:"name"`
	Attributes []InputAttribute `json:"attributes"`
}

type InputAttribute struct{
	ID int	`json:"id"`
	Value float32 `json:"value"`
	Description string `json:"description"`
}