package account

import (
	"errors"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

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

func (a account) OutputAccount() {
	color.Cyan(a.login)
	color.Red(a.password)
	color.Green(a.url)
}

func (a *account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	a.password = string(result)
}

func NewAccount(login, password, urlString string) (*account, error) {

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
