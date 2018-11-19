
package main

import (
	"fmt"
	"github.com/wgroeneveld/gogb/proc"
)


func main() {
	cpu := proc.Z80 { Clock: 1 }
	fmt.Println(fmt.Sprintf("%d", cpu.Clock))

	fmt.Println("sup!")
}
