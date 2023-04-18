package models

type FP000029I struct {
	Currency                string  `json:"currency" validate:"required"`
	Amount                  float64 `json:"amount" validate:"required"`
	SweepingType            string  `json:"sweepingType" validate:"required"`
	ParticipantClearingCode string  `json:"participantClearingCode" validate:"required"`
}
type FP000029O struct {
	Status string
}

func (*FP000029I) GetServiceKey() string {
	return "FP000029"
}

type FP000029Outgress struct {
	Body []byte `json:"body"`
}

func (*FP000029Outgress) GetServiceKey() string {
	return "FP200029"
}

type FP000029OutgressResponse struct {
	Body []byte `json:"body"`
}
