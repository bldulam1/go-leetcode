package multiply_strings

import (
	"fmt"
	"strings"
)

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	const c0 = 48

	num1Len, num2Len := len(num1), len(num2)
	if num1Len > num2Len {
		num2 = strings.Repeat("0", num1Len-num2Len) + num2
	} else if num1Len < num2Len {
		num1 = strings.Repeat("0", num2Len-num1Len) + num1
	}

	subProducts := make([][]byte, len(num1))
	longestSubProduct := 0

	for i1, c1 := len(num1)-1, 0; i1 >= 0; i1 -= 1 {
		subProducts[c1] = make([]byte, c1)
		carryOver := byte(0)

		for i2, c2 := len(num1)-1, 0; i2 >= 0; i2 -= 1 {
			n1, n2 := num1[i1]-c0, num2[i2]-c0
			prod := n1*n2 + carryOver
			carryOver = prod / 10
			subProducts[c1] = append(subProducts[c1], prod%10)

			c2 += 1
		}

		if carryOver > 0 {
			subProducts[c1] = append(subProducts[c1], carryOver)
		}

		subProductsLen := len(subProducts[c1])
		if subProductsLen > longestSubProduct {
			longestSubProduct = subProductsLen
		}
		c1 += 1
	}

	grandTotal := make([]byte, 0)
	carryOver := 0
	for ind := 0; ind < longestSubProduct; ind += 1 {
		total := carryOver
		for j := 0; j < len(subProducts); j++ {
			if ind < len(subProducts[j]) {
				total += int(subProducts[j][ind])
			}
		}
		carryOver = total / 10
		grandTotal = append([]byte{byte(total%10) + c0}, grandTotal...)
	}

	product := string(grandTotal)

	if carryOver > 0 {
		product = fmt.Sprintf("%d%s", carryOver, product)
	}

	for product[0] == c0 {
		product = product[1:]
	}
	return product
}
