package filereader

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Optimized_File_Reader(){

	file_demo, err := os.Open("love-line")
	if err != nil{
		log.Fatal(err)
	}

	defer file_demo.Close()

	bytesReader := make([]byte, 79)


	n, err := file_demo.Read(bytesReader)
	if err == io.EOF{
		fmt.Printf("End of file reached we go %d bytes on the last read", n)
	} else {
		fmt.Printf("Unexpected error reading from the file : %v\n", err)
	}
}