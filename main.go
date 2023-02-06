package main

import (
	"fmt"

	filereader "github.com/kwamekyeimonies/Golang-Microservices/file_reader"
)

func main() {
	// filereader.File_Reader_Byte()
	// filereader.Optimized_File_Reader()
	// filereader.IO_File_Reader_Test()
	payload := []byte("Heyya, values still does not match")
	filereader.Hash_And_BroadCast(filereader.NewHashReader(payload))

	solve_quadratic := filereader.Roots{A: 2.2, B: 3.3, C: 4.4}
	quadratic_results := solve_quadratic.Roots_Determinant()
	roots_status := filereader.Roots_Checker(quadratic_results)
	//cannot use quadratic results (variable of type float 64) in argument
	fmt.Println(roots_status)

}
