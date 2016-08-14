// Copyright 2016 Travelydia, Inc. All rights reserved.

package config

import (
	"github.com/gotravelydia/platform/log"

	"github.com/go-ini/ini"
)

const DefaultEnv = "staging"
const ConfigPath = ".cfg/config.cfg"

var ServiceConfig *Config

type Config struct {
	Env  string
	file *ini.File
}

func init() {
	file, err := ini.Load(ConfigPath)
	if err != nil {
		log.Error(err)
	}

	ServiceConfig = &Config{
		Env:  DefaultEnv,
		file: file,
	}
}

func (c *Config) GetString(key, defaultValue string) string {
	if !c.file.Section(c.Env).HasKey(key) {
		return defaultValue
	}

	return c.file.Section(c.Env).Key(key).String()
}

func (c *Config) GetInt(key string, defaultValue int) int {
	if !c.file.Section(c.Env).HasKey(key) {
		return defaultValue
	}

	value, err := c.file.Section(c.Env).Key(key).Int()
	if err != nil {
		log.Error(err)
		return defaultValue
	}

	return value
}
