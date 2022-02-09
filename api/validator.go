package api

import "github.com/go-playground/validator/v10"

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"JPY": true,
}

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return supportedCurrencies[currency]
	}
	return false
}
