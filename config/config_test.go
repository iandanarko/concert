package config_test

import (
	"testing"

	"github.com/iandanarko/concert/config"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	t.Run("Failed: env file not found", func(t *testing.T) {
		t.Parallel()
		cfg, err := config.New("../test/fixture/env.notexists")
		require.Error(t, err)
		require.Nil(t, cfg)
	})

	t.Run("Failed: incomplete env file", func(t *testing.T) {
		t.Parallel()
		cfg, err := config.New("../test/fixture/env.invalid")
		require.Error(t, err)
		require.Nil(t, cfg)
	})

	t.Run("Failed because invalid file", func(t *testing.T) {
		t.Parallel()
		cfg, err := config.New("../test/fixture/env.valid")
		require.Nil(t, err)
		require.NotEmpty(t, cfg)
	})
}
