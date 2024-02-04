package Measurements

import(
	Units "mensuration/Units"
	"errors"
	"math"
)

type Mass struct {
	Value float64
	Unit  Units.MassUnit
}

func NewMass(value float64, unit Units.MassUnit) Mass{
	return Mass{
		Value: value,
		Unit: unit,
	}
}

func(a Mass)ConvertTo(b interface{}) float64{

	aValueInMeters := a.Value * float64(Units.MassUnit(a.Unit).GetConversionUnit())
	aConvertedToTargetValue := aValueInMeters / (Units.MassUnit(b.(Mass).Unit).GetConversionUnit())

	return aConvertedToTargetValue

}


func (a Mass) CompareValue(b interface{}) (int,error){

	if(a.Value < 0 || b.(Mass).Value < 0){return 0,errors.New("negative values")}
	if(a.Value == 0 && b.(Mass).Value == 0){return 0,nil}

	aConvertedToTargetValue := a.ConvertTo(b)

	if(aConvertedToTargetValue == b.(Mass).Value){
		return 0,nil;
	}else if(aConvertedToTargetValue > b.(Mass).Value){
		return 1,nil;
	}else{
		return -1,nil;
	}

}

func (a Mass) Add(b interface{}) Mass{

	aConvertedToTargetValue := a.ConvertTo(b)
	return Mass{
		(aConvertedToTargetValue + b.(Mass).Value),
		b.(Mass).Unit,
	} 

}

func (a Mass) Subtract(b interface{}) Mass{

	aConvertedToTargetValue := a.ConvertTo(b)
	return Mass{
		math.Abs(aConvertedToTargetValue - b.(Mass).Value),
		b.(Mass).Unit,
	} 

}