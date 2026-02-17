package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
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

func newAccount(login, password, urlString string) (*account, error) {
	// Validate
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	return &account{
		login:    login,
		password: password,
		url:      urlString,
	}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

func main() {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, error := newAccount(login, password, url)
	if error != nil {
		fmt.Println("Неверный формат URL")
		return
	}

	myAccount.outputAccount()
}

func promptData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scan(&result)

	return result
}
