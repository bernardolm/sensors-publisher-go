package api

// https://www.home-assistant.io/docs/configuration/customizing-devices/#icon
// https://www.home-assistant.io/integrations/homeassistant/#icon
// https://www.home-assistant.io/integrations/mqtt/#device-discovery-payload
// https://www.home-assistant.io/integrations/mqtt/#discovery-migration-example-with-a-device-automation-and-a-sensor
// https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery
// https://www.home-assistant.io/integrations/mqtt/#single-component-discovery-payload
// https://www.home-assistant.io/integrations/mqtt/#supported-abbreviations-in-mqtt-discovery-messages

type Discovery struct {
	Entity
	Sensor

	// Components map[string]Component `json:"cmps,omitempty"` // optional // NOTE: new discovery version

	Device  Device `json:"dev,omitempty"` // optional
	Origin  Origin `json:"o,omitempty"`   // optional
	Payload string `json:"pl,omitempty"`  // optional
	Topic   string `json:"t,omitempty"`   // optional
}
