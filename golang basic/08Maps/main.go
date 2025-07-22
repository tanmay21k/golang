package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]string)
	map1["py"] = "python"
	map1["js"] = "javascript"
	map1["java"] = "java"
	map1["ts"] = "typescript"

	fmt.Println(map1)
	fmt.Println(map1["py"])
	delete(map1, "js")
	fmt.Println(map1)
	fmt.Println(len(map1))

	map2 := make(map[int]int)
	map2[0] = 32
	map2[1] = 232
	map2[2] = 2
}
