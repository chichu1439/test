package models

type PF000009I struct {
	StrategyId string `json:"strategyId" description:"Strategy id" validate:"required"`
}

type PF000009O struct {
}

func (*PF000009I) GetServiceKey() string {
	return "PF000009"
}
