package mailer

import "github.com/thanishsid/goserver/config"

type LinkMailData struct {
	To               string
	Subject          string
	Title            string
	PrimaryMessage   string
	SecondaryMessage string
	Link             string
}

type DefaultData[T any] struct {
	CompanyName string
	Data        T
}

func NewDataWithDefaults[T any](data T) DefaultData[T] {
	return DefaultData[T]{
		CompanyName: config.C.CompanyName,
		Data:        data,
	}
}
