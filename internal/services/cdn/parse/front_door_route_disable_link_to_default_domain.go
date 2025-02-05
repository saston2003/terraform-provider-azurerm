package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type FrontDoorRouteDisableLinkToDefaultDomainId struct {
	SubscriptionId                 string
	ResourceGroup                  string
	ProfileName                    string
	AfdEndpointName                string
	RouteName                      string
	DisableLinkToDefaultDomainName string
}

func NewFrontDoorRouteDisableLinkToDefaultDomainID(subscriptionId, resourceGroup, profileName, afdEndpointName, routeName, disableLinkToDefaultDomainName string) FrontDoorRouteDisableLinkToDefaultDomainId {
	return FrontDoorRouteDisableLinkToDefaultDomainId{
		SubscriptionId:                 subscriptionId,
		ResourceGroup:                  resourceGroup,
		ProfileName:                    profileName,
		AfdEndpointName:                afdEndpointName,
		RouteName:                      routeName,
		DisableLinkToDefaultDomainName: disableLinkToDefaultDomainName,
	}
}

func (id FrontDoorRouteDisableLinkToDefaultDomainId) String() string {
	segments := []string{
		fmt.Sprintf("Disable Link To Default Domain Name %q", id.DisableLinkToDefaultDomainName),
		fmt.Sprintf("Route Name %q", id.RouteName),
		fmt.Sprintf("Afd Endpoint Name %q", id.AfdEndpointName),
		fmt.Sprintf("Profile Name %q", id.ProfileName),
		fmt.Sprintf("Resource Group %q", id.ResourceGroup),
	}
	segmentsStr := strings.Join(segments, " / ")
	return fmt.Sprintf("%s: (%s)", "Front Door Route Disable Link To Default Domain", segmentsStr)
}

func (id FrontDoorRouteDisableLinkToDefaultDomainId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Cdn/profiles/%s/afdEndpoints/%s/routes/%s/disableLinkToDefaultDomain/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.ProfileName, id.AfdEndpointName, id.RouteName, id.DisableLinkToDefaultDomainName)
}

// FrontDoorRouteDisableLinkToDefaultDomainID parses a FrontDoorRouteDisableLinkToDefaultDomain ID into an FrontDoorRouteDisableLinkToDefaultDomainId struct
func FrontDoorRouteDisableLinkToDefaultDomainID(input string) (*FrontDoorRouteDisableLinkToDefaultDomainId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := FrontDoorRouteDisableLinkToDefaultDomainId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	if resourceId.ProfileName, err = id.PopSegment("profiles"); err != nil {
		return nil, err
	}
	if resourceId.AfdEndpointName, err = id.PopSegment("afdEndpoints"); err != nil {
		return nil, err
	}
	if resourceId.RouteName, err = id.PopSegment("routes"); err != nil {
		return nil, err
	}
	if resourceId.DisableLinkToDefaultDomainName, err = id.PopSegment("disableLinkToDefaultDomain"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}

// FrontDoorRouteDisableLinkToDefaultDomainIDInsensitively parses an FrontDoorRouteDisableLinkToDefaultDomain ID into an FrontDoorRouteDisableLinkToDefaultDomainId struct, insensitively
// This should only be used to parse an ID for rewriting, the FrontDoorRouteDisableLinkToDefaultDomainID
// method should be used instead for validation etc.
//
// Whilst this may seem strange, this enables Terraform have consistent casing
// which works around issues in Core, whilst handling broken API responses.
func FrontDoorRouteDisableLinkToDefaultDomainIDInsensitively(input string) (*FrontDoorRouteDisableLinkToDefaultDomainId, error) {
	id, err := resourceids.ParseAzureResourceID(input)
	if err != nil {
		return nil, err
	}

	resourceId := FrontDoorRouteDisableLinkToDefaultDomainId{
		SubscriptionId: id.SubscriptionID,
		ResourceGroup:  id.ResourceGroup,
	}

	if resourceId.SubscriptionId == "" {
		return nil, fmt.Errorf("ID was missing the 'subscriptions' element")
	}

	if resourceId.ResourceGroup == "" {
		return nil, fmt.Errorf("ID was missing the 'resourceGroups' element")
	}

	// find the correct casing for the 'profiles' segment
	profilesKey := "profiles"
	for key := range id.Path {
		if strings.EqualFold(key, profilesKey) {
			profilesKey = key
			break
		}
	}
	if resourceId.ProfileName, err = id.PopSegment(profilesKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'afdEndpoints' segment
	afdEndpointsKey := "afdEndpoints"
	for key := range id.Path {
		if strings.EqualFold(key, afdEndpointsKey) {
			afdEndpointsKey = key
			break
		}
	}
	if resourceId.AfdEndpointName, err = id.PopSegment(afdEndpointsKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'routes' segment
	routesKey := "routes"
	for key := range id.Path {
		if strings.EqualFold(key, routesKey) {
			routesKey = key
			break
		}
	}
	if resourceId.RouteName, err = id.PopSegment(routesKey); err != nil {
		return nil, err
	}

	// find the correct casing for the 'disableLinkToDefaultDomain' segment
	disableLinkToDefaultDomainKey := "disableLinkToDefaultDomain"
	for key := range id.Path {
		if strings.EqualFold(key, disableLinkToDefaultDomainKey) {
			disableLinkToDefaultDomainKey = key
			break
		}
	}
	if resourceId.DisableLinkToDefaultDomainName, err = id.PopSegment(disableLinkToDefaultDomainKey); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &resourceId, nil
}
