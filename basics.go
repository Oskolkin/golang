package main

import (
	"fmt"
)

func a(n int) int {
	return n * n
}

func b(a, b int) int {
	if a > b {
		return a
	}
	if a < b {
		return b
	}
	return max(a, b)

}

func c(n int) (r int) {
	r = n * 2
	return
}

func s(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("error")
	}
	return a / b, nil
}

func e(age int) (ok bool, err error) {
	if age < 0 {
		ok = false
		err = fmt.Errorf("error")
		return
	}
	if age >= 18 {
		ok = true
		return
	}
	ok = false
	return
}

func k(nums ...int) int {
	result := 0
	for _, v := range nums {
		if v < 0 {
			continue
		}
		result += v
	}
	return result
}

func v(nums []int) int {
	return k(nums...)
}

func n(m map[string]string) (string, error) {
	a, s := m["name"]
	if !s {
		return "", fmt.Errorf("значение не найдено")
	}
	return a, nil

}

func m(y int) string {
	switch {
	case y >= 90:
		return "A"
	case y >= 80:
		return "B"
	case y >= 70:
		return "C"
	}
	return "F"
}

func p(nums ...int) (sum int, err error) {
	if len(nums) == 0 {
		sum = 0
		err = fmt.Errorf("Ошибка")
		return
	}
	for _, l := range nums {
		if l == 0 {
			continue
		}
		if l%2 != 0 {
			sum += l
			return
		}
	}
	return
}

// 1
func w(nums ...int) (sum int) {
	for _, n := range nums {
		if n == 0 {
			continue
		}
		if n%2 != 0 {
			sum += n

		}
	}
	return
}

// 2
func q(nums []int) int {
	return sumOdd(nums...)
}

// 3
func avg(nums ...int) (float64, error) {
	sum := 0
	if len(nums) == 0 {
		return 0, fmt.Errorf("Error")
	}
	for _, a := range nums {
		sum += a
	}
	return float64(sum) / float64(len(nums)), nil
}

// 4
func mm(nums ...int) (min, max int, err error) {
	if len(nums) == 0 {
		return 0, 0, fmt.Errorf("Error")
	}
	min = nums[0]
	max = nums[0]
	for _, a := range nums[1:] {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return min, max, nil
}

// 5
func yk(m map[string]int) (int, error) {
	a, b := m["age"]
	if !b {
		return 0, fmt.Errorf("Error")
	}
	return a, nil
}

// 6
func ik(t int) string {
	switch {
	case t >= 30:
		return "hot"
	case t >= 15:
		return "warm"
	case t >= 0:
		return "cold"
	}
	return "freezing"
}

// 7
func cp(nums ...int) int {
	count := 0
	for _, a := range nums {
		if a == 0 {
			continue
		}
		if a > 0 {
			count++

		}
	}
	return count
}

// 8
func fE(nums ...int) (int, error) {
	b := 0
	for _, a := range nums {
		if a == 0 {
			continue
		}
		if a%2 == 0 {
			b = a
			return b, nil
		}
	}
	return b, fmt.Errorf("error")
}

// 9
func numberType(n int) string {
	switch {
	case n == 0:
		return "zero"
	case n > 0 && n%2 == 0:
		return "positive even"
	case n > 0 && n%2 != 0:
		return "positive odd"
	}
	return "negative"
}

// 10
func process2(nums ...int) (sum int, count int, err error) {

	if len(nums) == 0 {
		err = fmt.Errorf("error")
	}
	for _, a := range nums {
		if a == 0 {
			continue
		}
		if a > 0 && a%2 == 0 {
			sum += a
			count++
		}
	}
	return sum, count, err
}
