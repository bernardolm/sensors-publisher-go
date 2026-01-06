package api

// https://www.home-assistant.io/integrations/mqtt/#single-component-discovery-payload

type ____Component struct {
	// Type              string         `json:"type,omitempty"`                // optional

	AutomationType    AutomationType `json:"atype,omitempty"`        // optional
	DeviceClass       DeviceClass    `json:"dev_cla,omitempty"`      // optional
	Payload           string         `json:"pl,omitempty"`           // optional
	Platform          Platform       `json:"p,omitempty"`            // optional
	Subtype           string         `json:"stype,omitempty"`        // optional
	Topic             string         `json:"t,omitempty"`            // optional
	UniqueID          string         `json:"uniq_id,omitempty"`      // optional
	UnitOfMeasurement string         `json:"unit_of_meas,omitempty"` // optional
	ValueTemplate     string         `json:"val_tpl,omitempty"`      // optional
}
