package flavor

/**
 *
 * @author purvades
 */

// Image contains the Encryption policy and Metadata
type Image struct {
	Meta       Meta       `json:"meta,omitempty"`
	Encryption Encryption `json:"encryption,omitempty"`
}
