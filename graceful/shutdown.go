package graceful

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var instance *Shutdown

type Shutdown struct {
	signals    []os.Signal
	closeFuncs []func()
	once       sync.Once
}

// NewShutdown 返回默认的关闭钩子实例
func NewShutdown() *Shutdown {
	if instance == nil {
		instance = &Shutdown{
			signals:    []os.Signal{syscall.SIGINT, syscall.SIGTERM},
			closeFuncs: []func(){},
		}
	}
	return instance
}

// WithSignals 添加需要监听的信号
func (h *Shutdown) WithSignals(signals ...os.Signal) *Shutdown {
	h.signals = append(h.signals, signals...)
	return h
}

// AddCloseFuncs 添加关闭时需要执行的函数
func (h *Shutdown) AddCloseFuncs(funcs ...func()) *Shutdown {
	h.closeFuncs = append(h.closeFuncs, funcs...)
	return h
}

// Wait 等待一个信号，然后执行所有关闭函数
func (h *Shutdown) Wait() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, h.signals...)

	<-sigCh

	h.once.Do(func() {
		for _, f := range h.closeFuncs {
			f()
		}
		close(sigCh)
	})
}
