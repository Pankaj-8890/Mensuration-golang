package Measurements

import (
	"errors"
	"fmt"
	"math"
	Units "mensuration/Units"
	"reflect"
)
type Length struct {
	Value float64
	Unit  Units.LengthUnit
}

func NewLength(value float64, unit Units.LengthUnit) Length{
	return Length{
		Value: value,
		Unit: unit,
	}
}

func(a Length)ConvertTo(b interface{}) (interface{},error){

	if(reflect.TypeOf(a) != reflect.TypeOf(Length{}) || reflect.TypeOf(b) != reflect.TypeOf(Length{})) {return Length{},errors.New("can't convert")}

	aValueInMeters := a.Value * float64(Units.LengthUnit(a.Unit).GetConversionUnit())
	aConvertedToTargetValue := aValueInMeters / (Units.LengthUnit(b.(Length).Unit).GetConversionUnit())

	return Length{aConvertedToTargetValue,b.(Length).Unit},nil

}


func (a Length) CompareValue(b interface{}) (int,error){
	

	if(reflect.TypeOf(a) != reflect.TypeOf(Length{}) || reflect.TypeOf(b) != reflect.TypeOf(Length{})) {return 0,errors.New("can't compare")}

	if(a.Value < 0 || b.(Length).Value < 0){return 0,errors.New("negative values")}
	if(a.Value == 0 && b.(Length).Value == 0){return 0,nil}

	aConvertedToTargetValue,err := a.ConvertTo(b)

	if err!=nil{
		fmt.Printf("Invalid arrguments %v",err)
	}

	if(aConvertedToTargetValue.(Length).Value == b.(Length).Value){
		return 0,nil;
	}else if(aConvertedToTargetValue.(Length).Value > b.(Length).Value){
		return 1,nil;
	}else{
		return -1,nil;
	}

}

func (a Length) Add(b interface{}) (Length,error){

	if(reflect.TypeOf(a) != reflect.TypeOf(Length{}) || reflect.TypeOf(b) != reflect.TypeOf(Length{})) {return Length{},errors.New("can't add")}

	aConvertedToTargetValue,err := a.ConvertTo(b)
	if err!=nil{
		fmt.Printf("Invalid arrguments %v",err)
	}
	return Length{
		(aConvertedToTargetValue.(Length).Value + b.(Length).Value),
		b.(Length).Unit,
	},nil 

}

func (a Length) Subtract(b interface{}) (Length,error){

	if(reflect.TypeOf(a) != reflect.TypeOf(Length{}) || reflect.TypeOf(b) != reflect.TypeOf(Length{})) {return Length{},errors.New("can't subtract")}

	aConvertedToTargetValue,err := a.ConvertTo(b)
	if err!=nil{
		fmt.Printf("Invalid arrguments %v",err)
	}
	return Length{
		math.Abs(aConvertedToTargetValue.(Length).Value - b.(Length).Value),
		b.(Length).Unit,
	},nil 

}







