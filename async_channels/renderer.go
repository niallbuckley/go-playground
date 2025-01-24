package main

import (
	"fmt"
)


func RenderEachLine(ret *BusinessResult){
	run := true
	for run {
		select {
		// ok in the case of a channel indicates if the channel is still open or not.
		case st, ok := <-ret.Status:
			if !ok {
				// chanel has closed - stop running.
				run = false
				continue
			}

			//tput.Rc(os.Stdout)
			for i, p := range st {
				//tput.ClearLine(os.Stdout)
				fmt.Printf("%d- %s: %s\n", i, p.IP, p.State)
			}

			//fmt.Printf("\n")
			//tput.Cuu(os.Stdout, 1)

		case err, ok := <-ret.Error:
			if !ok {
				ret.Error = nil
				continue
			}

			fmt.Printf("Received error: %v\n", err)
		}
	}

}
