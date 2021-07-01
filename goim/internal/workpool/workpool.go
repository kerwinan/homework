package workpool

type WorkPool struct {
	taskQueue chan func()
}

func NewWorkPool(workPoolNum, maxTaskSize int64) *WorkPool {
	p := &WorkPool{
		taskQueue: make(chan func(), maxTaskSize),
	}
	for i := int64(0); i < workPoolNum; i++ {
		go p.start()
	}
	return p
}

func (w *WorkPool) Add(task func()) {
	w.taskQueue <- task
}

func (w *WorkPool) start() {
	// 捕获 panic
	for {
		(<-w.taskQueue)()
	}
}
