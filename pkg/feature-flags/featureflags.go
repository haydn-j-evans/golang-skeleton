package featureflags

import "sync"

type featureFlags struct {
	FeatureFlagMap map[string]feature `json:"featureflags"`
	mutex          sync.RWMutex
}

type feature struct {
	code            string `json:"code"`
	enabled         bool   `json:"enabled"`
	frontendRelated bool   `json:"frontendRelated"`
}
