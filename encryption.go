package flavor

/**
 *
 * @author purvades
 */

// Encryption contains information pertaining to the encryption policy of the image
type Encryption struct {
	EncryptionRequired   bool   `json:"encryption_required,omitempty"`
	KeyURL               string `json:"key_url,omitempty"`
	InitializationVector []byte `json:"initialization_vector,omitempty"`
	Digest               string `json:"digest,omitempty"`
}
