package eval

import (
	"strconv"
	"strings"
)

func Eval(query string) string {
	var result strings.Builder

	left := ""
	right := ""
	operator := rune(0)

	for i := 0; i < len(query); i++ {
		char := query[i]

		if char == ' ' {
			continue
		}

		if char == '+' || char == '-' || char == '*' || char == '/' {
			if operator == rune(0) {
				operator = rune(char)
			} else {
				num1, err := strconv.Atoi(strings.TrimSpace(left))
				if err != nil {
					return ""
				}

				num2, err := strconv.Atoi(strings.TrimSpace(right))
				if err != nil {
					return ""
				}

				var res int
				switch operator {
				case '+':
					res = num1 + num2
				case '-':
					res = num1 - num2
				case '*':
					res = num1 * num2
				case '/':
					if num2 == 0 {
						return "Division by zero error"
					}
					res = num1 / num2
				}
				left = strconv.Itoa(res)
				right = ""
				operator = rune(char)
			}
		} else {
			if operator == rune(0) {
				left += string(char)
			} else {
				right += string(char)
			}
		}
	}

	num1, err := strconv.Atoi(strings.TrimSpace(left))
	if err != nil {
		return ""
	}

	num2, err := strconv.Atoi(strings.TrimSpace(right))
	if err != nil {
		return ""
	}

	var res int
	switch operator {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		if num2 == 0 {
			return "Division by zero error"
		}
		res = num1 / num2
	default:
		return "Invalid operator"
	}

	result.WriteString(strconv.Itoa(res))

	return result.String()
}

