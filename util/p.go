package util

import (
	C "coconut/constant"
	"fmt"
	"runtime"
)

func PrintVersion() {
	fmt.Println(`                                 _   
                                | |  
  ___ ___   ___ ___  _ __  _   _| |_ 
 / __/ _ \ / __/ _ \| '_ \| | | | __|
| (_| (_) | (_| (_) | | | | |_| | |_ 
 \___\___/ \___\___/|_| |_|\__,_|\__|

` + "v" + C.Version + "_" + runtime.GOARCH + "_" + runtime.Version())
}
