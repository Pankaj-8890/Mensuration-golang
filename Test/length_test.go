package Test

import (
	"errors"
	"testing"

	MeasurementPkg "mensuration/Measurements"
	Units "mensuration/Units"

	"github.com/stretchr/testify/assert"
)


func TestConvertTo(t *testing.T){
    testCases := []TestCase{
		{
			TestName:                   "TestErrorWhenConvert1kmto1kg",
			MeasurementInstance:        MeasurementPkg.NewLength(100, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewMass(1, Units.KG),
			ExpectedValue:              MeasurementPkg.NewLength(0, ""),
			WantErr:                    errors.New("can't convert"),
		},
        {
			TestName:                   "TestErrorWhenConvert1cmto1m",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.CM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1, Units.M),
			ExpectedValue:              MeasurementPkg.NewLength(0.01, Units.M),
			WantErr:                    nil,
		},
        {
			TestName:                   "TestErrorWhenConvert1gto1kg",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.CM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1, Units.M),
			ExpectedValue:              MeasurementPkg.NewLength(0.001, Units.KM),
			WantErr:                    nil,
		},
	}
    for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
		    res,err := tc.MeasurementInstance.(MeasurementPkg.Length).ConvertTo(tc.AnotherMeasurementInstance)

            if (err == tc.WantErr && err != nil){
                // t.Errorf("%s:    res error %v, want %v",tc.TestName, Err,tc.WantErr)
                assert.Equal(t,err,tc.WantErr)
            }else if  res != tc.ExpectedValue {
                // t.Errorf("%s: Expected %v, but   res %v", tc.TestName, res, tc.ExpectedValue)
				assert.NotEqual(t,res,tc.ExpectedValue)
			}

		})
	}

}
func TestLengthComparison(t *testing.T) {
	testCases := []TestCase{
        {
			TestName:                   "TestErrorWhenCompare1lto1kg",
			MeasurementInstance:        MeasurementPkg.NewLength(100, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewMass(1, Units.MG),
			ExpectedValue:              MeasurementPkg.NewLength(0, ""),
			WantErr:                    errors.New("can't compare"),
		},
		{
			TestName:                   "TestFailedWithNegativeLength",
			MeasurementInstance:        MeasurementPkg.NewLength(-1, Units.CM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(-1000, Units.M),
			ExpectedValue:              0,
			WantErr:                    errors.New("negative values"),
		},
		{
			TestName:                   "TestPassedWhenLengthEqual",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1000, Units.M),
			ExpectedValue:              0,
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenLengthInequal",
			MeasurementInstance:        MeasurementPkg.NewLength(10, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1000, Units.M),
			ExpectedValue:              1,
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenLengthsAreZero",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1000, Units.M),
			ExpectedValue:              0,
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
		    res, err := tc.MeasurementInstance.(MeasurementPkg.Length).CompareValue(tc.AnotherMeasurementInstance)

			if (err == tc.WantErr && err != nil){
                // t.Errorf("%s:    res error %v, want %v",tc.TestName, Err,tc.WantErr)
                assert.Equal(t,err,tc.WantErr)
            }else if  res != tc.ExpectedValue {
                // t.Errorf("%s: Expected %v, but   res %v", tc.TestName, res, tc.ExpectedValue)
				assert.NotEqual(t,res,tc.ExpectedValue)
			}

		})
	}

}

func TestAddLengths(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestPassedWhenAdd1kmAnd1000m",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1000, Units.M),
			ExpectedValue:              MeasurementPkg.NewLength(2000, Units.M),
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenLengthInequal",
			MeasurementInstance:        MeasurementPkg.NewLength(10, Units.CM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1, Units.M),
			ExpectedValue:              MeasurementPkg.NewLength(1.1, Units.M),
			WantErr:                    nil,
		},
        {
			TestName:                   "TestErrorWhenAdd1kmto1kg",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewMass(1, Units.KG),
			ExpectedValue:              MeasurementPkg.NewLength(0, ""),
			WantErr:                    errors.New("can't add"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
		    res,err := tc.MeasurementInstance.(MeasurementPkg.Length).Add(tc.AnotherMeasurementInstance)

            if (err == tc.WantErr && err != nil) {
                // t.Errorf("%s: res error %v, want %v",tc.TestName, err,tc.WantErr)
                assert.Equal(t,err,tc.WantErr)
            }else if  res != tc.ExpectedValue {
                // t.Errorf("%s: Expected %v, but   res %v", tc.TestName, res, tc.ExpectedValue)
				assert.NotEqual(t,res,tc.ExpectedValue)
			}

		})
	}
}

func TestSubtractLengths(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestPassedWhenAdd1kmAnd1000m",
			MeasurementInstance:        MeasurementPkg.NewLength(1, Units.KM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1000, Units.M),
			ExpectedValue:              MeasurementPkg.NewLength(0, Units.M),
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenLengthInequal",
			MeasurementInstance:        MeasurementPkg.NewLength(10, Units.CM),
			AnotherMeasurementInstance: MeasurementPkg.NewLength(1, Units.M),
			ExpectedValue:              MeasurementPkg.NewLength(0.8, Units.M),
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
		    res,err := tc.MeasurementInstance.(MeasurementPkg.Length).Subtract(tc.AnotherMeasurementInstance)
            if (err == tc.WantErr && err != nil){
                // t.Errorf("%s:    res error %v, want %v",tc.TestName, Err,tc.WantErr)
                assert.Equal(t,err,tc.WantErr)
            }else if  res != tc.ExpectedValue {
                // t.Errorf("%s: Expected %v, but   res %v", tc.TestName, res, tc.ExpectedValue)
				assert.NotEqual(t,res,tc.ExpectedValue)
			}

		})
	}
}
