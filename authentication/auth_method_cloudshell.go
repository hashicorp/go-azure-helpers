package authentication

import (
	"fmt"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/cli"
	"github.com/hashicorp/go-multierror"
)

type cloudShellAuth struct {
	profile *azureCLIProfile
}

func (a cloudShellAuth) build(b Builder) (authMethod, error) {
	auth := cloudShellAuth{
		profile: &azureCLIProfile{
			clientId:       b.ClientID,
			environment:    b.Environment,
			subscriptionId: b.SubscriptionID,
			tenantId:       b.TenantID,
		},
	}
	profilePath, err := cli.ProfilePath()
	if err != nil {
		return nil, fmt.Errorf("Error loading the Azure CLI Profile Path from CloudShell: %+v", err)
	}

	profile, err := cli.LoadProfile(profilePath)
	if err != nil {
		return nil, fmt.Errorf("Azure CLI Authorization Profile was not found. Please re-launch your CloudShell.")
	}

	auth.profile.profile = profile

	err = auth.profile.populateFields()
	if err != nil {
		return nil, err
	}

	err = auth.profile.populateClientIdAndAccessToken()
	if err != nil {
		return nil, fmt.Errorf("Error populating CloudShell Access Tokens from the Azure CLI: %+v", err)
	}

	return auth, nil
}

func (a cloudShellAuth) isApplicable(b Builder) bool {
	return b.SupportsCloudShell && os.Getenv("ACC_CLOUD") != ""
}

func (a cloudShellAuth) getAuthorizationToken(oauthConfig *adal.OAuthConfig, endpoint string) (*autorest.BearerAuthorizer, error) {
	// load the refreshed tokens from the CloudShell Azure CLI credentials
	err := a.profile.populateClientIdAndAccessToken()
	if err != nil {
		return nil, fmt.Errorf("Error loading the refreshed CloudShell tokens: %+v", err)
	}

	spt, err := adal.NewServicePrincipalTokenFromManualToken(*oauthConfig, a.profile.clientId, endpoint, *a.profile.accessToken)
	if err != nil {
		return nil, err
	}

	auth := autorest.NewBearerAuthorizer(spt)
	return auth, nil
}

func (a cloudShellAuth) name() string {
	return "Parsing credentials from CloudShell"
}

func (a cloudShellAuth) populateConfig(c *Config) error {
	c.ClientID = a.profile.clientId
	c.Environment = a.profile.environment
	c.SubscriptionID = a.profile.subscriptionId
	c.TenantID = a.profile.tenantId
	return nil
}

func (a cloudShellAuth) validate() error {
	var err *multierror.Error

	errorMessageFmt := "A %s was not found in your Azure CLI Credentials.\n\nPlease re-launch your CloudShell."

	if a.profile == nil {
		return fmt.Errorf("Azure CLI Profile is nil - this is an internal error and should be reported.")
	}

	if a.profile.accessToken == nil {
		err = multierror.Append(err, fmt.Errorf(errorMessageFmt, "Access Token"))
	}

	if a.profile.clientId == "" {
		err = multierror.Append(err, fmt.Errorf(errorMessageFmt, "Client ID"))
	}

	if a.profile.subscriptionId == "" {
		err = multierror.Append(err, fmt.Errorf(errorMessageFmt, "Subscription ID"))
	}

	if a.profile.tenantId == "" {
		err = multierror.Append(err, fmt.Errorf(errorMessageFmt, "Tenant ID"))
	}

	return err.ErrorOrNil()
}
