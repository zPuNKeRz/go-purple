package main

import (
	"fmt"
)

func main() {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, error := newAccount(login, password, url)
	if error != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}

	myAccount.outputAccount()
}

func promptData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scanln(&result)

	return result
}
