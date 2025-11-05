package domain

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the domain group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("aci_vmm_domain", func(r *ujconfig.Resource) {
		r.Kind = "VMMDomain"
		r.ShortGroup = "domain"
	})

	p.AddResourceConfigurator("aci_physical_domain", func(r *ujconfig.Resource) {
		r.Kind = "PhysicalDomain"
		r.ShortGroup = "domain"
	})

	p.AddResourceConfigurator("aci_vlan_pool", func(r *ujconfig.Resource) {
		r.Kind = "VLANPool"
		r.ShortGroup = "domain"
	})

	p.AddResourceConfigurator("aci_fabric_node", func(r *ujconfig.Resource) {
		r.Kind = "FabricNode"
		r.ShortGroup = "domain"
	})

	p.AddResourceConfigurator("aci_pod_maintenance_group", func(r *ujconfig.Resource) {
		r.Kind = "PodMaintenanceGroup"
		r.ShortGroup = "domain"
	})
}
