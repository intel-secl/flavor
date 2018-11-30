package flavor

import (
	"fmt"

	"github.com/satori/go.uuid"
)

/**
 *
 * @author purvades
 */

// ImageFlavor is a flavor structure that contains whitelist metadata for a VM image
type ImageFlavor struct {
	Image Image `json:"flavor"`
}

// GetImageFlavor constructs a new ImageFlavor with the specified label, encryption policy, KMS url, and digest of the encrypted payload
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
