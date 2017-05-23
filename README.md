[![Build Status](https://travis-ci.org/calvernaz/go-configmap.svg?branch=master)](https://travis-ci.org/calvernaz/go-configmap)
[![Coverage Status](https://coveralls.io/repos/github/calvernaz/go-configmap/badge.svg?branch=master)](https://coveralls.io/github/calvernaz/go-configmap?branch=master)

# go-configmap

Old habits die hard. The motivation is get configuration settings from your environment and if they don't exist provide a default.

## Install

	 go get github.com/calvernaz/go-configmap

## Use

### Get from environment, otherwise fallbacks to default

```go
cfg := &ConfigMap{}
v, _ := cfg.GetEnvOrDefault("env", "default")
fmt.Println(v) // default
```

### Get from configuration, otherwise fallbacks to default

```go
cfg := &ConfigMap{ "db": "file.db" }
v, _ := cfg.GetOrDefault("db", "data.db")
fmt.Println(v) // file.db
```

### Get from configuration

```go
cfg := &ConfigMap{ "debug": "false" }
v, ok := cfg.Get("debug")
fmt.Println(v) // false
fmt.Println(ok) // true
```

### Merge configuration

```go
cfg := &ConfigMap{}
cfg.MergeConfig(ConfigMap{
  "a":   "b",
  "foo": "bar",
})

assert.EqualValues(t, &ConfigMap{"a": "b", "foo": "bar"}, cfg)
```

```go
cfg := &ConfigMap{"foo": "bar"}
cfg.MergeConfig(ConfigMap{"foo": "barbar", "xyz": "abc"})
assert.EqualValues(t, &ConfigMap{"foo": "barbar", "xyz": "abc"}, cfg)
```
