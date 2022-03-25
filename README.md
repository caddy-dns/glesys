Glesys module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). 

It can be used to manage DNS records with [Glesys](https://glesys.se).

## Caddy module name

```
dns.providers.glesys
```

## Config examples

### Caddyfile

```Caddyfile
somehost.example.org
tls {
        issuer acme {
                email "<your email for acme notifications>"
                dns glesys {
                        project "YOUR_GLESYS_PROJECT/USER"
                        api_key "YOUR_GLESYS_API_KEY"
                }
        }
}
respond "Hello, world!"
```

### JSON

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "glesys",
				"project": "YOUR_GLESYS_PROJECT/USER",
				"api_key": "YOUR_GLESYS_API_KEY"
			}
		}
	}
}
```

