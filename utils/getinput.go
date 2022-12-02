package utils

import "os"

func GetInput(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(data)
}
