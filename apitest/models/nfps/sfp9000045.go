package models

type Sfp9000045TpartRelation struct {
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	PartName         string `json:"partName"`
}

type Sfp9000045I struct {

}

func (h *Sfp9000045I) GetServiceKey() string {
	return "FP900045"
}
