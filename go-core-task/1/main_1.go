package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func stringBuilder(value ...interface{}) string {
	var builder strings.Builder
	for _, v := range value {
		switch valueType := v.(type) {
		case int:
			fmt.Printf("%d is an int\n", valueType)
			builder.WriteString(strconv.Itoa(valueType))
		case string:
			fmt.Printf("%v is a string\n", valueType)
			builder.WriteString(valueType)
		case float64:
			str := strconv.FormatFloat(valueType, 'f', -1, 64)
			fmt.Printf("%s is a float64\n", str)
			builder.WriteString(str)
		case complex64:
			fmt.Printf("%v is a complex64\n", valueType)
			builder.WriteString(strconv.FormatComplex(complex128(valueType), 'f', -1, 64))
		case bool:
			fmt.Printf("%v is a bool\n", valueType)
			builder.WriteString(fmt.Sprint(valueType))
		default:
			fmt.Printf("%v has a wrong type\n", valueType)
		}
	}
	return builder.String()
}

func runeConv(str string) []rune {
	return []rune(str)
}

func hash(str []rune) [32]byte {
	center := (len(str)) / 2
	str = append(str[:center], append([]rune("go-2024"), str[center:]...)...)
	return sha256.Sum256(([]byte)((string)(str)))
}

func main() {
	var (
		numDecimal     int       = 42       // Десятичная система
		numOctal       int       = 052      // Восьмеричная система
		numHexadecimal int       = 0x2A     // Шестнадцатиричная система
		pi             float64   = 3.14     // Тип float64
		name           string    = "Golang" // Тип string
		isActive       bool      = true     // Тип bool
		complexNum     complex64 = 1 + 2i   // Тип complex64
	)
	resultString := stringBuilder(
		numDecimal,
		numHexadecimal,
		numOctal,
		pi,
		name,
		isActive,
		complexNum)
	fmt.Println(resultString)
	resultRune := runeConv(resultString)
	fmt.Println(resultRune)
	resultSha := hash(resultRune)
	fmt.Printf("%x", resultSha)
}
