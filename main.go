package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Square(x int) int {
	return x * x
}
func Power(x int, n int) int {
	product := 1
	for i := 0; i < n; i++ {
		product *= x
	}
	return product
}
func main() {

	file, err := os.Create("fast.go")
	if err != nil {
		fmt.Println("Unable to open file")
		os.Exit(1)
	}
	_, err = file.WriteString("package main\nimport (\n\"fmt\"\n)\nfunc main(){\n fmt.Println(\"smth\")\n}")
	if err != nil {
		fmt.Println("Unable to write data")
		os.Exit(1)
	}
	file.Close()

	/*_, errc := exec.Command("go", "build", "fast.go").Output()
	if errc != nil {
		fmt.Println(errc)
		os.Exit(1)
	}*/

	_, erre := exec.Command("go", "run", "fast.go").Output()
	if erre != nil {
		fmt.Println(erre)
		os.Exit(1)
	}

}
