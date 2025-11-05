package networking

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the networking group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("aci_vrf", func(r *ujconfig.Resource) {
		r.Kind = "VRF"
		r.ShortGroup = "networking"
	})

	p.AddResourceConfigurator("aci_bridge_domain", func(r *ujconfig.Resource) {
		r.Kind = "BridgeDomain"
		r.ShortGroup = "networking"
	})

	p.AddResourceConfigurator("aci_subnet", func(r *ujconfig.Resource) {
		r.Kind = "Subnet"
		r.ShortGroup = "networking"
	})

	p.AddResourceConfigurator("aci_l3_outside", func(r *ujconfig.Resource) {
		r.Kind = "L3Outside"
		r.ShortGroup = "networking"
	})

	p.AddResourceConfigurator("aci_external_network_instance_profile", func(r *ujconfig.Resource) {
		r.Kind = "ExternalNetworkInstanceProfile"
		r.ShortGroup = "networking"
	})
}
