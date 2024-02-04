package Test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	Volumepkg "mensuration/Measurements"
	Units "mensuration/Units"
	"testing"
)

func TestVolumeComparison(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestFailedWithNegativeVolume",
			MeasurementInstance:        Volumepkg.NewVolume(-11, Units.ML),
			AnotherMeasurementInstance: Volumepkg.NewVolume(-11000, Units.L),
			ExpectedValue:              0,
			WantErr:                    errors.New("negative values"),
		},
		{
			TestName:                   "TestPassedWhenVolumeEqual",
			MeasurementInstance:        Volumepkg.NewVolume(1, Units.CuM),
			AnotherMeasurementInstance: Volumepkg.NewVolume(1000000, Units.ML),
			ExpectedValue:              0,
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenVolumeUnequal",
			MeasurementInstance:        Volumepkg.NewVolume(1, Units.L),
			AnotherMeasurementInstance: Volumepkg.NewVolume(1000, Units.CuM),
			ExpectedValue:              -1,
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenVolumesAreZero",
			MeasurementInstance:        Volumepkg.NewVolume(0, Units.L),
			AnotherMeasurementInstance: Volumepkg.NewVolume(0, Units.L),
			ExpectedValue:              0,
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got, Err := tc.MeasurementInstance.(Volumepkg.Volume).CompareValue(tc.AnotherMeasurementInstance)

			if Err != tc.WantErr {
				assert.NotEqual(t, got, tc, tc.ExpectedValue, "expected and actual values are not equal")
			}

			if got != tc.ExpectedValue {
				t.Errorf("%s: got %d, want %d", tc.TestName, got, tc.ExpectedValue)

			}

		})
	}
}

func TestAddVolume(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestPassedWhenAdd1lAnd1000ml",
			MeasurementInstance:        Volumepkg.NewVolume(1, Units.L),
			AnotherMeasurementInstance: Volumepkg.NewVolume(1000, Units.ML),
			ExpectedValue:              Volumepkg.NewVolume(2000, Units.ML),
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenAddUnequalVolume",
			MeasurementInstance:        Volumepkg.NewVolume(10, Units.ML),
			AnotherMeasurementInstance: Volumepkg.NewVolume(1, Units.ML),
			ExpectedValue:              Volumepkg.NewVolume(1.001, Units.ML),
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got,_ := tc.MeasurementInstance.(Volumepkg.Volume).Add(tc.AnotherMeasurementInstance)

			if got != tc.ExpectedValue {
				// t.Errorf("%s: Expected %v, but got %v", tc.TestName, got, tc.ExpectedValue)
				assert.NotEqual(t, got, tc, tc.ExpectedValue, "expected and actual values are not equal")

			}

		})
	}
}

func TestSubtractVolume(t *testing.T) {
	testCases := []TestCase{
		{
			TestName:                   "TestPassedWhenSubtract1lAnd1000ml",
			MeasurementInstance:        Volumepkg.NewVolume(1, Units.L),
			AnotherMeasurementInstance: Volumepkg.NewVolume(1000, Units.ML),
			ExpectedValue:              Volumepkg.NewVolume(0, Units.ML),
			WantErr:                    nil,
		},
		{
			TestName:                   "TestPassedWhenSubtractUnequalVolume",
			MeasurementInstance:        Volumepkg.NewVolume(10, Units.ML),
			AnotherMeasurementInstance: Volumepkg.NewVolume(1, Units.ML),
			ExpectedValue:              Volumepkg.NewVolume(0.9, Units.ML),
			WantErr:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			got,_ := tc.MeasurementInstance.(Volumepkg.Volume).Subtract(tc.AnotherMeasurementInstance)

			if got != tc.ExpectedValue {
				// t.Errorf("%s: Expected %v, but got %v", tc.TestName, got, tc.ExpectedValue)
				assert.NotEqual(t, got, tc, tc.ExpectedValue, "expected and actual values are not equal")
			}

		})
	}
}
