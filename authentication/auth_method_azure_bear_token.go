package authentication

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-multierror"
	"github.com/manicminer/hamilton/environments"
)

type azureBearTokenAuth struct {
	bearToken string
}

func (a azureBearTokenAuth) build(b Builder) (authMethod, error) {
	auth := azureBearTokenAuth{
		bearToken: b.BearToken,
	}

	return auth, nil
}

func (a azureBearTokenAuth) isApplicable(b Builder) bool {
	return b.BearToken != ""
}

func (a azureBearTokenAuth) getADALToken(_ context.Context, _ autorest.Sender, oauthConfig *OAuthConfig, endpoint string) (autorest.Authorizer, error) {
	auth := autorest.NewBearerAuthorizer(&BearTokenProvider{
		token: a.bearToken,
	})
	return auth, nil
}

func (a azureBearTokenAuth) getMSALToken(ctx context.Context, _ environments.Api, sender autorest.Sender, oauthConfig *OAuthConfig, endpoint string) (autorest.Authorizer, error) {
	return a.getADALToken(ctx, sender, oauthConfig, endpoint)
}

func (a azureBearTokenAuth) name() string {
	return "Bear Token"
}

func (a azureBearTokenAuth) populateConfig(c *Config) error {
	return nil
}

func (a azureBearTokenAuth) validate() error {
	var err *multierror.Error

	if a.bearToken == "" {
		err = multierror.Append(err, fmt.Errorf("A Bear Token must be configured"))
	}

	return err.ErrorOrNil()
}

type BearTokenProvider struct {
	token string
}

// OAuthToken implements the OAuthTokenProvider interface.  It returns the current access token.
func (btp *BearTokenProvider) OAuthToken() string {
	return btp.token
}