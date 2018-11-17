package flavor

import (
	"fmt"
	"testing"
)

func TestImageFlavorCreation(t *testing.T) {
	flavor, err := GetImageFlavor("Cirros-enc", true,
		"http://10.1.68.21:20080/v1/keys/73755fda-c910-46be-821f-e8ddeab189e9/transfer",
		[]byte("0123456"), "261209df1789073192285e4e408addadb35068421ef4890a5d4d434")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("Image Flavor:%s\n", flavor)
}
