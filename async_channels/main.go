// main.go
package main

import (
	"os"
)

func main() {
	logic := &fetchEtagCtx{}
	bargs := &BusinessArgs{
		IPs: []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"},
	}

	ret, err := logic.Spawn(bargs)
	if err != nil {
		panic(err)
	}

	tput := &Tput{}
	tput.Sc(os.Stdout)

	RenderEachLine(ret)
}

