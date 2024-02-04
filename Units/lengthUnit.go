package Units

type LengthUnit string

const (
	KM LengthUnit = "KM"
	CM LengthUnit = "CM"
	M  LengthUnit = "M"
	MM LengthUnit = "MM"
)

var metersPerUnit = map[LengthUnit]float64{
	KM: 1000.0,
	CM: 0.01,
	M:  1.0,
	MM: 0.001,
}

func (l LengthUnit) GetConversionUnit() float64 {
	return metersPerUnit[l]
}


