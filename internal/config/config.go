package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

type Host struct {
	Alias    string
	Hostname string
	User     string
	Port     string
}

func LoadHosts() ([]Host, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(home, ".ssh", "config")
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return []Host{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg, err := ssh_config.Decode(f)
	if err != nil {
		return nil, err
	}

	var hosts []Host
	seen := make(map[string]bool)

	for _, hostBlock := range cfg.Hosts {
		for _, pattern := range hostBlock.Patterns {
			alias := pattern.String()

			if strings.ContainsAny(alias, "*?") || seen[alias] {
				continue
			}

			seen[alias] = true

			h := Host{
				Alias:    alias,
				Hostname: getVal(cfg, alias, "Hostname"),
				User:     getVal(cfg, alias, "User"),
				Port:     getVal(cfg, alias, "Port"),
			}

			hosts = append(hosts, h)
		}
	}

	return hosts, nil
}

func getVal(cfg *ssh_config.Config, alias, key string) string {
	val, err := cfg.Get(alias, key)
	if err != nil {
		return ""
	}
	return val
}
