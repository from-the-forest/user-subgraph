package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running integration tests")
	path := "./integration/queries"
	err := readDir(path)
	if err != nil {
		panic(err)
	}
}

func readDir(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	names, _ := file.Readdirnames(0)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v", path, name)
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}
		fmt.Println(filePath)
		if fileInfo.IsDir() {
			readDir(filePath)
		}
	}
	return nil
}
