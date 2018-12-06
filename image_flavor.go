package flavor

import (
	"fmt"

	"github.com/satori/go.uuid"
)

/**
 *
 * @author purvades
 */

// ImageFlavor is a flavor for an image with the encryption requirement information
// and key details of an encrypted image. 
type ImageFlavor struct {
	Image Image `json:"flavor"`
}

// GetImageFlavor is used to create a new image flavor with the specified label, encryption policy,
// key url, and digest of the encrypted image
func GetImageFlavor(label string, encryptionRequired bool, keyURL string, digest string) (*ImageFlavor, error) {
	flavorID, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Unable to create uuid. ", err)
		return nil, nil
	}

	description := Description{
		Label:      label,
		FlavorPart: "IMAGE",
	}

	meta := Meta{
		ID:          flavorID.String(),
		Description: &description,
	}
	encryption := Encryption{
		EncryptionRequired:   encryptionRequired,
		KeyURL:               keyURL,
		Digest:               digest,
	}

	imageflavor := Image{
		Meta:       meta,
		Encryption: encryption,
	}

	flavor := ImageFlavor{
		Image: imageflavor,
	}
	return &flavor, nil
}
