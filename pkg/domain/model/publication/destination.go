package publication

// Destination identifies a delivery destination.
type Destination string

const (
	// DestinationMQTT identifies MQTT delivery.
	DestinationMQTT Destination = "mqtt"
	// DestinationInfluxDB identifies InfluxDB delivery.
	DestinationInfluxDB Destination = "influxdb"
	// DestinationPostgres identifies PostgreSQL delivery.
	DestinationPostgres Destination = "postgres"
)
