package main

import "fmt"

func main() {
	tr := make([]string, 0, 10)
	tr = append(tr, "1")
	tr = append(tr, "2")
	fmt.Println(tr)

	transactions := []float64{}

	fmt.Println("Введите транзации:")

	for {
		transaction, err := scanTransaction()

		if err != nil {
			break
		}

		transactions = append(transactions, transaction)
	}

	fmt.Println("Транзации: ", transactions)
	fmt.Println("Баланс: ", calculateTransactions(transactions))

}

func scanTransaction() (float64, error) {
	var transaction float64

	fmt.Print("Ввод: ")
	_, err := fmt.Scan(&transaction)

	return transaction, err
}

func calculateTransactions(transactions []float64) float64 {
	var balance float64

	for _, value := range transactions {
		balance = balance + value
	}

	return balance
}
