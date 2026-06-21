package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {

	if cardNumber == "" {
		return nil
	}

	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func LuhnCheck(cardNumber string) bool {
	if cardNumber == "" {
		return false
	}

	sum := 0
	double := false
	digits := 0

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := cardNumber[i] - '0'
		if digit > 9 {
			return false
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += int(digit)
		double = !double
		digits++
	}

	return digits > 0 && sum%10 == 0
}

func loadBankData(path string) ([]Bank, error) {
	var banks []Bank
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("неверный формат строки: %q", line)
		}

		banks = append(banks, Bank{Name: strings.TrimSpace(parts[0]), Prefix: strings.TrimSpace(parts[1])})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Ошибка при сканировании: %q", err)
	}

	return banks, nil
}

func getUserInput() string {
	fmt.Print("Введите номер карты (или Enter для выхода):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')        // читать до нажатия Enter
	input = strings.TrimSpace(input)           // убрать \n и пробелы по краям
	input = strings.ReplaceAll(input, " ", "") // убрать все пробелы внутри
	input = strings.ReplaceAll(input, "-", "") // убрать все дефисы
	return input
}

func main() {
	banks, err := loadBankData("banks.txt")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Загружено банков: ", len(banks))
	}
	for {
		cardNumber := getUserInput()
		if cardNumber == "" {
			fmt.Println("Спвсибо, до скорых встречь")
			break
		}
		isValid := LuhnCheck(cardNumber)

		fmt.Println("Валиден по Луне: ", isValid)

		bank := DetectBank(cardNumber, banks)
		if bank != nil {
			fmt.Println("Банк: ", bank.Name)
		} else {
			fmt.Println("Банк: не определен")
		}
	}
}
