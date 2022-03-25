package glesys

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsglesys "github.com/libdns/glesys"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnsglesys.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.glesys",
		New: func() caddy.Module { return &Provider{new(libdnsglesys.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()

	p.Provider.ApiKey = repl.ReplaceAll(p.Provider.ApiKey, "")
	p.Provider.Project = repl.ReplaceAll(p.Provider.Project, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// glesys [<project> api_key] {
//     project <project/user>
//     api_key <api_key>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.Project = d.Val()
		}
		if d.NextArg() {
			p.Provider.ApiKey = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "project":
				if p.Provider.Project != "" {
					return d.Err("Project already set")
				}
				if d.NextArg() {
					p.Provider.Project = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_key":
				if p.Provider.ApiKey != "" {
					return d.Err("API-Key already set")
				}
				if d.NextArg() {
					p.Provider.ApiKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Project == "" {
		return d.Err("missing Project")
	}
	if p.Provider.ApiKey == "" {
		return d.Err("missing API-Key")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
