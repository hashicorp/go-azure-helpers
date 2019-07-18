package authentication

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/hashicorp/go-multierror"
)

type servicePrincipalClientSecretMultitenantAuth struct {
	clientId       string
	clientSecret   string
	subscriptionId string
	tenantId       string
}

func (a servicePrincipalClientSecretMultitenantAuth) build(b Builder) (authMethod, error) {
	method := servicePrincipalClientSecretMultitenantAuth{
		clientId:       b.ClientID,
		clientSecret:   b.ClientSecret,
		subscriptionId: b.SubscriptionID,
		tenantId:       b.TenantID,
	}
	return method, nil
}

func (a servicePrincipalClientSecretMultitenantAuth) isApplicable(b Builder) bool {
	return b.SupportsClientSecretAuth && b.ClientSecret != ""
}

func (a servicePrincipalClientSecretMultitenantAuth) name() string {
	return "Service Principal / Client Secret"
}

func (a servicePrincipalClientSecretMultitenantAuth) getAuthorizationToken(sender autorest.Sender, oauth *MultiOAuth, endpoint string) (autorest.Authorizer, error) {
	spt, err := adal.NewMultiTenantServicePrincipalToken(*oauth.MultiTenantOauth, a.clientId, a.clientSecret, endpoint)
	if err != nil {
		return nil, err
	}

	spt.PrimaryToken.SetSender(sender)
	for _, t := range spt.AuxiliaryTokens {
		t.SetSender(sender)
	}

	auth := autorest.NewMultiTenantServicePrincipalTokenAuthorizer(spt)
	return auth, nil
}

func (a servicePrincipalClientSecretMultitenantAuth) populateConfig(c *Config) error {
	c.AuthenticatedAsAServicePrincipal = true
	return nil
}

func (a servicePrincipalClientSecretMultitenantAuth) validate() error {
	var err *multierror.Error

	fmtErrorMessage := "A %s must be configured when authenticating as a Service Principal using a Client Secret."

	if a.subscriptionId == "" {
		err = multierror.Append(err, fmt.Errorf(fmtErrorMessage, "Subscription ID"))
	}
	if a.clientId == "" {
		err = multierror.Append(err, fmt.Errorf(fmtErrorMessage, "Client ID"))
	}
	if a.clientSecret == "" {
		err = multierror.Append(err, fmt.Errorf(fmtErrorMessage, "Client Secret"))
	}
	if a.tenantId == "" {
		err = multierror.Append(err, fmt.Errorf(fmtErrorMessage, "Tenant ID"))
	}

	return err.ErrorOrNil()
}
