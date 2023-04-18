//Version: v1.0.0
package models

type NT000018I struct {
	RelateMcId string `json:"relateMcId" description:"Relate mc id" validate:"required"`
}

type NT000018O struct {
}

func (i *NT000018I) GetServiceKey() string {
	return "NT000018"
}
