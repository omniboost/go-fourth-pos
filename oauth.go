package fourth_pos

import (
	"net/url"
	"strings"
	"time"

	"github.com/joefitzgerald/passwordcredentials"
	"golang.org/x/oauth2"
)

const (
	scope                = ""
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://api.fourth.com/api/eposgateway/token",
				TokenURL:  "https://api.fourth.com/api/eposgateway/token",
				AuthStyle: oauth2.AuthStyleAutoDetect,
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	c.Config.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/token",
		TokenURL: baseURL.String() + "/token",
	}
}

type Oauth2PasswordConfig struct {
	passwordcredentials.Config
}

func NewOauth2PasswordConfig() *Oauth2PasswordConfig {
	config := &Oauth2PasswordConfig{
		Config: passwordcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			Username:     "",
			Password:     "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://api.fourth.com/api/eposgateway/token",
				TokenURL: "https://api.fourth.com/api/eposgateway/token",
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2PasswordConfig) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	c.Config.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/token",
		TokenURL: baseURL.String() + "/token",
	}
}
