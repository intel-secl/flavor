package main

import (
	"encoding/json"
	"intel/isecl/lib/flavor"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

func main() {

	if len(os.Args[0:]) < 2 {
		log.Fatal("Usage :  ./lib-flavor <methodname> <parameters>")
	}
	var methodName = os.Args[1]

	switch methodName {

	case "GetImageFlavor":
		if len(os.Args[1:]) < 5 {
			log.Fatal("Usage :  ./lib-flavor GetImageFlavor label encryptionRequired keyURL digest")
		}

		encryptionRequired, err := strconv.ParseBool(os.Args[3])
		if err != nil {
			log.Fatal("Invalid input : encryptionRequired must be a boolean value. ")
		}
		_, err = url.ParseRequestURI(os.Args[4])
		if err != nil {
			log.Fatal("Invalid input : keyURL ")
		}

		if !isValidDigest(os.Args[5]) {
			log.Fatal("Invalid input : digest must be a hexadecimal value and 64 characters in length.")
		}

		log.Println("Image flavor creation method called")
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

//isValidDigest method checks if the digest value is hexadecimal and 64 characters in length
func isValidDigest(value string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{64}$")
	return r.MatchString(value)
}
