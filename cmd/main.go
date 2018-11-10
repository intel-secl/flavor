package cmd

import (
	"lib-go-flavor/pkg/flavor"
)

func main() {
	label := "Cirros-enc"
	encryptionRequired := true
	keyURL := "http://10.1.68.21:20080/v1/keys/73755fda-c910-46be-821f-e8ddeab189e9/transfer"
	IV := []byte("0123456")
	digest := "261209df1789073192285e4e408addadb35068421ef4890a5d4d434"
	result,err := flavor.GetImageFlavor(label, encryptionRequired, keyURL, IV, digest)
	if err != nil {
		println(err)
	}
	println(result)
}