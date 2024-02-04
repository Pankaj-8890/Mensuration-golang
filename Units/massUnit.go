package Units

type MassUnit string

const (
	KG MassUnit = "KG"
	G  MassUnit = "G"
	MG MassUnit = "MG"
)

var massPerGram = map[MassUnit]float64{
	KG : 1000.0,
    MG : 0.001,
    G :  1.0,	
}

func (l MassUnit) GetConversionUnit() float64 {
	return massPerGram[l]
}
