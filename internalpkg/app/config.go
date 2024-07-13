package app

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/xaionaro-go/go2rtc/pkg/shell"
	"github.com/xaionaro-go/go2rtc/pkg/yaml"
)

func LoadConfig(v any) {
	for _, data := range configs {
		if err := yaml.Unmarshal(data, v); err != nil {
			Logger.Warn().Err(err).Send()
		}
	}
}

func PatchConfig(key string, value any, path ...string) error {
	if ConfigPath == "" {
		return errors.New("config file disabled")
	}

	// empty config is OK
	b, _ := os.ReadFile(ConfigPath)

	b, err := yaml.Patch(b, key, value, path...)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigPath, b, 0644)
}

type FlagConfig []string

func (c *FlagConfig) String() string {
	return strings.Join(*c, " ")
}

func (c *FlagConfig) Set(value string) error {
	*c = append(*c, value)
	return nil
}

var configs [][]byte

func SetConfigs(_configs [][]byte) {
	configs = _configs
}

func SetConfigPaths(confs FlagConfig) {
	var _configs [][]byte

	if confs == nil {
		confs = []string{"go2rtc.yaml"}
	}

	for _, conf := range confs {
		if len(conf) == 0 {
			continue
		}
		if conf[0] == '{' {
			// config as raw YAML or JSON
			_configs = append(_configs, []byte(conf))
		} else if data := parseConfString(conf); data != nil {
			_configs = append(_configs, data)
		} else {
			// config as file
			if ConfigPath == "" {
				ConfigPath = conf
			}

			if data, _ = os.ReadFile(conf); data == nil {
				continue
			}

			data = []byte(shell.ReplaceEnvVars(string(data)))
			_configs = append(_configs, data)
		}
	}
	SetConfigs(_configs)

	if ConfigPath != "" {
		if !filepath.IsAbs(ConfigPath) {
			if cwd, err := os.Getwd(); err == nil {
				ConfigPath = filepath.Join(cwd, ConfigPath)
			}
		}
		Info["config_path"] = ConfigPath
	}
}

func parseConfString(s string) []byte {
	i := strings.IndexByte(s, '=')
	if i < 0 {
		return nil
	}

	items := strings.Split(s[:i], ".")
	if len(items) < 2 {
		return nil
	}

	// `log.level=trace` => `{log: {level: trace}}`
	var pre string
	var suf = s[i+1:]
	for _, item := range items {
		pre += "{" + item + ": "
		suf += "}"
	}

	return []byte(pre + suf)
}
