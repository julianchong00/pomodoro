package config

import (
	"time"

	"emperror.dev/errors"
)

type TimerConfig struct {
	WorkingDuration time.Duration
	RestingDuration time.Duration
}

func NewConfig(options ...func(*TimerConfig) error) (*TimerConfig, error) {
	config := &TimerConfig{}

	for _, option := range options {
		err := option(config)
		if err != nil {
			return nil, errors.Wrap(err, "failed to apply functional option to timer config")
		}
	}

	return config, nil
}

func Work(duration time.Duration) func(*TimerConfig) error {
	return func(tc *TimerConfig) error {
		tc.WorkingDuration = duration
		return nil
	}
}

func Rest(duration time.Duration) func(*TimerConfig) error {
	return func(tc *TimerConfig) error {
		tc.RestingDuration = duration
		return nil
	}
}
