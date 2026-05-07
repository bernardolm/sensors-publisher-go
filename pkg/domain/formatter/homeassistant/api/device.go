package api

// https://www.home-assistant.io/integrations/mqtt/#supported-abbreviations-in-mqtt-discovery-messages
// https://www.home-assistant.io/integrations/sensor.mqtt/#device

type Device struct {
	ConfigurationURL string     `json:"cu,omitempty"`         // optional
	Connections      [][]string `json:"cns,omitempty"`        // optional
	HardwareVersion  string     `json:"hw,omitempty"`         // optional
	Identifiers      []string   `json:"ids,omitempty"`        // optional
	Manufacturer     string     `json:"mf,omitempty"`         // optional
	Model            string     `json:"mdl,omitempty"`        // optional
	ModelID          string     `json:"mdl_id,omitempty"`     // optional
	Name             string     `json:"name,omitempty"`       // optional
	SerialNumber     string     `json:"sn,omitempty"`         // optional
	SoftwareVersion  string     `json:"sw,omitempty"`         // optional
	SuggestedArea    string     `json:"sa,omitempty"`         // optional
	ViaDevice        string     `json:"via_device,omitempty"` // optional // https://www.home-assistant.io/integrations/sensor.mqtt/#via_device
}
