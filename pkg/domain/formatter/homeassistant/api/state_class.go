package api

// https://developers.home-assistant.io/docs/core/entity/sensor/#available-state-classes

type StateClass string

const (
	MeasurementStateClass      StateClass = "measurement"
	MeasurementAngleStateClass StateClass = "measurement_angle"
	TotalStateClass            StateClass = "total"
	TotalIncreasingStateClass  StateClass = "total_increasing"
)
