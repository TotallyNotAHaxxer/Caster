package CastHunter

import "fmt"

func Handler() {
	if x := recover(); x != nil {
		fmt.Println("Panic has occured -> ", x)
	}
}
