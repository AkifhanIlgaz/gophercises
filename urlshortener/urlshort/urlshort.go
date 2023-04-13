package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type pathUrl struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}

	var pathUrls []pathUrl

	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := map[string]string{}

	for _, pathUrl := range pathUrls {
		pathsToUrls[pathUrl.Path] = pathUrl.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type pathUrl struct {
		Path string `json:"path"`
		Url  string `json:"url"`
	}

	var pathUrls []pathUrl

	err := yaml.Unmarshal(json, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := map[string]string{}

	for _, pathUrl := range pathUrls {
		pathsToUrls[pathUrl.Path] = pathUrl.Url
	}

	return MapHandler(pathsToUrls, fallback), nil
}
