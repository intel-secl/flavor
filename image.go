package flavor

/**
 *
 * @author purvades
 */

type Image struct {
	Meta       *Meta       `json:"meta,omitempty"`
	Encryption *Encryption `json:"encryption,omitempty"`
}
