package requests

type QuantityUnitText struct {
	QuantityUnit     	string  `json:"QuantityUnit"`
	Language         	string  `json:"Language"`
	QuantityUnitName 	string  `json:"QuantityUnitName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
