package api

// Reference
// https://www.home-assistant.io/integrations/device_trigger.mqtt/#automation_type

type AutomationType string

const (
	TriggerAutomationType   AutomationType = "trigger"   // https://www.home-assistant.io/docs/automation/trigger/
	ConditionAutomationType AutomationType = "condition" // hhttps://www.home-assistant.io/docs/automation/condition/
	ActionAutomationType    AutomationType = "action"    // https://www.home-assistant.io/docs/automation/action/
)
