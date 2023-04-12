package internal

import "os"

func ReadFileBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadFileString(path string) (string, error) {
	bytes, err := ReadFileBytes(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
