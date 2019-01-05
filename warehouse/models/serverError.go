package models

type ServerError struct {
	Status             int
	ClientErrorMessage string
	Error              error
}
