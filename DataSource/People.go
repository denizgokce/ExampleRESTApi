package DataSource

import "sync"
import model "../Models"

type singleton struct {
	People []model.Person
}


var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
