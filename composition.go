package main

import "fmt"

// but since you can't anonymously fulfill an interface! we also need a
// struct that can take the composable functions
/*type Funkifier[A any, B any] interface {
	Do(A) B
}*/

// type parameters have to be associated with type *names*, so 
// it doesn't seem possible to have an generic function as a parameter
// aka something like 
// f func[A, B](a A) B
// 
// to get around this problem, we have to wrap up our composable functions
// into something else
type Funky [A any, B any] struct {
	Do func(a A) B
}
//func (f funky[A, B]) Do(a A) B { return f.do(a) }

func ID [A any] (a A) A {
	return a
}

// Almost unreadable but surprisingly it works!
// composes *f* after *g*
func comp[A any, B any, C any](f Funky[B, C], g Funky[A, B]) Funky[A, C] {
	// using Funkifier instead of Funky doesn't work due to failures of inference
	return Funky[A, C]{ Do: func(a A) C { return f.Do(g.Do(a)) } }
}

func main() {
	add1 := Funky[int, int]  {Do: func(n int) int { return n+1 }}
	isEven := Funky[int, bool] {Do: func(n int) bool { return n%2 == 0 }}
	constant := Funky[bool, string] {Do: func(b bool) string { return "const" }}
	count := Funky[string, int] {Do: func(s string) int { return len(s) }}
	intID := Funky[int, int]{Do: ID[int]}

	fmt.Println(add1.Do(2))
	fmt.Println(isEven.Do(2))
	fmt.Println(comp(isEven,add1).Do(2))
	fmt.Println(comp(constant, comp(isEven,add1)).Do(2))
	fmt.Println(comp(count, comp(constant, comp(isEven,add1))).Do(2))
	fmt.Println(comp(intID, comp(count, comp(constant, comp(isEven,comp(add1, intID))))).Do(2))
}
