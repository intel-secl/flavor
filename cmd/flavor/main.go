package main

import (
	"encoding/json"
	"intel/isecl/lib/flavor"
	"intel/isecl/lib/common/validation"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"fmt"
)

func main() {

	if len(os.Args[0:]) < 2 {
		log.Fatal("Usage :  ./flavor <methodname> <parameters>")
	}
	var methodName = os.Args[1]

	switch methodName {

	case "GetImageFlavor":
		if len(os.Args[1:]) < 5 {
			log.Fatal("Usage :  ./flavor GetImageFlavor label encryptionRequired keyURL digest")
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

		inputArr := []string{os.Args[2]}
		if validateLabelErr := validation.ValidateStrings(inputArr); validateLabelErr != nil {
			fmt.Println(err.Error())
			os.Exit(1)
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

	case "GetContainerImageFlavor":
		if len(os.Args[1:]) < 6 {
            createContainerImageFlavorUsage()
		}

		encryptionRequired, err := strconv.ParseBool(os.Args[3])
		if err != nil {
			log.Println("Invalid input : encryptionRequired must be a boolean value. ")
			createContainerImageFlavorUsage()
		}

		integrityEnforced, err := strconv.ParseBool(os.Args[5])
		if err != nil {
			log.Println("Invalid input : integrityEnforced must be a boolean value. ")
			createContainerImageFlavorUsage()
		}

		if encryptionRequired {
			_, err = url.ParseRequestURI(os.Args[4])
			if err != nil {
				log.Println("Invalid input : keyURL ")
				createContainerImageFlavorUsage()
			}
		}

		if integrityEnforced {
			_, err = url.ParseRequestURI(os.Args[6])
			if err != nil {
				log.Fatal("Invalid input : notaryURL ")
				createContainerImageFlavorUsage()
			}
		}

		log.Println("Container Image flavor creation method called")
		containerImageFlavor, err := flavor.GetContainerImageFlavor(os.Args[2], encryptionRequired, os.Args[4], integrityEnforced, os.Args[6])
		if err != nil {
			log.Println(err)
		}
		json, err := json.Marshal(containerImageFlavor)
		if err != nil {
			log.Fatal("Failed to serialize flavor:", err)
		}
		log.Println(string(json))

	default:
		log.Println("Invalid method name. \n Expected values: GetImageFlavor or GetContainerImageFlavor")
	}
}

//isValidDigest method checks if the digest value is hexadecimal and 64 characters in length
func isValidDigest(value string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{64}$")
	return r.MatchString(value)
}

func createContainerImageFlavorUsage() {
	log.Fatal("Usage :  ./flavor GetContainerImageFlavor label encryptionRequired keyURL integrityEnforced notaryURL")
}
