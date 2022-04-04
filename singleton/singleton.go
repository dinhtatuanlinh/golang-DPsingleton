package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{count: 100}
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}