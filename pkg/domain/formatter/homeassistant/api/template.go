package api

// Example value template:
// With given payload:
// { "state": "ON", "temperature": 21.902, "humidity": null }
// JSON
// Template {{ value_json.temperature | round(1) }} renders to 21.9.
// Template {{ value_json.humidity }} renders to None.

type Template string
