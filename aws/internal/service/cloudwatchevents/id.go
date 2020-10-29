package cloudwatchevents

import (
	"fmt"
	"strings"
)

const DefaultEventBusName = "default"

const PermissionIDSeparator = "/"

func PermissionCreateID(eventBusName, statementID string) string {
	if eventBusName == "" || eventBusName == DefaultEventBusName {
		return statementID
	}
	return eventBusName + PermissionIDSeparator + statementID
}

func PermissionParseID(id string) (string, string, error) {
	parts := strings.Split(id, PermissionIDSeparator)
	if len(parts) == 1 && parts[0] != "" {
		return DefaultEventBusName, parts[0], nil
	}
	if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
		return parts[0], parts[1], nil
	}

	return "", "", fmt.Errorf("unexpected format for ID (%q), expected <event-bus-name>"+PermissionIDSeparator+"<statement-id> or <statement-id>", id)
}

const ruleIDSeparator = "/"

func RuleCreateID(eventBusName, ruleName string) string {
	if eventBusName == "" || eventBusName == DefaultEventBusName {
		return ruleName
	}
	return eventBusName + ruleIDSeparator + ruleName
}

func RuleParseID(id string) (string, string, error) {
	parts := strings.Split(id, ruleIDSeparator)
	if len(parts) == 1 && parts[0] != "" {
		return DefaultEventBusName, parts[0], nil
	}
	if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
		return parts[0], parts[1], nil
	}

	return "", "", fmt.Errorf("unexpected format for ID (%q), expected <event-bus-name>"+ruleIDSeparator+"<rule-name> or <rule-name>", id)
}
