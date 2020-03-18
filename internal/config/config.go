package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
	"path/filepath"
	"runtime"
)

const envFileName = ".env"

func init() {
	MustLoadEnv()
}

// NewConfig returns the settings from the environment.
func NewConfig() *Config {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

// RootDir returns root dir of project
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../..")
}

// Load .env file in root path of module and return err
func LoadEnv() error {
	envFile := filepath.Join(RootDir(), envFileName)
	return gotenv.Load(envFile)
}

// MustLoadEnv is the same as LoadEnv but panics if an error occurs
func MustLoadEnv() {
	err := LoadEnv()
	if err != nil {
		panic(err)
	}
}
