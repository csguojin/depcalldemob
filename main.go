package depcalldemob

import (
	"fmt"

	"github.com/csguojin/depcalldemod"
)

func MyFunc1() {
	fmt.Println("depcalldemoB.MyFunc1")
	depcalldemod.MyFunc1()
	depcalldemod.MyFunc2()
}
