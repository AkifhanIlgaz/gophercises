module github.com/AkifhanIlgaz/gophercises

go 1.20

require gopkg.in/yaml.v3 v3.0.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

require (
	github.com/boltdb/bolt v1.3.1
	github.com/lib/pq v1.10.9
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.7.0
	golang.org/x/net v0.9.0
)

replace urlshort => ./urlshortener/handler/
