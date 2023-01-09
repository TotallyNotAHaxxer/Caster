package CastHunter

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

var CommonKeys = []string{
	"macAddress",
	"deviceID",
	"ID",
	"SourceVersion",
	"Version",
	"DSzr",
	"Tname",
	"TV",
	"Apple",
	":",
	"$",
	"statusFlagsRpiX",
	"audioLate",
	"tsRpkUmodelXfeat",
	"AliveLowPower",
}

func DumpFile(filename string) {
	GrabSig(filename)
	var CommonFoundKeys []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 256)
	fmt.Println("\033[31m")
	for {
		_, x := reader.Read(buffer)
		if x != nil {
			break // eof or some read error
		}
		dumpkeys := hex.Dump(buffer)
		for _, k := range dumpkeys {
			switch k {
			case ' ':
				fmt.Print("\033[38;5;50m")
			case '|':
				fmt.Print("\033[38;5;56m")
			}
			for i := range CommonKeys {
				if strings.Contains(string(k), CommonKeys[i]) {
					fmt.Print("\033[31m")
					CommonFoundKeys = append(CommonFoundKeys, string(k))
				}
			}
			fmt.Printf("%c", k)
		}
	}
	cols := []string{"Common Keys located"}
	var rows [][]string
	if CommonFoundKeys != nil {
		CommonFoundKeys = ExterminateExtraVals(CommonFoundKeys)

		for i := 0; i < len(CommonFoundKeys); i++ {
			rows = append(rows, []string{CommonFoundKeys[i]})
		}
	}
	if rows != nil {
		DrawTableSepColBased(rows, cols)
	}
}
