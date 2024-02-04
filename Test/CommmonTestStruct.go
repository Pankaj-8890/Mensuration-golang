package Test

type TestCase struct {
	TestName                   string
	MeasurementInstance        interface{}
	AnotherMeasurementInstance interface{}
	ExpectedValue              interface{}
	WantErr                    error
}
