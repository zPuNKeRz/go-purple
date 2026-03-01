package file

import (
	"fmt"
	"os"
)

func ReadFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}

func WriteFile(content []byte, name string) {

	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")
}
