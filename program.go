package main

func calc(n int) int {
	return n * n
}

func main() {
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)

	in1 <- 10
	in2 <- 23

	merge2Channels(calc, in1, in2, out, 5)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	go func() {
		ww1 := make(chan chan int, n)
		ww2 := make(chan chan int, n)

		go func() {
			for i := 0; i < n; i++ {
				f1 := <-<-ww1
				f2 := <-<-ww2
				out <- f1 + f2
			}
		}()

		input := func(in <-chan int, ww chan chan int) {
			for i := 0; i < n; i++ {
				w := make(chan int)
				ww <- w
				x := <-in
				go func(w chan int, x int) {
					w <- fn(x)
				}(w, x)
			}
		}

		go input(in1, ww1)
		go input(in2, ww2)
	}()

}
