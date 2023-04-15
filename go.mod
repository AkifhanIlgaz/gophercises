module github.com/AkifhanIlgaz/gophercises

go 1.20

require gopkg.in/yaml.v3 v3.0.1

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	golang.org/x/net v0.7.0 // indirect
)

require (
	github.com/PuerkitoBio/goquery v1.8.1
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

replace urlshort => ./urlshortener/handler/
