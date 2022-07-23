package mailer

import "github.com/thanishsid/goserver/config"

type LinkMailTemplateData struct {
	Title            string
	PrimaryMessage   string
	SecondaryMessage string
	Link             string
}

type TemplateData[T any] struct {
	CompanyName string
	Data        T
}

func NewDataWithDefault[T any](data T) TemplateData[T] {
	return TemplateData[T]{
		CompanyName: config.C.CompanyName,
		Data:        data,
	}
}
