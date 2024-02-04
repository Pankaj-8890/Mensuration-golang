package Measurements

import(
	Units "mensuration/Units"
	"errors"
	"math"
	"reflect"
	"fmt"
)

type Volume struct {
	Value float64
	Unit  Units.VolumeUnit
}

func NewVolume(value float64, unit Units.VolumeUnit) Volume{
	return Volume{
		Value: value,
		Unit: unit,
	}
}

func(a Volume)ConvertTo(b interface{}) (interface{},error){

	if(reflect.TypeOf(a) != reflect.TypeOf(Volume{}) || reflect.TypeOf(b) != reflect.TypeOf(Volume{})) {return Volume{},errors.New("can't convert")}

	aValueInMeters := a.Value * float64(Units.VolumeUnit(a.Unit).GetConversionUnit())
	aConvertedToTargetValue := aValueInMeters / (Units.VolumeUnit(b.(Volume).Unit).GetConversionUnit())

	return aConvertedToTargetValue,nil

}


func (a Volume) CompareValue(b interface{}) (int,error){

	if(a.Value < 0 || b.(Volume).Value < 0){return 0,errors.New("negative values")}
	if(a.Value == 0 && b.(Volume).Value == 0){return 0,nil}

	aConvertedToTargetValue,err := a.ConvertTo(b)

	if err!=nil{
		fmt.Printf("Invalid arrguments %v",err)
	}

	if(aConvertedToTargetValue.(Volume).Value == b.(Volume).Value){
		return 0,nil;
	}else if(aConvertedToTargetValue.(Volume).Value > b.(Volume).Value){
		return 1,nil;
	}else{
		return -1,nil;
	}

}

func (a Volume) Add(b interface{}) (Volume,error){

	if(reflect.TypeOf(a) != reflect.TypeOf(Volume{}) || reflect.TypeOf(b) != reflect.TypeOf(Volume{})) {return Volume{},errors.New("can't add")}

	aConvertedToTargetValue,err := a.ConvertTo(b)
	if err!=nil{
		fmt.Printf("Invalid arrguments %v",err)
	}
	return Volume{
		(aConvertedToTargetValue.(Volume).Value + b.(Volume).Value),
		b.(Volume).Unit,
	},nil  

}

func (a Volume) Subtract(b interface{}) (Volume,error){

	if(reflect.TypeOf(a) != reflect.TypeOf(Volume{}) || reflect.TypeOf(b) != reflect.TypeOf(Volume{})) {return Volume{},errors.New("can't subtract")}

	aConvertedToTargetValue,err := a.ConvertTo(b)
	if err!=nil{
		fmt.Printf("Invalid arrguments %v",err)
	}
	return Volume{
		math.Abs(aConvertedToTargetValue.(Volume).Value - b.(Volume).Value),
		b.(Volume).Unit,
	},nil 

}









// package Volume

// type Volume struct {
// 	Value float64
// 	Unit  VolumeUnit
// }

// func(l *Volume)GetValue() float64{
// 	return l.Value
// }

// func(l *Volume)GetUnit() interface{}{
// 	return l.Unit
// }


// func(l *Volume)SetValue(L float64){
// 	l.Value = L
// }


// func (l *Volume) GetConversionUnit() float64 {
// 	return metersPerUnit[l.Unit]
// }


// func NewVolume(value float64, unit VolumeUnit) Volume{
// 	return Volume{
// 		Value: value,
// 		Unit: unit,
// 	}
// }

