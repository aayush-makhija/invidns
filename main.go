package invidns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider wraps the provider implementation as a Caddy module.
type provider struct {
	*Provider
}

func init() {
	caddy.RegisterModule(provider{})
}

// CaddyModule returns the Caddy module information.
func (provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.invidns",
		New: func() caddy.Module { return &provider{new(Provider)} },
	}
}

// Provision sets up the provider by resolving placeholders and sending the request.
// Implements caddy.Provisioner.
func (p *provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.URL = repl.ReplaceAll(p.Provider.URL, "")
	p.Provider.APIToken = repl.ReplaceAll(p.Provider.APIToken, "")
	p.Provider.OverrideDomain = repl.ReplaceAll(p.Provider.OverrideDomain, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	invidns [<api_token>] {
//	    url <requestbin_url>
//	    username <username>
//	    password <password>
//	}
func (p *provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {

			case "url":
				if !d.NextArg() {
					return d.ArgErr()
				}
				if p.Provider.URL != "" {
					return d.Err("URL already set")
				}
				p.Provider.URL = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_token":
				if !d.NextArg() {
					return d.ArgErr()
				}
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				p.Provider.APIToken = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			case "override_domain":
				if !d.NextArg() {
					return d.ArgErr()
				}
				if p.Provider.OverrideDomain != "" {
					return d.Err("Override domain already set")
				}
				p.Provider.OverrideDomain = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*provider)(nil)
	_ caddy.Provisioner     = (*provider)(nil)
)
