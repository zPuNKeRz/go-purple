package main

import (
	"fmt"
	"go-demo-4/account"
	"go-demo-4/file"
)

func main() {
	file.WriteFile("Привет!!! Я файл", "file.txt")
	file.ReadFile("file.txt")

	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, error := account.NewAccount(login, password, url)
	if error != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	myAccount.OutputAccount()

}

func promptData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scanln(&result)

	return result
}
