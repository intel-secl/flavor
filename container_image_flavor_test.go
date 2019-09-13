package flavor

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerImageFlavorCreation(t *testing.T) {

	//flavor with encryption enabled and integrity enforced
	flavorInput, err := GetContainerImageFlavor("10.105.168.101:5000/hello-world:encrypted", true,
		"http://10.1.68.21:20080/v1/keys/73755fda-c910-46be-821f-e8ddeab189e9/transfer", true, "https://10.105.168.141:4443")
	if err != nil {
		log.Println(err)
	}
	assert.NoError(t, err)
	flavor, err := json.Marshal(flavorInput)
	assert.NoError(t, err)
	assert.NotNil(t, flavor)

	//Flavor with only integrity enforced
	flavorInput, err = GetContainerImageFlavor("10.105.168.101:5000/hello-world:signed", false, "",
		true, "https://10.105.168.141:4443")
	if err != nil {
		log.Println(err)
	}
	flavor, err = json.Marshal(flavorInput)
	assert.NoError(t, err)
	assert.NotNil(t, flavor)

	//Flavor with only encryption
	flavorInput, err = GetContainerImageFlavor("10.105.168.101:5000/hello-world:encrypted", true,
		"http://10.1.68.21:20080/v1/keys/73755fda-c910-46be-821f-e8ddeab189e9/transfer", false, "")
	if err != nil {
		log.Println(err)
	}
	flavor, err = json.Marshal(flavorInput)
	assert.NoError(t, err)
	assert.NotNil(t, flavor)

	//Flavor with no encryption and integrity enforced
	flavorInput, err = GetContainerImageFlavor("hello-world", false, "", false, "")
	if err != nil {
		log.Println(err)
	}
	flavor, err = json.Marshal(flavorInput)
	assert.NoError(t, err)
	assert.NotNil(t, flavor)
}
