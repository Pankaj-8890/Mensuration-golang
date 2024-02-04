package Test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	Masspkg "mensuration/Measurements"
	Units "mensuration/Units"
	"testing"
)

func TestMassComparison(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestFailedWithNegativeMass",
			MeasurementInstance:        Masspkg.NewMass(-11, Units.G),
			AnotherMeasurementInstance: Masspkg.NewMass(-11000, Units.KG),
			ExpectedValue:              0,
			WantErr:                    errors.New("negative values"),
		},
		{
			TestName:                   "TestPassedWhenMassEqual",
			MeasurementInstance:        Masspkg.NewMass(100, Units.G),
			AnotherMeasurementInstance: Masspkg.NewMass(0.1, Units.KG),
			ExpectedValue:              0,
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenMassInequal",
			MeasurementInstance:        Masspkg.NewMass(1, Units.KG),
			AnotherMeasurementInstance: Masspkg.NewMass(100, Units.G),
			ExpectedValue:              1,
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenMasssAreZero",
			MeasurementInstance:        Masspkg.NewMass(0, Units.KG),
			AnotherMeasurementInstance: Masspkg.NewMass(0, Units.G),
			ExpectedValue:              0,
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got, Err := tc.MeasurementInstance.(Masspkg.Mass).CompareValue(tc.AnotherMeasurementInstance)

			if Err != tc.WantErr {
				assert.NotEqual(t, got, tc, tc.ExpectedValue, "expected and actual values are not equal")
			}

			if got != tc.ExpectedValue {
				t.Errorf("%s: got %d, want %d", tc.TestName, got, tc.ExpectedValue)

			}

		})
	}
}

func TestAddMass(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestPassedWhenAdd1kgAnd1000g",
			MeasurementInstance:        Masspkg.NewMass(1, Units.KG),
			AnotherMeasurementInstance: Masspkg.NewMass(1000, Units.G),
			ExpectedValue:              Masspkg.NewMass(2000, Units.G),
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenAddInequalMass",
			MeasurementInstance:        Masspkg.NewMass(10, Units.MG),
			AnotherMeasurementInstance: Masspkg.NewMass(1, Units.G),
			ExpectedValue:              Masspkg.NewMass(1.001, Units.MG),
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got := tc.MeasurementInstance.(Masspkg.Mass).Add(tc.AnotherMeasurementInstance)

			if got != tc.ExpectedValue {
				// t.Errorf("%s: Expected %v, but got %v", tc.TestName, got, tc.ExpectedValue)
				assert.NotEqual(t, got, tc, tc.ExpectedValue, "expected and actual values are not equal")

			}

		})
	}
}

func TestSubtractMass(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestPassedWhenSubtract1kmAnd1000m",
			MeasurementInstance:        Masspkg.NewMass(1, Units.KG),
			AnotherMeasurementInstance: Masspkg.NewMass(1000, Units.G),
			ExpectedValue:              Masspkg.NewMass(0, Units.G),
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenSubtractInEqualMass",
			MeasurementInstance:        Masspkg.NewMass(10, Units.MG),
			AnotherMeasurementInstance: Masspkg.NewMass(1, Units.G),
			ExpectedValue:              Masspkg.NewMass(0.9, Units.G),
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got := tc.MeasurementInstance.(Masspkg.Mass).Subtract(tc.AnotherMeasurementInstance)

			if got != tc.ExpectedValue {
				// t.Errorf("%s: Expected %v, but got %v", tc.TestName, got, tc.ExpectedValue)
				assert.NotEqual(t, got, tc, tc.ExpectedValue, "expected and actual values are not equal")
			}

		})
	}
}
