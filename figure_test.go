package figure

import (
	"fmt"
)

func ExampleAlphabet() {
	myFigure := NewFigure("Hello world", "alphabet", true)
	fmt.Println(myFigure)
	// Output:
	// H  H     l l                     l    d
	// H  H     l l                     l    d
	// HHHH eee l l ooo   w   w ooo rrr l  ddd
	// H  H e e l l o o   w w w o o r   l d  d
	// H  H ee  l l ooo    w w  ooo r   l  ddd
}
