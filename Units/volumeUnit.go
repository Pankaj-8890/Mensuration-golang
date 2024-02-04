package Units

type VolumeUnit string

const (
	L  VolumeUnit = "L"
	ML VolumeUnit = "ML"
	CuM VolumeUnit = "CuM"
)

var volumePerMl = map[VolumeUnit]float64{
	L:  1.0,
    CuM:1000.0,
    ML: 0.001,
}

func (l VolumeUnit) GetConversionUnit() float64 {
	return volumePerMl[l]
}
