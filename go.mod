module github.com/caddy-dns/glesys

go 1.16

replace github.com/libdns/glesys => ../libdns-glesys

require (
	github.com/caddyserver/caddy/v2 v2.4.1
	github.com/libdns/glesys v0.0.0-00010101000000-000000000000
)
