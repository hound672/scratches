package main

type Worker interface {
	Do() 
}

type worker struct {
}

func (w *worker) Do()  {
}

func caller(w Worker) {
	w.Do()
}

func main() {
	w := &worker{}
	caller(w)
}
