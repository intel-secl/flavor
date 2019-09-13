package flavor

/**
 *
 * @author purvades
 */

// Image struct defines the metadata of the image and
// encryption details such as key URL, digest etc.
type Image struct {
	Meta               Meta        `json:"meta"`
	EncryptionRequired bool        `json:"encryption_required"`
	Encryption         *Encryption `json:"encryption,omitempty"`
	IntegrityEnforced  bool        `json:"integrity_enforced"`
	Integrity          *Integrity  `json:"integrity,omitempty"`
}
