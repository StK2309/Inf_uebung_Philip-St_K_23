package rekursion

// Erwartet eine ganze Zahl n >= 0.
// Liefert die n-te Zahl der Fibonacci-Folge.
// Die Funktion soll rekursiv implementiert werden.
//
// Beispiele für die Folge:
// n = 0 → 0
// n = 1 → 1
// n = 2 → 1
// n = 3 → 2
// n = 4 → 3
// n = 5 → 5
//
// Hinweis: Die rekursive Definition lautet:
// Fibonacci(n) = Fibonacci(n-1) + Fibonacci(n-2)
// mit den Basisfällen: Fibonacci(0) = 0, Fibonacci(1) = 1
func Fibonacci2(n int) int {
	//TODO Fibonacci in recursiv
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci(n int) int {
	//TODO
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
