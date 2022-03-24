package main

import (
	"fmt"
	// "strings"
	"os"
)

func forOne(fileData string) (content string, err error) {

	file, err := os.Open(fileData)

	if err != nil {
		return
	}
	defer file.Close()
	data := make([]byte, 1)
	for {
		n, err := file.Read(data)
		if err != nil { // если конец файла
			break
		}
		content += string(data[:n]) // записать содержимое файла в переменную
	}
	return
}

func forAssociation(fileData, twoFileData, result string) (err error) {
	var content string

	fileOne, err := os.Open(fileData)

	if err != nil {
		return
	}

	fileTwo, err := os.Open(twoFileData)

	if err != nil {
		return
	}

	resultFile, err := os.Create(result)

	if err != nil {
		return
	}

	defer fileOne.Close()
	defer fileTwo.Close()
	defer resultFile.Close()
	data := make([]byte, 1)

	for {
		n, err := fileOne.Read(data)
		if err != nil { // если конец файла
			break
		}
		content += string(data[:n])
	}

	for {
		n, err := fileTwo.Read(data)
		if err != nil { // если конец файла
			break
		}
		content += string(data[:n])
	}

	resultFile.WriteString(content) // записать содержимое 2 файла в resultFile.txt

	return
}

func main() {

	fmt.Println("26.5 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Написать программу аналог cat.")
	fmt.Println("------------")

	firstFile := os.Args

	if len(firstFile) == 0 {
		fmt.Println("Ошибка")
		return

	} else if len(firstFile) == 2 {
		result, err := forOne(firstFile[1])

		if err != nil {
			fmt.Println("Ошибка", err)
			return
		}

		fmt.Println(result)

	} else if len(firstFile) == 3 {

		result, err := forOne(firstFile[1])

		if err != nil {
			fmt.Println("Ошибка", err)
			return
		}

		resultSecond, err := forOne(firstFile[2])

		if err != nil {
			fmt.Println("Ошибка", err)
			return
		}

		fmt.Println(result, resultSecond)
		// Про 2 файла написано – «программа соединяет их и печатает содержимое обоих файлов на экран», но ниже в таблице показано, что выводит 2 файла последовательно. Сделал как в таблице.
	} else {
		err := forAssociation(firstFile[1], firstFile[2], firstFile[3])
		if err != nil {
			fmt.Println("Ошибка", err)
			return
		}
		fmt.Println("Успешное слияние файлов!")
	}
}
