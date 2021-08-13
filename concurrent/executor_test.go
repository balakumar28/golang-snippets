package concurrent

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type MockRunnable struct {
	run func()
}

func (m MockRunnable) Run() {
	m.run()
}

func TestFixedPoolExecutor_Init(t *testing.T) {
	executor := FixedPoolExecutor(5)
	assert.Equal(t, 5, cap(executor.(*fixedPoolExecutor).workers))
}

func TestFixedPoolExecutor_Submit(t *testing.T) {
	executed := false
	executor := FixedPoolExecutor(1)
	executor.Submit(MockRunnable{run: func() {
		executed = true
	}})
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, true, executed)
}

func TestFixedPoolExecutor_Submit_Blocking(t *testing.T) {
	executed := 0
	command := MockRunnable{run: func() {
		time.Sleep(10 * time.Millisecond)
		executed += 1
	}}
	executor := FixedPoolExecutor(1)
	executor.Submit(command)
	start := time.Now().Nanosecond()
	executor.Submit(command)
	end := time.Now().Nanosecond()

	time.Sleep(21 * time.Millisecond)
	assert.Equal(t, 2, executed)
	assert.True(t, (end - start)/int(time.Millisecond) >= 10)
}
