package main

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i <= Upper; i++ { // Loop 
		if i%15 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}}
