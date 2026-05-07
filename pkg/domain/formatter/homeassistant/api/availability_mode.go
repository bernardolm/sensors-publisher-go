package api

type AvailabilityMode string

const (
	AllAvailabilityMode    AvailabilityMode = "all"
	AnyAvailabilityMode    AvailabilityMode = "any"
	LatestAvailabilityMode AvailabilityMode = "latest"
)
