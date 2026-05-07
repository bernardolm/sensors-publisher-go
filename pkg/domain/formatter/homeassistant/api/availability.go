package api

// https://www.home-assistant.io/integrations/mqtt/#mqtt-configuration-variables
// https://www.home-assistant.io/integrations/mqtt/#supported-abbreviations-in-mqtt-discovery-messages
// https://www.home-assistant.io/integrations/mqtt/#using-availability-topics

type Availability struct {
	PayloadAvailable    string   `json:"pl_avail,omitempty"`     // optional
	PayloadNotAvailable string   `json:"pl_not_avail,omitempty"` // optional
	Topic               string   `json:"t,omitempty"`            // required
	ValueTemplate       Template `json:"val_tpl,omitempty"`      // optional
}
