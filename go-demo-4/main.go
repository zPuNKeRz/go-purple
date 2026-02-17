package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimestamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
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

	// Проверка логин
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	// Проверка URL
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAccount := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	// Проверка пароля
	if password == "" {
		newAccount.generatePassword(12)
	}

	return newAccount, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

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

	// AccountWithTimestamp
	myAccountWithTimestamp := accountWithTimestamp{
		account: account{
			login:    login,
			password: password,
			url:      url,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	fmt.Println("WithTimestamp", myAccountWithTimestamp, myAccountWithTimestamp.createdAt.Format(time.DateOnly))
}

func promptData(prompt string) string {
	fmt.Print(prompt + " ")
	var result string
	fmt.Scanln(&result)

	return result
}
