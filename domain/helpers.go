package domain

type Validateable interface {
	Validate() error
}

func CheckValidity(obj Validateable) error {
	return obj.Validate()
}
