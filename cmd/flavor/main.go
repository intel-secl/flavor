package main

import (
	"encoding/json"
	"intel/isecl/lib/flavor"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args[0:]) < 2 {
		log.Fatal("Usage :  ./lib-flavor <methodname> <parameters>")
	}
	var methodName = os.Args[1]

	switch methodName {

	case "GetImageFlavor":
		log.Println("Image flavor creation method called")
		if len(os.Args[1:]) < 5 {
			log.Fatal("Usage :  ./lib-flavor GetImageFlavor label encryptionRequired keyURL digest")
		}
		encryptionRequired, _ := strconv.ParseBool(os.Args[3])
		imageFlavor, err := flavor.GetImageFlavor(os.Args[2], encryptionRequired, os.Args[4], os.Args[5])
		if err != nil {
			log.Println(err)
		}
		json, err := json.Marshal(imageFlavor)
		if err != nil {
			log.Fatal("Failed to serialize flavor:", err)
		}
		log.Println(string(json))

	default:
		log.Println("Invalid method name. \nExpected values: GetImageFlavor")
	}
}
