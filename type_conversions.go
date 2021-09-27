package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := "19"
	edad_str,_ := strconv.Atoi(edad)

	fmt.Println(edad_str + 1)
}
