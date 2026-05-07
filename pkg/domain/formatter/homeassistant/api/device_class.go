package api

// https://developers.home-assistant.io/docs/core/entity/sensor/#available-device-classes
// https://www.home-assistant.io/integrations/homeassistant/#device-class
// https://www.home-assistant.io/integrations/sensor/#device-class

type DeviceClass string

const (
	// atmospheric_pressure Atmospheric pressure in cbar, bar, hPa, mmHg, inHg, kPa, mbar, Pa or psi
	AtmosphericPressureDeviceClass DeviceClass = "atmospheric_pressure"

	// humidity Percentage of humidity in the air in %
	HumidityDeviceClass DeviceClass = "humidity"

	// battery Percentage of battery that is left in %
	BatteryDeviceClass DeviceClass = "battery"

	// date Date string (ISO 8601)
	DateDeviceClass DeviceClass = "date"

	// enum Has a limited set of (non-numeric) states
	EnumDeviceClass DeviceClass = "enum"

	// moisture Percentage of water in a substance in %
	MoistureDeviceClass DeviceClass = "moisture"

	// ph Potential hydrogen (pH) value of a water solution
	PhDeviceClass DeviceClass = "ph"

	// power Power in mW, W, kW, MW, GW or TW
	PowerDeviceClass DeviceClass = "power"

	// pressure Pressure in mPa, Pa, hPa, kPa, bar, cbar, mbar, mmHg, inHg, inH₂O or psi
	PressureDeviceClass DeviceClass = "pressure"

	// signal_strength Signal strength in dB or dBm
	SignalStrengthDeviceClass DeviceClass = "signal_strength"

	// temperature Temperature in °C, °F or K
	TemperatureDeviceClass DeviceClass = "temperature"

	// temperature_delta Temperature difference between two measurements in °C, °F, or K
	TemperatureDeltaDeviceClass DeviceClass = "temperature_delta"

	// timestamp Datetime object or timestamp string (ISO 8601)
	TimestampDeviceClass DeviceClass = "timestamp"

	// voltage Voltage in V, mV, µV, kV, MV
	VoltageDeviceClass DeviceClass = "voltage"

	// volume Generic volume in L, mL, gal, fl. oz., m³, ft³, CCF, or MCF
	VolumeDeviceClass DeviceClass = "volume"

	// volume_flow_rate Volume flow rate in m³/h, m³/min, m³/s, ft³/min, L/h, L/min, L/s, gal/h, gal/min, or mL/s
	VolumeFlowRateDeviceClass DeviceClass = "volume_flow_rate"

	// volume_storage Generic stored volume in L, mL, gal, fl. oz., m³, ft³, CCF, or MCF
	VolumeStorageDeviceClass DeviceClass = "volume_storage"

	// water Water consumption in L, gal, m³, ft³, CCF, or MCF
	WaterDeviceClass DeviceClass = "water"
)
