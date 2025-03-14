package fourth_pos_test

import (
	"context"
	"os"
	"testing"

	fourth_pos "github.com/omniboost/go-fourth-pos"
)

var (
	client *fourth_pos.Client
)

func TestMain(m *testing.M) {
	username := os.Getenv("OAUTH_USERNAME")
	password := os.Getenv("OAUTH_PASSWORD")
	// tokenURL := os.Getenv("OAUTH_TOKEN_URL")

	// Default oausth2 flow
	oauthConfig := fourth_pos.NewOauth2PasswordConfig()
	oauthConfig.Username = username
	oauthConfig.Password = password

	// set alternative token url
	// if tokenURL != "" {
	// 	oauthConfig.TokenURL = tokenURL
	// }

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client = fourth_pos.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
