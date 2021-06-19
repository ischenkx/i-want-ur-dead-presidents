package models

type Entity struct {
	ID        string
	Title     string
	ShortDesc string
	LongDesc  string
	MoneyGoal int64
	OwnerID   string
}