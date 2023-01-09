package CastHunter

import (
	"fmt"
	"math/rand"
	"os"
)

func GenerateNumber() int {
	var minimum, maximum int
	minimum = 1000
	maximum = 10000
	return rand.Intn(maximum-minimum) + minimum
}

func WriteToFile(filename string, data []string) (WrittenFile string) {
	_, x := os.Open(filename)
	if os.IsExist(x) {
		fmt.Println("[-[ Error: (WARNING) -> Creating new file since -> ", filename, " Already exists")
		filename = filename + fmt.Sprint(GenerateNumber())
		fmt.Println("[-] Warn : (WARNING) -> New file is ", filename)
	}
	f, x := os.Create(filename)
	if x != nil {
		ErrorHandler[130]()
		fmt.Print(x)
		return
	}
	defer f.Close()
	for i := range data {
		_, x = f.WriteString(data[i])
		if x != nil {
			ErrorHandler[140]()
			fmt.Print(x)
			return
		}
	}
	WrittenFile = filename
	return WrittenFile
}
