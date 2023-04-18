//Version: v1
package models

type IC000066Request struct {
	CustomerId           string `json:"customerId" validate:"required"`
	CreditApplySerialNum string `json:"creditApplySerialNum" validate:"required"`
	OperationType        string `json:"operationType" validate:"required"` //1- Create; 2- Increase; 3- Decrease; 4- Freeze; 5- Unfreeze; 6- Expire
}

type IC000066Response struct {
}
