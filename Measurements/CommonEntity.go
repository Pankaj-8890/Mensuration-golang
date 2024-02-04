package Measurements

type CommonMethods interface {
	ConvertTo(interface{},interface{}) interface{}
	Compare(interface{}) (int, error)
	Add(interface{}) interface{}
	Subtract(interface{}) interface{}
}
