package flavor

/**
 *
 * @author purvades
 */

// Encryption contains information pertaining to the encryption policy of the image
type Encryption struct {
	KeyURL string `json:"key_url,omitempty"`
	Digest string `json:"digest,omitempty"`
}
