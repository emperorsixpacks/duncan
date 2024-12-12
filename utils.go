package duncan

import (
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
	resolveVars(&duncanConfig)
	if err != nil {
		return nil, err
	}
	return duncanConfig, nil
}

func resolveVars(config *map[string]interface{}) {
	for k, v := range *config {
		if str, ok := v.(string); ok {
			(*config)[k] = resolveEnv(str)
		}
	}
}

func resolveEnv(value string) string {
  fmt.Println("hello")
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
