package outgress

type PD000078I struct {
	ProductId  string `json:"productId" description:"Product id" validate:"required"`
	NotifyType string `json:"notifyType" description:"Notify type" validate:"required"`
}

type PD000078O struct {
	ListStrategy []ListStrategy `json:"listStrategy" description:"List strategy"`
}

type ListStrategy struct {
	StrategyId int `json:"strategyId" description:"Strategy id"`
}

func (i *PD000078I) GetServiceKey() string {
	return "PD000078"
}
