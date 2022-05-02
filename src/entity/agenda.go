package entity

type Agenda struct {
	Commitments []Commitment
	Owner       Politician
	PublicURL   string
}
