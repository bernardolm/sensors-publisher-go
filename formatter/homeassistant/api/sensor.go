package api

// https://developers.home-assistant.io/docs/core/entity/sensor
// https://developers.home-assistant.io/docs/core/entity/sensor/#available-state-classes
// https://developers.home-assistant.io/docs/core/entity#generic-properties
// https://github.com/home-assistant/core/blob/dev/homeassistant/components/mqtt/sensor.py
// https://www.home-assistant.io/integrations/sensor.mqtt
// https://www.home-assistant.io/integrations/sensor.mqtt/#sensor-mqtt-configuration-variables

type Sensor struct {
	Availability              []Availability    `json:"avty,omitempty"`          // optional
	AvailabilityMode          string            `json:"avty_mode,omitempty"`     // optional
	AvailabilityTemplate      Template          `json:"avty_tpl,omitempty"`      // optional
	AvailabilityTopic         string            `json:"avty_t,omitempty"`        // optional
	DefaultEntityId           string            `json:"def_ent_id,omitempty"`    // optional
	Device                    map[string]Device `json:"dev,omitempty"`           // optional
	DeviceClass               DeviceClass       `json:"dev_cla,omitempty"`       // optional
	EnabledByDefault          bool              `json:"en,omitempty"`            // optional
	Encoding                  string            `json:"e,omitempty"`             // optional
	EntityCategory            string            `json:"ent_cat,omitempty"`       // optional
	EntityPicture             string            `json:"ent_pic,omitempty"`       // optional
	ExpireAfter               int               `json:"exp_aft,omitempty"`       // optional
	ForceUpdate               bool              `json:"frc_upd,omitempty"`       // optional
	Icon                      string            `json:"ic,omitempty"`            // optional // for Entity
	JsonAttributesTemplate    Template          `json:"json_attr_tpl,omitempty"` // optional
	JsonAttributesTopic       string            `json:"json_attr_t,omitempty"`   // optional
	LastResetValueTemplate    Template          `json:"lrst_val_tpl,omitempty"`  // optional
	Name                      string            `json:"name,omitempty"`          // optional
	Options                   []string          `json:"ops,omitempty"`           // optional
	PayloadAvailable          string            `json:"pl_avail,omitempty"`      // optional
	PayloadNotAvailable       string            `json:"pl_not_avail,omitempty"`  // optional
	Platform                  Platform          `json:"p,omitempty"`             // required  // NOTE: Must be 'sensor'. Only allowed and required in MQTT auto discovery device messages.
	Qos                       int               `json:"qos,omitempty"`           // optional
	StateClass                StateClass        `json:"stat_cla,omitempty"`      // optional
	StateTopic                string            `json:"stat_t,omitempty"`        // required
	SuggestedDisplayPrecision int               `json:"sug_dsp_prc,omitempty"`   // optional
	UniqueID                  string            `json:"uniq_id,omitempty"`       // optional
	UnitOfMeasurement         string            `json:"unit_of_meas,omitempty"`  // optional
	ValueTemplate             string            `json:"val_tpl,omitempty"`       // optional
}
