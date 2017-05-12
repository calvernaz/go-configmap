package configmap

import (
	"errors"
	"os"
	"reflect"
)

// ConfigMap is a map that holds the configuration
// and implements some helper methods
type ConfigMap map[string]interface{}

// Get returns the configuration value
func (c ConfigMap) Get(key string) (interface{}, bool) {
	v, ok := c[key]
	if ok {
		if notEmptyOrNil(v) {
			return v, ok
		}
	}
	return nil, false
}

// GetOrDefault gets a configuration value if present, otherwise
// fallbacks to the provided value. In case the default provided is invalid
// an error is return
func (c ConfigMap) GetOrDefault(key string, defaultValue interface{}) (interface{}, error) {
	v, ok := c[key]
	if ok {
		if notEmptyOrNil(v) {
			c[key] = v
			return v, nil
		}

		if notEmptyOrNil(defaultValue) {
			c[key] = defaultValue
			return defaultValue, nil
		}
		return nil, errors.New("Provided default value is invalid")
	}

	if notEmptyOrNil(defaultValue) {
		c[key] = defaultValue
		return defaultValue, nil
	}
	return nil, errors.New("Provided default value is invalid")
}

// GetEnvOrDefault returns the environment variable if present
func (c ConfigMap) GetEnvOrDefault(key string, value interface{}) (interface{}, error) {
	env := os.Getenv(key)
	if env != "" {
		c[key] = env
		return env, nil
	}

	v, err := c.GetOrDefault(key, value)
	if err != nil {
		return nil, err
	}
	c[key] = value
	return v, nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func notEmptyOrNil(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.IsValid() && !isEmptyValue(v)
}

/*
func main() {

	cfg := ConfigMap{
		"a": "b", // covered
		"c": "",  // covered
		"d": 0,
		"e": []int{1, 2, 3},
		"f": []string{"one", "two", "three"},
		"g": func() { fmt.Println("Function Handler") },
	}

	fmt.Printf("ConfigMap size: %v\n", unsafe.Sizeof(cfg))
	fmt.Printf("&ConfigMap size: %v\n", unsafe.Sizeof(&cfg))
	fmt.Println(cfg)

	var v interface{}
	var b bool

	v, b = cfg.GetOrDefault("c", "d")
	fmt.Printf("map[c] exists (should be false) %v\n", b)
	fmt.Printf("Should map[c]=d %v\n", v)
}
*/
