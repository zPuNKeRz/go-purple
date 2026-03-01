package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (a *Account) OutputAccount() {
	color.Cyan(a.Login)
	color.Red(a.Password)
	color.Green(a.Url)
	fmt.Println(a)
}

func (a *Account) generatePassword(n int) {
	result := make([]rune, n)
	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	a.Password = string(result)
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
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Проверка пароля
	if password == "" {
		newAccount.generatePassword(12)
	}

	return newAccount, nil
}
