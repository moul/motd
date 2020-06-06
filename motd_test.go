package motd_test

import (
	"fmt"
	"strings"
	"testing"

	"moul.io/banner"
	"moul.io/motd"
)

func Example() {
	fmt.Print(motd.Default())
}

func TestDefault(t *testing.T) {
	ret := motd.Default()
	fmt.Println(ret)

	banner := banner.Inline("motd.test")
	if !strings.Contains(ret, banner) {
		t.Errorf("should contain the banner of motd.test")
	}
}
