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
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func IsArmstrong(n int) bool {
	absNum := int(math.Abs(float64(n)))
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
	n = int(math.Abs(float64(n))) 
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
