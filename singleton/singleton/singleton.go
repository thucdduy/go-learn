package singleton

import "sync"

type Singleton interface {
	AddOne() int
}

// Doi tuong duy nhat su dung xuyen suot chuong trinh
type singleton struct {
	count int
}

var (
	instance *singleton
	runonce  sync.Once
)

func GetInstance() Singleton {

	runonce.Do(func() {
		instance = &singleton{count: 100}
	})
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
