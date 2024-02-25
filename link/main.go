package main

import (
	"fmt"
	link "main/Link"
	"os"
)

func main() {
	var fileName string
	fmt.Printf("write a file name current available are [ex1.html, ex2.html, ex3.html,ex4.html]:")
	fmt.Scanln(&fileName)
	file, err := openFile("./" + fileName)

	if err != nil {
		return
	}
	link.TraverseHtml(file)
	defer file.Close()

}

func openFile(path string) (*os.File, error) {
	filePath := path
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	return file, nil
}
