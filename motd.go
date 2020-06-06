package motd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"moul.io/banner"
)

func Default() string {
	out := ""

	exe, _ := os.Executable()
	args0 := os.Args[0]
	if exe != "" {
		args0 = path.Base(exe)
	}
	out += banner.Inline(args0) + "\n"

	facts := []string{}
	if num := runtime.NumCPU(); num > 0 {
		facts = append(facts, fmt.Sprintf("%d CPUs", num))
	}
	if exe != "" {
		facts = append(facts, exe)
	}
	hostname, err := os.Hostname()
	if err == nil {
		facts = append(facts, hostname)
	}
	facts = append(facts, runtime.Version())
	out += strings.Join(facts, ", ") + "\n"
	return out
}
