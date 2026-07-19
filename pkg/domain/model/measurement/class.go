package measurement

// Class identifies the information delivered by a measurement sensor.
type Class string

const (
	ClassTemperature         Class = "temperature"
	ClassAtmosphericPressure Class = "atmospheric_pressure"
	ClassHumidity            Class = "humidity"
)
