package filereader

import (
	"log"
	"os"
)

func File_Reader_Byte() {

	file_demo, err := os.Open("love-line")
	if err != nil {
		log.Fatal("Error opening file", err)
	}

	defer file_demo.Close()

	//Store words in byte
	bytesReader := make([]byte, 79)

	//Read data
	n, err := file_demo.Read(bytesReader)
	if err != nil {
		log.Fatal("Error occured whiles reading from file", err)
	}

	log.Printf("Read \"%s\"  into bytesRead (%d bytes)", string(bytesReader), n)
}
