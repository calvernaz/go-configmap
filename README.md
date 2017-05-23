[![Build Status](https://travis-ci.org/calvernaz/go-configmap.svg?branch=master)](https://travis-ci.org/calvernaz/go-configmap)
[![Coverage Status](https://coveralls.io/repos/github/calvernaz/go-configmap/badge.svg?branch=master)](https://coveralls.io/github/calvernaz/go-configmap?branch=master)

# go-configmap

Old habits die hard. The motivation is get configuration settings from your environment and if they don't exist provide a default.

## Install

	 go get github.com/calvernaz/go-configmap

## Use

```go
cfg := &ConfigMap{}
v, _ := cfg.GetEnvOrDefault("env", "default")
fmt.Println(v) // default
```

## API

- Get(key string)
- GetOrDefault(key string, defaultValue interface{})
- GetEnvOrDefault(key string, value interface{})
- MergeConfig(config ConfigMap)
