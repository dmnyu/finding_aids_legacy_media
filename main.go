package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	er := []byte("_ER_")
	files, _ := os.ReadDir("eads")
	for _, file := range files {
		fileBytes, err := os.ReadFile(filepath.Join("eads", file.Name()))
		if err != nil {
			fmt.Println("ERROR:", err.Error())
		}
		if bytes.Contains(fileBytes, er) {
			oldPath := filepath.Join("eads", file.Name())
			newpath := filepath.Join("ers", file.Name())
			os.Rename(oldPath, newpath)
		}

	}

}
