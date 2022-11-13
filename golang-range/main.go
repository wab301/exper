package main

func length() int {
	println("length")
	return 5
}

func list() []string {
	println("list")
	return []string{"1", "2", "3"}
}

var (
	watch = make(chan int, 10)
)

func next() <-chan int {
	println("chan")
	return watch
}

type watch2 struct {
	watch chan string
}

func (w *watch2) next() <-chan string {
	println("watch2")
	return w.watch
}

func main() {
	// len
	for i := 0; i < length(); i++ {
		println(i)
	}

	// range array
	for s := range list() {
		println(s)
	}

	// range chan
	for i := 0; i < 10; i++ {
		watch <- i
	}
	for b := range next() {
		println(b)
		if b == 9 {
			break
		}
	}

	w := &watch2{
		watch: make(chan string, 10),
	}
	for i := 0; i < 10; i++ {
		w.watch <- "a"
	}

	for res := range w.next() {
		println(res)
	}
}
