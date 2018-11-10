package flavor

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"lib-go-flavor/pkg/flavor/flavors"
	"lib-go-flavor/pkg/flavor/images"
)
/**
 *
 * @author purvades
 */

type imageFlavor struct {
	Image images.Image	`json:"flavor"`
}

func GetImageFlavor(label string, encryptionRequired bool, keyURL string, initializationVector []byte, digest string) (string,error) {
    var meta flavors.Meta
    uuid1, err := uuid.NewV4()
    if err != nil {
		fmt.Println("Unable to create uuid. ", err)
		return "",nil
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

    var flavor imageFlavor
	flavor.Image.Meta = meta
	flavor.Image.Encryption = encryption
    return serialize(flavor)
}

func serialize(flavor imageFlavor) (string,error) {
    bytes, err := json.Marshal(flavor)
    if err != nil {
        fmt.Println("Can't serislize", err)
        return "",err
    }
    return string(bytes),nil
}

func deserialize(imageFlavorJson string) (imageFlavor,error) {
    var flavor imageFlavor
    err := json.Unmarshal([]byte(imageFlavorJson), &flavor)
    if err != nil {
        fmt.Println("Can't deserislize", err)
		return flavor,err
    }
   return flavor,nil
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
