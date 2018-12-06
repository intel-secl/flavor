package flavor

/**
 *
 * @author purvades
 */

// Image struct defines the metadata of the image and
// encryption details such as key URL, digest etc.
type Image struct {
	Meta       Meta       `json:"meta,omitempty"`
	Encryption Encryption `json:"encryption,omitempty"`
}
