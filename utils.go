package duncan

import (
	"errors"
	"fmt"
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
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ymltoMap(file []byte) (map[string]interface{}, error) {
	var duncanConfig map[string]interface{}
	err := yaml.Unmarshal(file, &duncanConfig)
	resolveConfig(&duncanConfig)
	if err != nil {
		return nil, err
	}
	fmt.Println(duncanConfig)
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
	redis_config := config["redis"].(map[string]interface{})
	database := config["database"].(map[string]interface{})
	config["redis"] = resolveConfigVars(redis_config)
	config["database"] = resolveDatabseConnnection(database)
	return config
}

func resolveDatabseConnnection(config map[string]interface{}) map[string]interface{} {
	master := config["master"].(map[string]interface{})
	// slave := config["master"].(map[string]interface{})
	config["master"] = resolveConfigVars(master)
	// config["slave"] = resolveConfigVars(slave)
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
func resolveConfigVars(config map[string]interface{}) map[string]interface{} {
	for k, v := range config {
		if str, ok := v.(string); ok {
			config[k] = resolveEnv(str)
		}
	}
	return config
}
