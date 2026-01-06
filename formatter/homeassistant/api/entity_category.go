package api

// https://developers.home-assistant.io/docs/core/entity/#registry-properties

type EntityCategory string

const (
	ConfigEntityCategory     EntityCategory = "config"
	DiagnosticEntityCategory EntityCategory = "diagnostic"
)
