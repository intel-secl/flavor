package main

import (
	"encoding/json"
	"fmt"
	"intel/isecl/lib/common/validation"
	"intel/isecl/lib/flavor"
	"net/url"
	"os"
	"strconv"
)

func main() {

	if len(os.Args[0:]) < 2 {
		fmt.Printf("Usage : %s <methodname> <parameters>\n", os.Args[0])
		os.Exit(1)
	}
	var methodName = os.Args[1]

	switch methodName {

	case "GetImageFlavor":
		if len(os.Args[1:]) < 5 {
			fmt.Printf("Usage : %s GetImageFlavor label encryptionRequired keyURL digest\n", os.Args[0])
			os.Exit(1)
		}

		uriValue, _ := url.Parse(os.Args[4])
		protocol := make(map[string]byte)
		protocol["https"] = 0
		if validateURLErr := validation.ValidateURL(os.Args[4], protocol, uriValue.RequestURI()); validateURLErr != nil {
			fmt.Printf("Invalid key URL format: %s\n", validateURLErr.Error())
			os.Exit(1)
		}

		if validateDigestErr := validation.ValidateBase64String(os.Args[5]); validateDigestErr != nil {
			fmt.Printf("Invalid digest: %s\n", validateDigestErr.Error())
			os.Exit(1)
		}

		inputArr := []string{os.Args[2]}
		if validateLabelErr := validation.ValidateStrings(inputArr); validateLabelErr != nil {
			fmt.Printf("Invalid label: %s\n", validateLabelErr.Error())
			os.Exit(1)
		}

		encryptionRequired, err := strconv.ParseBool(os.Args[3])
		if err != nil {
			fmt.Println("Invalid input : encryptionRequired must be a boolean value. ")
			os.Exit(1)
		}

		fmt.Println("Creating image flavor...")
		imageFlavor, err := flavor.GetImageFlavor(os.Args[2], encryptionRequired, os.Args[4], os.Args[5])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		json, err := json.Marshal(imageFlavor)
		if err != nil {
			fmt.Printf("Failed to serialize flavor: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(json))

	case "GetContainerImageFlavor":
		if len(os.Args[1:]) < 6 {
			createContainerImageFlavorUsage()
		}

		encryptionRequired, err := strconv.ParseBool(os.Args[3])
		if err != nil {
			fmt.Println("Invalid input : encryptionRequired must be a boolean value.")
			createContainerImageFlavorUsage()
		}

		integrityEnforced, err := strconv.ParseBool(os.Args[5])
		if err != nil {
			fmt.Println("Invalid input : integrityEnforced must be a boolean value.")
			createContainerImageFlavorUsage()
		}

		if encryptionRequired {
			_, err = url.ParseRequestURI(os.Args[4])
			if err != nil {
				fmt.Println("Invalid input : keyURL")
				createContainerImageFlavorUsage()
			}
		}

		if integrityEnforced {
			_, err = url.ParseRequestURI(os.Args[6])
			if err != nil {
				fmt.Println("Invalid input : notaryURL")
				createContainerImageFlavorUsage()
			}
		}

		fmt.Println("Creating container image flavor...")
		containerImageFlavor, err := flavor.GetContainerImageFlavor(os.Args[2], encryptionRequired, os.Args[4], integrityEnforced, os.Args[6])
		if err != nil {
			fmt.Println(err)
		}
		json, err := json.Marshal(containerImageFlavor)
		if err != nil {
			fmt.Printf("Failed to serialize flavor: %s", err.Error())
		}
		fmt.Println(string(json))

	default:
		fmt.Println("Invalid method name. \n Expected values: GetImageFlavor or GetContainerImageFlavor")
	}
}

func createContainerImageFlavorUsage() {
	fmt.Println("Usage :  ./flavor GetContainerImageFlavor label encryptionRequired keyURL integrityEnforced notaryURL")
	os.Exit(1)
}
