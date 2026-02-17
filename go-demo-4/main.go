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

func (a account) outputAccount() {
	fmt.Println(a.login, a.password, a.url)
}

func (a *account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	a.password = string(result)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

func main() {
	login := promtData("Введите логин:")
	//password := promtData("Введите пароль:")
	url := promtData("Введите URL:")

	account1 := account{
		login: login,
		url:   url,
	}

	account1.generatePassword(10)
	account1.outputAccount()
}

func promtData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scan(&result)

	return result
}
