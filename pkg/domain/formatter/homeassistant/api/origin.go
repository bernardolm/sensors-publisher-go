package api

// https://www.home-assistant.io/integrations/mqtt/#adding-information-about-the-origin-of-a-discovery-message
// https://www.home-assistant.io/integrations/mqtt/#discovery-payload
// https://www.home-assistant.io/integrations/mqtt/#supported-abbreviations-in-mqtt-discovery-messages

type Origin struct {
	Name            string `json:"name,omitempty"` // required
	SoftwareVersion string `json:"sw,omitempty"`   // optional
	SupportURL      string `json:"url,omitempty"`  // optional
}
