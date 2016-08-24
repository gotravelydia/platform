// Copyright 2016 Travelydia, Inc. All rights reserved.

package config

import (
	"errors"

	"github.com/gotravelydia/platform/log"

	"github.com/go-ini/ini"
)

const DefaultEnv = "staging"
const ConfigPath = "config.cfg"

var ServiceConfig *Config

type Config struct {
	Env  string
	file *ini.File
}

type RDSConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	PoolSize int
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

func (c *Config) GetString(key string) (string, error) {
	if !c.file.Section(c.Env).HasKey(key) {
		return "", errors.New("Missing key.")
	}

	return c.file.Section(c.Env).Key(key).String(), nil
}

func (c *Config) GetInt(key string) (int, error) {
	if !c.file.Section(c.Env).HasKey(key) {
		return 0, errors.New("Missing key.")
	}

	value, err := c.file.Section(c.Env).Key(key).Int()
	if err != nil {
		return 0, errors.New("Error parsing key.")
	}

	return value, nil
}

func (c *Config) GetStringDefault(key, defaultValue string) string {
	if !c.file.Section(c.Env).HasKey(key) {
		return defaultValue
	}

	return c.file.Section(c.Env).Key(key).String()
}

func (c *Config) GetIntDefault(key string, defaultValue int) int {
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

func (c *Config) GetRDSConfig() (*RDSConfig, error) {
	host, err := c.GetString("rds.core.host")
	if err != nil {
		return nil, err
	}

	username, err := c.GetString("rds.core.username")
	if err != nil {
		return nil, err
	}

	password, err := c.GetString("rds.core.password")
	if err != nil {
		return nil, err
	}

	return &RDSConfig{
		Host:     host,
		Port:     c.GetIntDefault("rds.core.port", 3306),
		Database: c.GetStringDefault("rds.core.database", "core"),
		Username: username,
		Password: password,
		PoolSize: c.GetIntDefault("rds.core.pool", 1),
	}, nil
}
