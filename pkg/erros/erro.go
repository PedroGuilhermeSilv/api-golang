package errors

import "errors"

var ErrorIDIsRequired = errors.New("id is required")
var ErrorNameIsRequired = errors.New("name is required")
var ErrorPriceIsRequired = errors.New("price is required")
var ErrorPriceIsInvalid = errors.New("price is invalid")
var ErrorNameIsInvalid = errors.New("name is invalid")
var ErrorIDIsInvalid = errors.New("id is invalid")



var ErrorProductNotFound = errors.New("product not found")
var ErrorProductDelete = errors.New("erro product delete")
