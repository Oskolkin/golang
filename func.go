package main

import (
	"fmt"
)

func sumDivBy3(nums ...int) int {
	sum := 0
	for _, a := range nums {
		if a%3 == 0 {
			sum += a
		}
		if a == 0 {
			continue
		}
	}
	return sum
}

func () {
	fmt.Println(sumDivBy3(1, 0, 2, 3, 6, -3))
	fmt.Println(sumDivBy3Slice([]int{0, 1, 2, 3, 3, 12}))
	fmt.Println(sumNoZero(0, 0, 1, 2, 3, 4, 5))
	fmt.Println(minPositive(1, 0, 2, 3, 6, -3))
	fmt.Println(mustGet(map[string]string{
		"key0": "wow",
		"key1": "row",
	}, "key1"))
	fmt.Println(sizeCategory(1000))
	fmt.Println(countEvenNoZero(1, 0, 2, 122, 6, -3))
	fmt.Println(firstGreaterThan(100, 0, 2, 122, 6, -3))
	fmt.Println(op(map[string]int{
		"a": 4,
		"b": 6,
	}, "mul"))
	fmt.Println(process3(0, 0, 5, 5, 10))
}

func sumDivBy3Slice(nums []int) int {
	return sumDivBy3(nums...)
}

func sumNoZero(nums ...int) (int, error) {
	sum := 0
	if len(nums) == 0 {
		return 0, fmt.Errorf("Пустой слайс")
	}
	for _, a := range nums {
		if a == 0 {
			continue
		}
		sum += a
	}
	if sum == 0 {
		return 0, fmt.Errorf("Все числа = 0")
	}
	return sum, nil
}

func minPositive(nums ...int) (min int, err error) {
	if len(nums) == 0 {
		min = 0
		err = fmt.Errorf("Пустой слайс")
		return
	}
	for _, a := range nums {
		if a > 0 {
			min = a
			break
		}
	}
	for _, a := range nums {
		if a == 0 || a < 0 {
			continue
		}
		if a > 0 && a < min {
			min = a
		}
	}
	if min == 0 {
		min = 0
		err = fmt.Errorf("Нет положительных")
		return
	}
	return min, nil
}

func mustGet(m map[string]string, key string) (string, error) {
	a, b := m[key]
	if !b {
		return "0", fmt.Errorf("Отсутствует ключ")
	}
	return a, nil
}

func sizeCategory(n int) string {
	switch {
	case n >= 1000:
		return "big"
	case n >= 100:
		return "medium"
	case n >= 0:
		return "small"
	}
	return "negative"
}

func countEvenNoZero(nums ...int) int {
	count := 0
	for _, a := range nums {
		if a == 0 {
			continue
		}
		if a%2 == 0 {
			count++
		}
	}
	return count
}

func firstGreaterThan(limit int, nums ...int) (int, error) {
	b := 0
	for _, a := range nums {
		if a > limit {
			b = a
			break
		}
	}
	if b == 0 {
		return 0, fmt.Errorf("Ошибка")
	}
	return b, nil
}

func op(m map[string]int, operation string) (int, error) {
	a, c := m["a"]
	b, d := m["b"]
	if !c {
		return 0, fmt.Errorf("Отсутствует ключ 1")
	}
	if !d {
		return 0, fmt.Errorf("Отсутствует ключ 2")
	}
	switch {
	case operation == "sum":
		return a + b, nil
	case operation == "diff":
		return a - b, nil
	case operation == "mul":
		return a * b, nil
	}
	return 0, fmt.Errorf("Неверная операция")

}

func process3(nums ...int) (sum int, count int, err error) {
	if len(nums) == 0 {
		err = fmt.Errorf("Пустой слайс")
		return
	}
	for _, a := range nums {
		if a == 0 {
			continue
		}
		if a > 0 && a%5 == 0 {
			sum += a
			count++
		}
	}
	return sum, count, nil
}
