//go:build windows || darwin || linux
// +build windows darwin linux

package data

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Load loads the config files.
func Load() error {
	configPath, err := path()
	if err != nil {
		return errors.Wrap(err, "failed to get config path")
	}

	if err := os.MkdirAll(configPath, 0700); err != nil {
		return errors.Wrap(err, "failed to create config directory")
	}

	configFilePath := filepath.Join(configPath, "config.yaml")
	if _, err := os.Stat(configFilePath); err != nil {
		if os.IsNotExist(err) {
			// File does not exist
			return nil
		} else {
			return errors.Wrap(err, "failed to stat config file")
		}
	} else {
		// File exists
		file, err := os.Open(configFilePath)
		if err != nil {
			return errors.Wrap(err, "failed to open config file")
		}
		defer file.Close()
		if err := yaml.NewDecoder(file).Decode(ConfigData); err != nil {
			return errors.Wrap(err, "failed to decode config file")
		}
	}

	return nil
}

// Save saves the config files to disk.
func Save() error {
	configPath, err := path()
	if err != nil {
		return errors.Wrap(err, "failed to get config path")
	}

	if err := os.MkdirAll(configPath, 0700); err != nil {
		return errors.Wrap(err, "failed to create config directory")
	}

	configFilePath := filepath.Join(configPath, "config.yaml")
	file, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.Wrap(err, "failed to create config file")
	}
	defer file.Close()
	if err := yaml.NewEncoder(file).Encode(ConfigData); err != nil {
		return errors.Wrap(err, "failed to encode config file")
	}

	return nil
}

// path returns the path to the config files.
func path() (string, error) {
	if xdgConfigHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
		return filepath.Join(xdgConfigHome, "semaphore"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home directory")
	}

	return filepath.Join(home, ".semaphore"), nil
}
