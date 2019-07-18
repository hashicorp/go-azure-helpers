package authentication

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

type authMethod interface {
	build(b Builder) (authMethod, error)

	isApplicable(b Builder) bool

	getAuthorizationToken(sender autorest.Sender, oauthConfig *adal.OAuthConfig, endpoint string) (*autorest.BearerAuthorizer, error)

	getMultiTenantAuthorizationToken(sender autorest.Sender, oauthConfig *adal.MultiTenantOAuthConfig, endpoint string) (*autorest.MultiTenantServicePrincipalTokenAuthorizer, error)

	name() string

	populateConfig(c *Config) error

	validate() error
}
