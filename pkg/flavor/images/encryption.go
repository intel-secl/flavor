package images

/**
 *
 * @author purvades
 */

type Encryption struct {
	EncryptionRequired   bool   `json:"encryption_required,omitempty"`
	KeyURL               string `json:"key_url,omitempty"`
	InitializationVector []byte `json:"initialization_vector,omitempty"`
	Digest               string `json:"digest,omitempty"`
}
