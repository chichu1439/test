//Version: v0.0.1
package models

type PI000009I struct {
	StrategyId string `json:"strategyId" validate:"required" description:"strategy Id"`
}

type PI000009O struct {
}

func (*PI000009I) GetServiceKey() string {
	return "PI000009"
}
