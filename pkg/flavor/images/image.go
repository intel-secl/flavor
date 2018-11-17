package images

import (
	"lib-go-flavor/pkg/flavor/flavors"
)

/**
 *
 * @author purvades
 */

type Image struct {
	Meta       *flavors.Meta `json:"meta,omitempty"`
	Encryption *Encryption   `json:"encryption,omitempty"`
}
