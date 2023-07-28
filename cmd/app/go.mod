module github.com/jmwample/nested_mods/cmd/app

go 1.20

replace github.com/jmwample/nested_modules/mods/server => ../../mods/server

require (
	github.com/jmwample/nested_modules/mods/server v0.0.0
	gitlab.com/yawning/obfs4.git v0.0.0-20230723031256-efdc692691f7
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/dchest/siphash v1.2.3 // indirect
	gitlab.com/yawning/edwards25519-extra.git v0.0.0-20220726154925-def713fd18e4 // indirect
	gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/goptlib v1.4.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
