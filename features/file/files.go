package file

import (
	"fmt"
	"os"
)

func WriteToFile(fileName string, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file is closed when we're done
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}