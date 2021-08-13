package concurrent

type Runnable interface {
	Run()
}

type Executor interface {
	Submit(Runnable)
}

type fixedPoolExecutor struct {
	PoolSize    int
	workers     chan struct{}
}

func FixedPoolExecutor(poolSize int) Executor {
	exec := &fixedPoolExecutor{PoolSize: poolSize}
	exec.workers = make(chan struct{}, exec.PoolSize)
	return exec
}

func (exec *fixedPoolExecutor) Submit(runnable Runnable) {
	exec.workers <- struct{}{}
	go exec.start(runnable)
}

func (exec *fixedPoolExecutor) start(runnable Runnable) {
	runnable.Run()
	<- exec.workers
}
