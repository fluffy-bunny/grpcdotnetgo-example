package config

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
)

// ExampleConfig type
type ExampleConfig struct {
	Port             int             `mapstructure:"PORT"`
	Mode             string          `mapstructure:"MODE"`
	OIDCConfig       oidc.OIDCConfig `mapstructure:"OIDC_CONFIG"`
	EnableTransient2 bool            `mapstructure:"ENABLE_TRANSIENT_2"`
}

// Config type
type Config struct {
	ApplicationEnvironment    string        `json:"applicationEnvironment" mapstructure:"APPLICATION_ENVIRONMENT"`
	Example                   ExampleConfig `json:"example" mapstructure:"EXAMPLE"`
	ClaimsPrincipalMiddleware string        `json:"claimsPrincipalMiddleware" mapstructure:"CLAIMS_PRINCIPAL_MIDDLEWARE"`
}

// GetOIDCConfig gets config
func (c *Config) GetOIDCConfig() oidc.IOIDCConfig {
	return &c.Example.OIDCConfig
}

// ConfigDefaultJSON default yaml
var ConfigDefaultJSON = []byte(`
{
	"APPLICATION_ENVIRONMENT": "in-environment",
	"CLAIMS_PRINCIPAL_MIDDLEWARE": "development",
	"EXAMPLE": {
	  "ENABLE_TRANSIENT_2": true,
	  "PORT": 1111,
	  "OIDC_CONFIG": {
		"AUTHORITY": "https://in-environment/",
		"CRON_REFRESH_SCHEDULE": "@every 0h1m0s"
	  }
	}
  }
`)
