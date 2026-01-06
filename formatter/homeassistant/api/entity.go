package api

// https://pictogrammers.com/library/mdi/
// https://www.home-assistant.io/integrations/homeassistant/#editing-entity-settings-in-yaml

type Entity struct {
	AssumedState      bool        `json:"assumed_state,omitempty"` // optional
	DeviceClass       DeviceClass `json:"dev_cla,omitempty"`       // optional
	EntityPicture     string      `json:"ent_pic,omitempty"`       // optional
	FriendlyName      string      `json:"friendly_name,omitempty"` // optional
	Icon              string      `json:"ic,omitempty"`            // optional
	InitialState      bool        `json:"initial_state,omitempty"` // optional
	UnitOfMeasurement string      `json:"unit_of_meas,omitempty"`  // optional
}
