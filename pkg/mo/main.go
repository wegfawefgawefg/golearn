package main

import (
	"errors"
	"log"

	"github.com/samber/lo"
	"github.com/samber/mo"
)

func fail_if_odd(a int) mo.Result[int] {
	if a%2 != 0 {
		return mo.Err[int](errors.New("odd number"))
	}
	return mo.Ok(a)
}

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		res := fail_if_odd(i)
		if res.IsError() {
			log.Println(res.Error())
		}
	}

	// map
	numbers := []int{1, 2, 3, 4, 5}
	for _, n := range numbers {
		res := fail_if_odd(n)
		if res.IsError() {
			log.Println(res.Error())
		}
		num := res.MustGet()
		log.Println(num)
	}

	lo.Map(numbers, func(n int, index int) mo.Result[int] {
		res := fail_if_odd(n)
		if res.IsError() {
			log.Println(res.Error())
		}
		return res
	})
}
