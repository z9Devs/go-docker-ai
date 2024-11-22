package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

// print struct json format to console
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func GetFileContent(filePath string) (string, error) {
	// Leggi il contenuto del file
	fContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error while reading the file: %v\n", err)
		return "", err
	}

	// Assegna il contenuto del file a una variabile
	content := string(fContent)

	return content, nil
}
