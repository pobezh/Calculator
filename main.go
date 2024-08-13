package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// отображение для перевода римских цифр(+ число 10) в арадские
var romeToArab = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// функция для определения является ли строка набором римских цифр (от 0 до 10)
func romanString(roman string) bool {
	contRomeSymbols := "IXV" //контейнер римских цифр
	romeChar := true
	for i := 0; i < len(roman); i++ {
		if !strings.Contains(contRomeSymbols, string(roman[i])) {
			romeChar = false
			break
		}
	}
	return romeChar
}

// отображение для перевода арабских цифр(+ число 10) в римские
var arabToRome = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

// отображение римских чисел
var romanMap = []struct {
	decVal int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// функция перевода десятичного числа из арабской системы счисления в римскую
func decimalToRoman(num int) string {
	result := ""
	for _, pair := range romanMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

func main() {
	//считываем строку
	scan := bufio.NewScanner(os.Stdin)
	_ = scan.Scan()
	text := scan.Text()

	errCode := 0 //переменная для кода ошибки, 0 - нет ошибок

	contOperand := "0123456789IXV" //контейнер символов операндов
	contOperator := "+-*/"         //контейнер символов операторов
	romeOperands := false          //наличие римских переменных
	arabOperands := false          //наличие арабских переменных
	var a, b, result int           // переменные для хранения операндов и результата
	strExpression := ""            //строка выражения без пробелов
	operator := ""                 //оператор (/,*,-,+)

	//удаляем пробелы из строки
	for i := 0; i < len(text); i++ {
		if string(text[i]) != " " {
			strExpression += string(text[i])
		}
	}

	//вычисляем какой оператор
	for j := 0; j < len(strExpression); j++ {
		if !strings.Contains(contOperand, string(strExpression[j])) { //пропускаем символы операндов
			if strings.Contains(contOperator, string(strExpression[j])) { //оставляем символ оператора
				operator += string(strExpression[j])
			} else { //когда присутствует символ не операнда и не оператора
				errCode = 1
				break
			}
		}
	}

	if len(operator) == 0 { //нет оператора в исходной строке
		errCode = 1
	} else if len(operator) > 1 { //больше одного оператора в исходной строке
		if string(strExpression[0]) == "-" {
			errCode = 5 //первое число отрицательное
		} else {
			errCode = 2
		}
	} else { //один оператор в исходном выражении
		operands := strings.Split(strExpression, operator) //разделяем переменные операндов оператором
		if operands[1] == "" {                             //нет 2-го операнда
			errCode = 1
		} else { //2 существующих операнда
			if numberOne, err := strconv.Atoi(operands[0]); err == nil { //первое число арабское
				if numberTwo, err2 := strconv.Atoi(operands[1]); err2 == nil { //второе число арабское
					if (numberOne >= 1 && numberOne <= 10) && (numberTwo >= 1 && numberTwo <= 10) { //оба числа от 1 до 10
						arabOperands = true
						a = numberOne
						b = numberTwo
					} else { //числа не от 1 до 10
						errCode = 5
					}
				} else { //второе числа не арабское
					if romanString(operands[1]) { //второе число римское
						errCode = 3
					} else { //второе число не римское
						errCode = 1
					}
				}
			} else { //первое число не арабское
				if !romanString(operands[0]) { //первое число не римское
					errCode = 1
				} else { //первое число римское
					if _, err2 := strconv.Atoi(operands[1]); err2 == nil { //второе число арабское
						errCode = 3
					} else { //второче число не арабское
						if !romanString(operands[1]) { //второе число не римское
							errCode = 1
						} else { //второе число римское
							//проверяем римские числа от 1 до 10
							numberRomeOne, okRomeOne := romeToArab[operands[0]]
							numberRomeTwo, okRomeTwo := romeToArab[operands[1]]
							if okRomeOne && okRomeTwo { //оба числа от 1 до 10
								romeOperands = true
								a = numberRomeOne
								b = numberRomeTwo
							} else { //числа не от 1 до 10
								errCode = 5
							}
						}
					}
				}
			}
			//проводим вычисление операции
			switch operator {
			case "-":
				result = a - b
			case "+":
				result = a + b
			case "*":
				result = a * b
			case "/":
				result = a / b
			}
			//если даны арабские числа
			if arabOperands {
				fmt.Println(result)
			} else if romeOperands { //если даны римские числа
				if result <= 0 { //если ответ не положительное число
					errCode = 4
				} else {
					fmt.Println(decimalToRoman(result))
				}
			}
		}
	}

	switch errCode {
	case 1:
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
	case 2:
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	case 3:
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
	case 4:
		fmt.Println("Выдача паники, так как в римской системе есть только положительные целые числа.")
	case 5:
		fmt.Println("Выдача паники, так как введенные данные не подходят под условие - числа от 1 до 10 включительно")
	}
}
