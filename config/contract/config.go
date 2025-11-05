package contract

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures the contract group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("aci_contract", func(r *ujconfig.Resource) {
		r.Kind = "Contract"
		r.ShortGroup = "contract"
	})

	p.AddResourceConfigurator("aci_contract_subject", func(r *ujconfig.Resource) {
		r.Kind = "ContractSubject"
		r.ShortGroup = "contract"
	})

	p.AddResourceConfigurator("aci_filter", func(r *ujconfig.Resource) {
		r.Kind = "Filter"
		r.ShortGroup = "contract"
	})

	p.AddResourceConfigurator("aci_filter_entry", func(r *ujconfig.Resource) {
		r.Kind = "FilterEntry"
		r.ShortGroup = "contract"
	})
}
