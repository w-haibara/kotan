package worker

import "github.com/charmbracelet/log"

var procQueue = make(chan Proc)

type Proc struct {
	Name string
	Func func()
}

func init() {
	go func() {
		for {
			proc := <-procQueue
			log.Info("Dequeue", "name", proc.Name)
			go proc.Func()
		}
	}()
}

func Enqueue(name string, f func()) {
	enqueue(Proc{
		Name: name,
		Func: f,
	})
}

func enqueue(proc Proc) {
	log.Info("Enqueue", "name", proc.Name)
	procQueue <- proc
}
