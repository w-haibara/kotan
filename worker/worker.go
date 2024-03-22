package worker

var queue = make(chan func())

func init() {
	go func() {
		for {
			f := <-queue
			go f()
		}
	}()
}

func Enqueue(f func()) {
	queue <- f
}
