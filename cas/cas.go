package cas

import (
	"os"

	"github.com/kevalsabhani/go-cas-storage/config"
	"github.com/kevalsabhani/go-cas-storage/utils"
)

func StoreInCAS(content string) (string, error) {
	key := utils.GenerateContentHash(content)
	if err := os.WriteFile(config.StoragePath+key, []byte(content), 0644); err != nil {
		return "", err
	}
	return key, nil
}

func RetrieveFromCAS(key string) (string, error) {
	content, err := os.ReadFile(config.StoragePath + key)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
