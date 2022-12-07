package models

type SMP0000015I struct {
}

type SMP0000015O struct {
	Wallets []Wallets
}

func (i *SMP0000015I) GetServiceKey() string {
	return "MP000015"
}
