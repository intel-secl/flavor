package flavor

import (
	"fmt"

	"github.com/satori/go.uuid"
)

/**
 *
 * @author purvades
 */

type ImageFlavor struct {
	Image Image `json:"flavor"`
}

func GetImageFlavor(label string, encryptionRequired bool, keyURL string, initializationVector []byte,
	digest string) (*ImageFlavor, error) {
	uuid1, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Unable to create uuid. ", err)
		return nil, nil
	}

	description := Description{
		Label:      label,
		FlavorPart: "IMAGE",
	}

	meta := Meta{
		ID:          uuid1.String(),
		Description: &description,
	}
	encryption := Encryption{
		EncryptionRequired:   encryptionRequired,
		KeyURL:               keyURL,
		InitializationVector: initializationVector,
		Digest:               digest,
	}

	imageflavor := Image{
		Meta:       &meta,
		Encryption: &encryption,
	}

	flavor := ImageFlavor{
		Image: imageflavor,
	}
	return &flavor, nil
}
