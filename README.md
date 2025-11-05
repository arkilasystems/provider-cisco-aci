# Crossplane Provider for Cisco ACI

`provider-cisco-aci` is a [Crossplane](https://crossplane.io/) provider
that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for Cisco ACI
(Application Centric Infrastructure).

## Overview

This provider enables management of Cisco ACI infrastructure resources through Crossplane, wrapping the [CiscoDevNet/aci Terraform provider](https://github.com/CiscoDevNet/terraform-provider-aci).

### Supported Resources

The provider currently includes configurations for the following ACI resource groups:

- **Tenant Resources**: Tenants, Application Profiles, Application EPGs
- **Networking Resources**: VRFs, Bridge Domains, Subnets, L3 Outside, External Network Instance Profiles
- **Contract Resources**: Contracts, Contract Subjects, Filters, Filter Entries
- **Domain Resources**: VMM Domains, Physical Domains, VLAN Pools, Fabric Nodes, Pod Maintenance Groups

## Prerequisites

- Kubernetes cluster with Crossplane installed
- Access to a Cisco ACI APIC controller
- Terraform v1.5.7 (for schema generation)

## Getting Started

### Generate Provider Schema

Before building the provider, you need to generate the Terraform provider schema:

```console
make config/schema.json
```

Note: This requires access to the Terraform Registry. If you encounter network restrictions, ensure your environment can access `registry.terraform.io`.

## Developing

### Run Code Generation

After generating the schema, run the code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

### Configuration

Create a ProviderConfig with your ACI credentials. The credentials should be stored in a Kubernetes Secret:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: aci-credentials
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "url": "https://your-apic-url.com",
      "username": "admin",
      "password": "your-password",
      "insecure": "true"
    }
---
apiVersion: aci.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: aci-credentials
      namespace: crossplane-system
      key: credentials
```

### Build and Install

Build, push, and install:

```console
make all
```

Build binary only:

```console
make build
```

Run against a Kubernetes cluster:

```console
make run
```

## Example Usage

### Creating a Tenant

```yaml
apiVersion: tenant.aci.crossplane.io/v1alpha1
kind: Tenant
metadata:
  name: example-tenant
spec:
  forProvider:
    name: production-tenant
    description: "Production tenant managed by Crossplane"
  providerConfigRef:
    name: default
```

### Creating an Application Profile

```yaml
apiVersion: tenant.aci.crossplane.io/v1alpha1
kind: ApplicationProfile
metadata:
  name: example-app-profile
spec:
  forProvider:
    name: web-app
    tenantDn: "uni/tn-production-tenant"
    description: "Web application profile"
  providerConfigRef:
    name: default
```

## Provider Configuration

The provider uses the following configuration from the Makefile:

- **Terraform Version**: 1.5.7
- **Provider Source**: CiscoDevNet/aci
- **Provider Version**: 2.17.0

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/arkilasystems/provider-cisco-aci/issues).
