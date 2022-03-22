package main

import (
	"fmt"
	// "strings"
	"os"
)

func forOne(fileData string) (content string) {

	file, err := os.Open(fileData)

	if err != nil {
		fmt.Println("Файл отсутствует", err)
	}
	defer file.Close()
	data := make([]byte, 1)
	for {
		n, err := file.Read(data)
		if err != nil { // если конец файла
			break // выходим из цикла
		}
		content += string(data[:n]) // записать содержимое файла в переменную
	}
	return
}

func forAssociation(fileData, twoFileData, result string) (content string) {

	fileOne, err := os.Open(fileData)

	if err != nil {
		fmt.Println("Файл отсутствует", err)
	}

	fileTwo, err := os.Open(twoFileData)

	if err != nil {
		fmt.Println("Файл отсутствует", err)
	}

	resultFile, err := os.Create(result)

	if err != nil {
		fmt.Println("Файл отсутствует", err)
	}

	defer fileOne.Close()
	defer fileTwo.Close()
	defer resultFile.Close()
	data := make([]byte, 1)

	for {
		n, err := fileOne.Read(data)
		if err != nil { // если конец файла
			break // выходим из цикла
		}
		content += string(data[:n])
	}

	for {
		n, err := fileTwo.Read(data)
		if err != nil { // если конец файла
			break // выходим из цикла
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
		fmt.Println(forOne(firstFile[1]))
	} else if len(firstFile) == 3 {
		fmt.Println(forOne(firstFile[1]))
		fmt.Println(forOne(firstFile[2]))
		// Про 2 файла написано – «программа соединяет их и печатает содержимое обоих файлов на экран», но ниже в таблице показано, что выводит 2 файла последовательно. Сделал как в таблице.
	} else {
		forAssociation(firstFile[1], firstFile[2], firstFile[3])
	}
}
