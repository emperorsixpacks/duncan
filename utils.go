package duncan

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func validPath(configPath string) error {
	_, err := os.Stat(configPath)
	if !os.IsNotExist(err) {
		return err
	}
	return nil
}
func loadConfig(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	config, err := ymltoMap(file)
	newConfig, err := yaml.Marshal(config)
	if err != nil {
		return nil, err
	}
	return newConfig, nil
}

func ymltoMap(file []byte) (map[string]interface{}, error) {
	var duncanConfig map[string]interface{}
	err := yaml.Unmarshal(file, &duncanConfig)
	err = resolveConfig(&duncanConfig)
	if err != nil {
		return nil, err
	}
	return duncanConfig, nil
}

func resolveConfig(config *map[string]interface{}) error {
	app_config, ok := (*config)["app"].(map[string]interface{})
	connection_config, ok := (*config)["connections"].(map[string]interface{})
	if !ok {
		return errors.New("Invalid Config")
	}
	(*config)["app"] = resolveConfigVars(app_config)
	(*config)["connections"] = resolveConnectionConfig(connection_config)

	return nil
}

func resolveConnectionConfig(config map[string]interface{}) map[string]interface{} {
	for k, v := range config {
		if innerKey, ok := v.(map[string]interface{}); ok {
			config[k] = resolveConfigVars(innerKey)
		}
	}
	return config
}

func resolveConfigVars(config map[string]interface{}) map[string]interface{} {
	for k, v := range config {
		if str, ok := v.(string); ok {
			config[k] = resolveEnv(str)
		}
	}
	return config
}

func resolveEnv(value string) string {
	return os.Getenv(resolvePlaceHolder(value))
}

func resolvePlaceHolder(value string) string {
	if strings.Contains(value, "${") {
		last_index := len(value) - 1
		first_index := 2
		env_value := value[first_index:last_index]
		return env_value
	}
	return value
}
