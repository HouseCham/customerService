package model

type ValidatorError struct {
	Tag   string 
	Field string 
	Err   string
}