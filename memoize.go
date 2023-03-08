package main

import "fmt"

type Funky [A any, B any] struct {
	Do func(a A) B
}

func fib(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

type Cache[A comparable, B any] struct {
	cache map[A]B
}
func NewCache[A comparable, B any]() *Cache[A, B] {
	return &Cache[A, B]{ cache: make(map[A]B,0) }
}
func (c *Cache[A, B]) Memoize(f Funky[A, B]) Funky[A, B] {
	g := func(a A) B {
		if v, ok := c.cache[a]; ok {
			return v
		}
		b := f.Do(a)
		c.cache[a] = b
		return b
	}
	return Funky[A, B] { Do: g}
}

func main() {
	fibF := Funky[int, int] { Do: fib }
	fibMemo := NewCache[int,int]().Memoize(fibF)
	fmt.Println(fibMemo.Do(50))
	fmt.Println(fibMemo.Do(50))
}
