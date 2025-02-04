package mathutils

import (
	"math"
	"strconv"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			sum += i
			otherDiv := n / i
			if otherDiv != 1 {
				sum += otherDiv
			}
		}
	}
	return sum == n
}

func IsArmStrong(n int) bool {
	absNum := n
	if n < 0 {
		absNum = -n
	}
	str := strconv.Itoa(absNum)
	power := len(str)
	sum := 0
	for _, ch := range str {
		digit := int(ch - '0')
		sum += int(math.Pow(float64(digit), float64(power)))
	}
	return sum == absNum
}

func DigitSum(n int) int {
	absNum := n
	if n < 0 {
		absNum = -n
	}
	sum := 0

	for absNum != 0 {
		sum += absNum % 10
		absNum /= 10
	}
	return sum
}
