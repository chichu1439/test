package models

type PD000089I struct {
	ProductId string      `json:"productId" description:"Product Id"`
	ServiceId string      `json:"serviceId" description:"Service Id"`
	FeatureId string      `json:"featureId" description:"Feature Id"`
	ListLimit []ListLimit `json:"listLimit" description:"List limit"`
}

type ListLimit struct {
	LimitTarget string `json:"limitTarget" description:"Limit target"`
	TargetValue string `json:"targetValue" description:"Target value"`
}

type PD000089O struct {
	ListLimitReturn []ListLimitReturn `json:"listLimitReturn" description:"List limit return"`
}

type ListLimitReturn struct {
	LimitTarget     string `json:"limitTarget" description:"Limit target"`
	TargetValue     string `json:"targetValue" description:"Target value"`
	UpperLimitValue string `json:"upperLimitValue" description:"Upper limit value"`
	LowerLimitValue string `json:"lowerLimitValue" description:"Lower limit value"`
	UpperPass       string `json:"upperPass" description:"Upper pass"`
	LowerPass       string `json:"lowerPass" description:"Lower pass"`
}

func (*PD000089I) GetServiceKey() string {
	return "PD000089"
}
