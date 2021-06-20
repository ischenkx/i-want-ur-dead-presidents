package models

type Entity struct {
	ID               string
	Title            string
	ShortDesc        string
	LongDesc         string
	MoneyGoal        float64
	OwnerID          string
	DirectorFullName string
	FullCompanyName  string
	Inn              string
	ORGNN            string
	CompanyEmail     string
	OwnerFullName    string
	OwnerPost        string
	PassportData     string
	PictureUrl       string
	ActivityField    string
}
