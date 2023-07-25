package requests

type QuantityUnit struct {
	QuantityUnit 		string  `json:"QuantityUnit"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
