package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

// contains проверяет, содержится ли символ в срезе символов
func contains(slice []rune, char rune) bool {
	for _, item := range slice {
		if item == char {
			return true
		}
	}
	return false
}

// Выполнение операции
func useOperation(Arg1 int, Arg2 int, oper rune) int {
	var result int
	switch oper {
	case '+':
		result = Arg1 + Arg2
	case '-':
		result = Arg1 - Arg2
	case '*':
		result = Arg1 * Arg2
	case '/':
		if Arg2 != 0 {
			result = Arg1 / Arg2
		}
  }
  return result
}
// Проверка на римское число
func isRoman(str string) bool {
  romanNums := map[string]int {
    "I"    : 1,
    "II"   : 2,
    "III"  : 3,
    "IV"   : 4,
    "V"    : 5,
    "VI"   : 6,
    "VII"  : 7,
    "VIII" : 8,
    "IX"   : 9,
    "X"    : 10,
  }
  _,err := romanNums[str]
  return err
}
// Конвертация римского числа в целое число
func romanToInt (str string) int { /*romanNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}*/
  romanNums := map[string]int {
    "I"    : 1,
    "II"   : 2,
    "III"  : 3,
    "IV"   : 4,
    "V"    : 5,
    "VI"   : 6,
    "VII"  : 7,
    "VIII" : 8,
    "IX"   : 9,
    "X"    : 10,
  }
  
  result,_ := romanNums[str]
  return result
  
}
// Функция для преобразования целого числа в римское число
func intToRoman(num int) string {
	vals := []struct {
		val  int
		roman string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, v := range vals {
		for num >= v.val {
			roman += v.roman
			num -= v.val
		}
	}
	return roman
}

func main() {
  symbols := []rune{'+', '-', '*', '/'}
  var operationIndex int
  var operationChar rune
  var flagFoundIndex = true
  var countOper int
  reader := bufio.NewReader(os.Stdin)
  
  
  
  for {
    fmt.Print("\nINPUT:\n")
    input, _ := reader.ReadString('\n') // Чтение строки до символа новой строки
    fmt.Printf("OUTPUT:\n")
    result := strings.ReplaceAll(input, " ", "")
    countOper = 0
  
	  for i, char := range result {
		  // Проверяем, является ли текущий символ одним из целевых символов
		  if contains(symbols, char) {
			  countOper++
			  if countOper == 1 {
				  operationIndex = i
				  operationChar = char
				  flagFoundIndex = false
			  } else {
				  panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			  }
		  }
	  }
	  // Если символы не найдены, выводим сообщение
    if flagFoundIndex {
      panic("Выдача паники, так как строка не является математической операцией.")
    }
    // Получаем аргументы для арифметической операции
    strArg1 := result[:operationIndex]
    strArg2 := result[operationIndex+1:]
    strArg2 = strings.ReplaceAll(strArg2, "\n", "")
    // Проверка на принадлежность чисел к одной системе счисления
    if isRoman(strArg1) != isRoman(strArg2) {
      panic("Выдача паники, так как используются одновременно разные системы счисления или используются значения римской системы счисления, не удовлетворяющие интервалу от I до X.")
    }
    var resultOper int
    // Есл числиа римские, то выводим результат операции в римской системе счисления
    if isRoman(strArg1) && isRoman(strArg2) {
      val1 := romanToInt(strArg1)
      val2 := romanToInt(strArg2)
      resultOper = useOperation(val1, val2, operationChar)
      if resultOper < 1 {
        panic("Выдача паники, так как в римской системе нет отрицательных чисел и числа 0.")
      }
      fmt.Printf("%s\n", intToRoman(resultOper))
      // Если числа не римские, то выводим результат в арабской системе счисления
    } else {
        val1, err := strconv.Atoi(strArg1)
        if err != nil {
          panic("Выдача паники, так как не удалось преобразовать первый операнд в число.")
        }
        if (val1 < 1) || (val1 > 10) {
        panic("Выдача паники, так как первый операнд не входит в диапозон от 1 до 10.")
        }
        val2, err := strconv.Atoi(strArg2)
        if err != nil {
          panic("Выдача паники, так как не удалось преобразовать второй операнд в число.")
        }
        if (val2 < 1) || (val2 > 10) {
          panic("Выдача паники, так как второй операнд не входит в диапозон от 1 до 10.")
        }
        resultOper = useOperation(val1, val2, operationChar)
        fmt.Printf("%d\n", resultOper)
      }
  }
}