//Version: v1
package models

type FP900066Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FP900066Response struct {
	Status string `json:"status"`
}

func (*FP900066Request) GetServiceKey() string {
	return "FP900066"
}
