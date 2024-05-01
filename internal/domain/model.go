package domain

type Model interface {
	ConvertFromArray(fields []interface{})
	GetFields() []interface{}
}
