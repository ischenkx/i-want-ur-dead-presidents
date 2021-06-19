package models

type Product struct {
	Id string
	Inn string
}

type Score struct {
	OverallScore int
	CourtScore int
	FinKoefScore int
	SmartScore int
}

type Response struct {
	Id       string
	Inn      string
	Name     string
	FullName string
	Score    Score
}