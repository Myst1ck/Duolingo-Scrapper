package shared

import "fmt"

type Loggerer interface {
	LogData(data string)
}

type Logger struct{}

var Log Logger

func (logger Logger) LogData(data string) {
	fmt.Println(data)
}
