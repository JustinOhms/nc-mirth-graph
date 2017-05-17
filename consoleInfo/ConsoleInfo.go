package consoleInfo

import "fmt"

type ConsoleInfo struct {
	verbose bool
}

func (c *ConsoleInfo) PrintVerbose(s string) {
	if c.verbose {
		c.PrintAlways(s)
	}
}

func (c *ConsoleInfo) PrintAlways(s string) {
	fmt.Println(s)
}

func NewConsoleInfo(verbose bool) *ConsoleInfo {
	return &ConsoleInfo{verbose: verbose}
}
