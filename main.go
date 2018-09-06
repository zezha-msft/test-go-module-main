package main

import ("fmt"
	v2 "github.com/zezha-msft/test-go-module-depone/v2/awesome"
	v3 "github.com/zezha-msft/test-go-module-depone/v3/awesome"
)

func main() {
	fmt.Println("SIMPLE MAIN")
	v2.PrintMessage()
	v3.PrintMessage()
}
