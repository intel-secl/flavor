package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"lib-go-flavor/cmd/flavors"
	"lib-go-flavor/cmd/images"
)
/**
 *
 * @author purvades
 */

func GetImageFlavor(label string, encryptionRequired bool, keyURL string, initializationVector []byte, digest string) string {
    var meta flavors.Meta
    uuid1, err := uuid.NewV4()
    if err != nil {
		fmt.Println("Unable to create uuid. ", err)
	}
	meta.ID = uuidToString(uuid1)

	var description flavors.Description
    description.Label = label
    description.FlavorPart = "IMAGE"
    meta.Description = description

    var encryption images.Encryption
    encryption.EncryptionRequired = encryptionRequired
    encryption.KeyURL = keyURL
    encryption.InitializationVector = initializationVector
    encryption.Digest = digest

    var imageFlavor images.Image
	imageFlavor.Meta = meta
	imageFlavor.Encryption = encryption
    return serialize(imageFlavor)
}

func serialize(imageFlavor images.Image) string {
    bytes, err := json.Marshal(imageFlavor)
    if err != nil {
        fmt.Println("Can't serislize", imageFlavor)
    }
    return string(bytes)
}

func deserialize(imageFlavorJson string) images.Image {
    var imageFlavor images.Image
    err := json.Unmarshal([]byte(imageFlavorJson), &imageFlavor)
    if err != nil {
        fmt.Println("Can't deserislize", imageFlavorJson)
    }
   return imageFlavor
}

func uuidToString(u uuid.UUID) string {
		buf := make([]byte, 36)

		hex.Encode(buf[0:8], u[0:4])
		buf[8] = '-'
		hex.Encode(buf[9:13], u[4:6])
		buf[13] = '-'
		hex.Encode(buf[14:18], u[6:8])
		buf[18] = '-'
		hex.Encode(buf[19:23], u[8:10])
		buf[23] = '-'
		hex.Encode(buf[24:], u[10:])

		return string(buf)
}
