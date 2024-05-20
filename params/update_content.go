package params

import (
	"strconv"
	"strings"
	"time"
)

type UpdateContentParams struct {
	Attribute      string
	AttributeValue string
}

func (p UpdateContentParams) Body() map[string]interface{} {
	return map[string]interface{}{
		"attribute":      p.Attribute,
		"attributeValue": p.AttributeValue,
	}
}

type UpdateContentOption func(*UpdateContentParams)

func WithName(name string) UpdateContentOption {
	return func(params *UpdateContentParams) {
		params.Attribute = "name"
		params.AttributeValue = name
	}
}

func WithDescription(description string) UpdateContentOption {
	return func(params *UpdateContentParams) {
		params.Attribute = "description"
		params.AttributeValue = description
	}
}

func WithTags(tags []string) UpdateContentOption {
	return func(params *UpdateContentParams) {
		params.Attribute = "tags"
		params.AttributeValue = strings.Join(tags, ",")
	}
}

func WithPublic(public bool) UpdateContentOption {
	return func(params *UpdateContentParams) {
		params.Attribute = "public"
		params.AttributeValue = "true"
		if !public {
			params.AttributeValue = "false"
		}
	}
}

func WithExpiry(expiry time.Time) UpdateContentOption {
	return func(params *UpdateContentParams) {
		expiryIntStr := strconv.FormatInt(expiry.Unix(), 10)
		params.Attribute = "expiry"
		params.AttributeValue = expiryIntStr
	}
}

func WithPassword(password string) UpdateContentOption {
	return func(params *UpdateContentParams) {
		params.Attribute = "password"
		params.AttributeValue = password
	}
}
