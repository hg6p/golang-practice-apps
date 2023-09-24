package urlshort

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml map[string]string, fallback http.Handler) (http.HandlerFunc, error) {

	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := yml[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}, nil
}

type T struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func ParseYaml(pathToFile string) []T {

	// Read YAML data from a file
	yamlData, errRf := os.ReadFile(pathToFile)
	if errRf != nil {
		log.Fatalf("error reading YAML file: %v", errRf)
	}
	t := []T{}
	err := yaml.Unmarshal([]byte(yamlData), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	return t

}

func CreateMap(yamlData []T) map[string]string {
	yamlMap := make(map[string]string)
	for _, item := range yamlData {
		yamlMap[item.Path] = item.URL
	}

	return yamlMap
}
