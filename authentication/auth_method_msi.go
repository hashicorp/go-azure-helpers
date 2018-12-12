package authentication

import (
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/hashicorp/go-multierror"
)

type managedServiceIdentityAuth struct {
	endpoint string
	clientID string
}

func (a managedServiceIdentityAuth) build(b Builder) (authMethod, error) {
	endpoint := b.MsiEndpoint
	if endpoint == "" {
		msiEndpoint, err := adal.GetMSIVMEndpoint()
		if err != nil {
			return nil, fmt.Errorf("Error determining MSI Endpoint: ensure the VM has MSI enabled, or configure the MSI Endpoint. Error: %s", err)
		}
		endpoint = msiEndpoint
	}

	log.Printf("[DEBUG] Using MSI endpoint %q", endpoint)

	auth := managedServiceIdentityAuth{
		endpoint: endpoint,
		clientID: b.ClientID,
	}
	return auth, nil
}

func (a managedServiceIdentityAuth) isApplicable(b Builder) bool {
	return b.SupportsManagedServiceIdentity
}

func (a managedServiceIdentityAuth) name() string {
	return "Managed Service Identity"
}

func (a managedServiceIdentityAuth) getAuthorizationToken(oauthConfig *adal.OAuthConfig, resource string) (*autorest.BearerAuthorizer, error) {
	log.Printf("[DEBUG] getAuthorizationToken using MSI endpoint %q", a.endpoint)
	log.Printf("[DEBUG] getAuthorizationToken using client_id %q", a.clientID)
	log.Printf("[DEBUG] Calling getAuthorizationToken for resource: %q", resource)

	if a.clientID == "" {
		spt, err := adal.NewServicePrincipalTokenFromMSI(a.endpoint, resource)
		if err != nil {
			return nil, err
		}
		auth := autorest.NewBearerAuthorizer(spt)
		return auth, nil
	} else {
		spt, err := adal.NewServicePrincipalTokenFromMSIWithUserAssignedID(a.endpoint, resource, a.clientID)
		if err != nil {
			return nil, fmt.Errorf("failed to get oauth token from MSI for user assigned identity: %v", err)
		}
		
		auth := autorest.NewBearerAuthorizer(spt)
		return auth, nil
	}
}

func (a managedServiceIdentityAuth) populateConfig(c *Config) error {
	// nothing to populate back
	return nil
}

func (a managedServiceIdentityAuth) validate() error {
	var err *multierror.Error

	if a.endpoint == "" {
		err = multierror.Append(err, fmt.Errorf("An MSI Endpoint must be configured"))
	}

	return err.ErrorOrNil()
}
