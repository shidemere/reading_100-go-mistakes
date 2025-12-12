package main

import (
	"io"
	"os"
	"sync"
)

type Foo struct {
	// Вот пример встраивания, нужно написать тип без имени
	// Это позволит сделать "продвижение". То есть мы можем обращаться к методам Bar как будто 
	// это методы Foo
	Bar
}

type Bar struct {
	Baz int
}

func fooBar() {
	foo := Foo{}
	foo.Baz = 42
}

// Вот тут скрыта неявная проблема, связанная со встраиванием. 
// Если оно делается для полей которые мне нужно скрыть - это все ломает, так как 
// само по себе встраивание открывает извне доступ к полю
type InMem struct {
	sync.Mutex
	m map[string]int
}

func New() *InMem {
	return &InMem{m: make(map[string]int)}
}

func (i *InMem) Get(key string) (int, bool) {
	i.Lock()
	v, contains := i.m[key]
	i.Unlock()
	return v, contains
}

type Logger struct {
	writeCloser io.WriteCloser
}

func (l Logger) Write(p []byte) (int, error) {
	return l.writeCloser.Write(p)
}

func (l Logger) Close() error {
	return l.writeCloser.Close()
}

func main() {
	l := Logger{writeCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
	// m := New();
	// m.Mutex.Unlock() -> Я не должен был получить сюда доступ, но из за продвижения получил
	// А если бы не продвигал - мог бы сам управлять доступом за счёт регистра.
}
