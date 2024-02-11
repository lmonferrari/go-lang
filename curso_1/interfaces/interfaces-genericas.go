package main

import "fmt"

func interfaceGenerica(interfaceg interface{}) {
	fmt.Println(interfaceg)
}

func main() {
	interfaceGenerica(1)
	interfaceGenerica("string")
	interfaceGenerica(1.5)
}
