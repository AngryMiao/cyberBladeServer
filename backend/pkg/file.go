package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// PathIsExist return this file whether exist
func PathIsExist(directory string) bool {
	_, err := os.Stat(directory)
	if err != nil {
		return false
	}
	return true
}

func ReadFileToString(path string) (string, error) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Convert []byte to string and print to screen
	text := string(content)

	return text, nil
}

func ReadAMMasterMapField(keyboard string) ([][]string, error) {
	var keyMapJson [][]string
	file, err := os.Open(fmt.Sprintf("%s.json", keyboard))
	if err != nil {
		return keyMapJson, fmt.Errorf("Error opening file: %s", err.Error())
	}
	defer file.Close()

	// 读取文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		return keyMapJson, fmt.Errorf("Error reading file: %s", err.Error())
	}

	json.Unmarshal(data, &keyMapJson)
	return keyMapJson, nil
}
