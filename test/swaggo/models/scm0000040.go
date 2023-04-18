package models

type CM000040I struct {
}

type CM000040O struct {
	RandomKey string `json:"randomKey" description:"Random key"`
}

func (*CM000040I) GetServiceKey() string {
	return "CM000040"
}
