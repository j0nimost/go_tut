package main

import (
	"fmt"
	"math"

	"github.com/j0nimost/go_tut/packages/str_util"
)

func main() {
	fmt.Println(math.Floor(3.14))
	fmt.Println(math.Ceil(3.14))
	fmt.Println(math.Sqrt(64))
	fmt.Println(str_util.Reverse("Hello World"))
}
