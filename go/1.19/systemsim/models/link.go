package models

//----------------------------------------------------------------------------------------
// Link
//----------------------------------------------------------------------------------------
type LinkType struct{
	ID int	`json:"id"`
	Name string `json:"name"`
	Items []ItemType `json:"items"`
	Equation string `json:"equation"`
	EquationEvaluated string `json:"equation_evaluated"`
	Output string `json:"output"`
}

type ItemType struct{
	ID int	`json:"id"`
	ItemID int `json:"item_id"`
	LinkType string `json:"link_type"`
	ValueID int `json:"value_id"`
}