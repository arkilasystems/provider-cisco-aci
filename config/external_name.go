package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Tenant Resources
	"aci_tenant":               config.IdentifierFromProvider,
	"aci_application_profile":  config.IdentifierFromProvider,
	"aci_application_epg":      config.IdentifierFromProvider,

	// Networking Resources
	"aci_vrf":                  config.IdentifierFromProvider,
	"aci_bridge_domain":        config.IdentifierFromProvider,
	"aci_subnet":               config.IdentifierFromProvider,

	// Contract Resources
	"aci_contract":             config.IdentifierFromProvider,
	"aci_contract_subject":     config.IdentifierFromProvider,
	"aci_filter":               config.IdentifierFromProvider,
	"aci_filter_entry":         config.IdentifierFromProvider,

	// Physical Resources
	"aci_l3_outside":           config.IdentifierFromProvider,
	"aci_external_network_instance_profile": config.IdentifierFromProvider,

	// Domain and Pool Resources
	"aci_vmm_domain":           config.IdentifierFromProvider,
	"aci_physical_domain":      config.IdentifierFromProvider,
	"aci_vlan_pool":            config.IdentifierFromProvider,

	// Fabric Resources
	"aci_fabric_node":          config.IdentifierFromProvider,
	"aci_pod_maintenance_group": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
