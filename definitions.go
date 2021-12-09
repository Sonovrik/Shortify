package Short

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func SetDefaultEnvs() error {
	envs := map[string]string{
		"ROOT_PATH": "./",
	}

	for key, value := range envs {
		if err := os.Setenv(key, value); err != nil {
			return errors.New("error while set default env ROOT_PATH: " + err.Error())
		}
	}

	return nil
}

func SetRootEnvs(key string) error {
	rootPath, err := getRootPath()
	if err != nil {
		return err
	}

	if err = os.Setenv(key, rootPath); err != nil {
		return errors.New("error while set env ROOT_PATH: " + err.Error())
	}

	return nil
}

func getRootPath() (string, error) {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("error while get root folder")
	}

	rootPath := filepath.Dir(b)

	return rootPath, nil
}
