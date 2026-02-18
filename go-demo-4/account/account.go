package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

type Account struct {
	login    string `json:"login"`
	password string `json:"password"`
	url      string `json:"url"`
}

type AccountWithTimestamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

func (a Account) OutputAccount() {
	color.Cyan(a.login)
	color.Red(a.password)
	color.Green(a.url)
}

func (a *Account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	a.password = string(result)
}

func NewAccount(login, password, urlString string) (*Account, error) {

	// Проверка логин
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	// Проверка URL
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAccount := &Account{
		login:    login,
		password: password,
		url:      urlString,
	}

	field, _ := reflect.TypeOf(newAccount).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))

	// Проверка пароля
	if password == "" {
		newAccount.generatePassword(12)
	}

	return newAccount, nil
}
