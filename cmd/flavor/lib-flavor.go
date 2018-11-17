package main

import (
	"lib-go-flavor/pkg/flavor"
	"log"
	"os"
	"strconv"
)

func main() {

	var methodName string = os.Args[1]

	switch methodName {

	case "GetImageFlavor":
		log.Printf("Image flavor creation method called")
		if len(os.Args[1:]) < 5 {
			log.Fatal("Usage :  ./lib-flavor GetImageFlavor label encryptionRequired keyURL initializationVector digest")
		}
		encryptionRequired, _ := strconv.ParseBool(os.Args[3])
		imageFlavor, err := flavor.GetImageFlavor(os.Args[2], encryptionRequired, os.Args[4], []byte(os.Args[5]), os.Args[6])
		if err != nil {
			log.Printf(err.Error())
		}
		log.Printf(imageFlavor)

	default:
		log.Printf("Invalid method name. \nExpected values: GetImageFlavor")
	}
}
