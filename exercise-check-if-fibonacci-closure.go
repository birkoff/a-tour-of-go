package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	return func() int {
		first := checkiffib(n)
		for ;!first; {
			n = n+1
			first = checkiffib(n)
		}
		thenum := n
		n = n+1
		return thenum
	}
}

func checkiffib(n int) bool {
	var a,b int = 0, 1
	if n==a || n==b {
		return true
	}
	var c = a+b
	for ;c<=n; {
		if c == n {
			return true
		}
		a=b
		b=c
		c=a+b
	}
	return false
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
