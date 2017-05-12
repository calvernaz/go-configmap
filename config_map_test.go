package configmap

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Get
func Test_ShouldReturnValue(t *testing.T) {
	cfg := &ConfigMap{"b": "c"}
	v, ok := cfg.Get("b")
	assert.Equal(t, "c", v, "Should return \"c\"")
	assert.True(t, ok, "Should be true")
}

func Test_ShouldReturnValueIsNotPresent(t *testing.T) {
	cfg := &ConfigMap{"b": "c"}
	v, ok := cfg.Get("d")
	assert.Nil(t, v, "Should be nil")
	assert.False(t, ok, "Should be false")
}

func Test_ShouldReturnValueIsNotPresentWhenIsEmpty(t *testing.T) {
	cfg := &ConfigMap{"b": ""}
	v, ok := cfg.Get("b")
	assert.Nil(t, v, "Should be nil")
	assert.False(t, ok, "Should be false")
}

// GetOrDefault
func Test_ShouldReturnExistentValue(t *testing.T) {
	cfg := &ConfigMap{"b": "c"}
	v, _ := cfg.GetOrDefault("b", "z")
	assert.Equal(t, "c", v, "Should return \"z\"")
}

func Test_ShouldThrowErrorIfKeyDoesNotExistAndDefaultValueIsInvalid(t *testing.T) {
	cfg := &ConfigMap{"d": "e"}
	_, err := cfg.GetOrDefault("z", nil)
	assert.EqualError(t, err, "Provided default value is invalid", "An error was expected")
}

func Test_ShouldReturnDefaultValueIfKeyDoesNotExist(t *testing.T) {
	cfg := &ConfigMap{"d": "e"}
	v, err := cfg.GetOrDefault("g", "h")
	assert.Nil(t, err, "No error should be thrown")
	assert.Equal(t, "h", v, "Should return fallback value \"h\"")

	dfv, ok := cfg.Get("g")
	assert.Equal(t, dfv, "h", "Should add variable in MapConfig")
	assert.True(t, ok, "Should be true")
}

func Test_ShouldThrowErrorIfKeyExistButDefaultValueIsInvalid(t *testing.T) {
	cfg := &ConfigMap{"a": ""}
	_, err := cfg.GetOrDefault("a", nil)
	assert.EqualError(t, err, "Provided default value is invalid", "An error was expected")
}

func Test_ShouldFallbackToDefaultValueWhenEmpty(t *testing.T) {
	cfg := &ConfigMap{"a": ""}
	v, _ := cfg.GetOrDefault("a", "b")
	assert.Equal(t, "b", v, "Should fallback to \"b\"")
}

// GetEnvOrDefault
func Test_ShouldReturnEnvValue(t *testing.T) {
	cfg := &ConfigMap{}
	os.Setenv("env", "variable")
	v, _ := cfg.GetEnvOrDefault("env", "default")
	assert.Equal(t, "variable", v, "Should return environment variable")

	ev, ok := cfg.Get("env")
	assert.Equal(t, ev, "variable", "Should add environment variable in MapConfig")
	assert.True(t, ok, "Should be true")
	os.Unsetenv("env")
}

func Test_ShouldReturnFallbackValueIfEnvValueNotPresent(t *testing.T) {
	cfg := &ConfigMap{}
	v, _ := cfg.GetEnvOrDefault("env", "default")
	assert.Equal(t, "default", v, "Should return environment variable")

	ev, ok := cfg.Get("env")
	assert.Equal(t, ev, "default", "Should add environment variable in MapConfig")
	assert.True(t, ok, "Should be true")

}

func Test_ShouldReturnErrorIfEnvValueIsNotPresentAndDefaultIsInvalid(t *testing.T) {
	cfg := &ConfigMap{}
	v, err := cfg.GetEnvOrDefault("env", nil)

	assert.Error(t, err, "Should throw an error")
	assert.Nil(t, v, "Return value should be nil")
}

// Random values
func Test_ShouldReturnSliceOfStrings(t *testing.T) {
	cfg := &ConfigMap{"ss": []string{"one", "two", "three"}}
	var v interface{}
	var err error
	var ok bool

	v, err = cfg.GetOrDefault("ss", []string{})
	assert.Len(t, v, 3, "Slice size should be 3")
	assert.Nil(t, err, "Should have no errors")

	v, ok = cfg.Get("ss")
	assert.Len(t, v, 3, "Slice size should be 3")
	assert.True(t, ok, "Should return ok")

	v, err = cfg.GetOrDefault("zz", []string{})
	assert.Error(t, err, "Should throw an error")
	assert.Nil(t, v, "Should return value nil")
}

func Test_ShouldReturnBool(t *testing.T) {
	cfg := &ConfigMap{"bt": true}
	var v interface{}

	v, _ = cfg.GetOrDefault("bt", nil)
	assert.Equal(t, true, v, "Should return true")

	v, _ = cfg.GetOrDefault("bnx", false)
	assert.Equal(t, false, v, "Should return false")

	v, _ = cfg.Get("bt")
	assert.Equal(t, true, v, "Should be true")

	var ok bool
	v, ok = cfg.Get("bx")
	assert.Nil(t, v, "Should be nil")
	assert.False(t, ok, "Should be false")
}
