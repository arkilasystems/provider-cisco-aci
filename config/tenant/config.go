package tenant

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the tenant group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("aci_tenant", func(r *ujconfig.Resource) {
		r.Kind = "Tenant"
		r.ShortGroup = "tenant"
	})

	p.AddResourceConfigurator("aci_application_profile", func(r *ujconfig.Resource) {
		r.Kind = "ApplicationProfile"
		r.ShortGroup = "tenant"
	})

	p.AddResourceConfigurator("aci_application_epg", func(r *ujconfig.Resource) {
		r.Kind = "ApplicationEPG"
		r.ShortGroup = "tenant"
	})
}
