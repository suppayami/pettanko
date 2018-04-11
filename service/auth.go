package service

import (
	"errors"

	"github.com/suppayami/authgo/strategy"
)

var (
	// LocalStrategy is an authentication strategy
	LocalStrategy *strategy.LocalStrategy
)

func init() {
	LocalStrategy, _ = strategy.NewLocalStrategy(
		nil,
		func(credentials strategy.LocalCredentials) error {
			if credentials.Username == "test" && credentials.Password == "test" {
				return nil
			}
			return errors.New("wrong username or password")
		},
	)
}
