package CastHunter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// all devices
func RokuDevStartConsoleall() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	Banner()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("| Welcome to the interactive RokuConsole for the Caster framework     |")
	fmt.Println("| this console allows you to basically make your console a remote     |")
	fmt.Println("| which also means you can do whatever you want as far as key presses |")
	fmt.Println("| below this box will be a list of commands for this console          |")
	fmt.Println("------------------------------------------------------------------------")
	cols := []string{"Key", "Description"}
	var rows [][]string
	rows = append(rows,
		[]string{"up", "goes up one"},
		[]string{"home", "goes to the home page"},
		[]string{"play", "hits the play button"},
		[]string{"down", "goes down one slide"},
		[]string{"left", "goes left one slide"},
		[]string{"right", "goes right one side"},
		[]string{"OK", "presses OK"},
		[]string{"rewind", "rewinds"},
		[]string{"fastforward", "fast forwards"},
		[]string{"options", "hits options"},
		[]string{"poweroff", "poweroff or powerdown the TV"},
		[]string{"vup", "volume up"},
		[]string{"vdown", "volume down"},
		[]string{"mute", "volume mute"},
		[]string{"views", "Show sent requests INT"},
		[]string{"exitc", "returns back to the caster console"},
	)
	DrawTableSepColBased(rows, cols)
	input := bufio.NewReader(os.Stdin)
	fmt.Print("\033[38;5;50m(\033[38;5;57mRoku@*HOSTS\033[38;5;50m)\033[38;5;163m>> ")
	success := 0
	for {
		Command, _ := input.ReadString('\n')
		Command = strings.Replace(Command, "\n", "", -1)
		if len(Command) != 0 {
			getexec := Keys[Command]

			if getexec == "" {
				if Command == "views" {
					DrawUtilsBox(fmt.Sprint(success))
				} else if Command == "exitc" {
					fmt.Println("\x1b[H\x1b[2J\x1b[3J")
					Banner()
					return
				} else {
					fmt.Println("Uhoh! the command did not exist within the map, try again? ")
				}
			} else {
				for _, h := range RokuHosts {
					go func(h string) {
						url := fmt.Sprintf(getexec, h)
						NewPostNoData(url, false)
						success++
					}(h)
				}
			}
		}
		time.Sleep(2 * time.Second)
		fmt.Print("\033[38;5;50m(\033[38;5;57mRoku@*HOSTS\033[38;5;50m)\033[38;5;163m>> ")
	}
}

func RokuDevStartConsole() {
	fmt.Println("\x1b[H\x1b[2J\x1b[3J")
	Banner()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("| Welcome to the interactive RokuConsole for the Caster framework     |")
	fmt.Println("| this console allows you to basically make your console a remote     |")
	fmt.Println("| which also means you can do whatever you want as far as key presses |")
	fmt.Println("| below this box will be a list of commands for this console          |")
	fmt.Println("------------------------------------------------------------------------")
	cols := []string{"Key", "Description"}
	var rows [][]string
	rows = append(rows,
		[]string{"up", "goes up one"},
		[]string{"home", "goes to the home page"},
		[]string{"play", "hits the play button"},
		[]string{"down", "goes down one slide"},
		[]string{"left", "goes left one slide"},
		[]string{"right", "goes right one side"},
		[]string{"OK", "presses OK"},
		[]string{"rewind", "rewinds"},
		[]string{"fastforward", "fast forwards"},
		[]string{"options", "hits options"},
		[]string{"poweroff", "poweroff or powerdown the TV"},
		[]string{"vup", "volume up"},
		[]string{"vdown", "volume down"},
		[]string{"mute", "volume mute"},
		[]string{"views", "Show sent requests INT"},
		[]string{"exitc", "returns back to the caster console"},
	)
	DrawTableSepColBased(rows, cols)
	input := bufio.NewReader(os.Stdin)
	fmt.Printf("\033[38;5;50m(\033[38;5;57mRoku@%s\033[38;5;50m)\033[38;5;163m>> ", TargetMain)
	success := 0
	for {
		Command, _ := input.ReadString('\n')
		Command = strings.Replace(Command, "\n", "", -1)
		if len(Command) != 0 {
			getexec := Keys[Command]

			if getexec == "" {
				if Command == "views" {
					DrawUtilsBox(fmt.Sprint(success))
				} else if Command == "exitc" {
					fmt.Println("\x1b[H\x1b[2J\x1b[3J")
					Banner()
					return
				} else {
					fmt.Println("Uhoh! the command did not exist within the map, try again? ")
				}
			} else {
				url := fmt.Sprintf(getexec, TargetMain)
				NewPostNoData(url, false)
				success++
			}
		}
		fmt.Printf("\033[38;5;50m(\033[38;5;57mRoku@%s\033[38;5;50m)\033[38;5;163m>> ", TargetMain)
	}
}
