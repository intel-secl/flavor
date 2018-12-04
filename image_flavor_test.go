package flavor

import (
	"encoding/json"
	"log"
	"testing"
)

func TestImageFlavorCreation(t *testing.T) {
	flavor, err := GetImageFlavor("Cirros-enc", true,
		"http://10.1.68.21:20080/v1/keys/73755fda-c910-46be-821f-e8ddeab189e9/transfer",
		"261209df1789073192285e4e408addadb35068421ef4890a5d4d434")
	if err != nil {
		log.Println(err)
	}
	json, err := json.Marshal(flavor)

	log.Printf("Image Flavor:%s\n", string(json))
}
