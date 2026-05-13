package config

import "errors"

var (
	ErrRuleNotFound          = errors.New("ward: config: rule not found in registry")
	ErrRuleAlreadyRegistered = errors.New("ward: config: rule already registered and cannot be overwritten")
)
