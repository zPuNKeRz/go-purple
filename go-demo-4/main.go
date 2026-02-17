package main

import (
	"fmt"
	"math/rand"
)

type account struct {
	login    string
	password string
	url      string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

func main() {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите URL:")

	account1 := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputAccount(&account1)
}

func promtData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scan(&result)

	return result
}

func outputAccount(account *account) {
	fmt.Println(account.login, account.password, account.url)
}

func generatePassword(n int) string {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(result)
}
