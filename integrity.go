package flavor

// Integrity contains information pertaining to the Integrity policy of the image
type Integrity struct {
	NotaryURL string `json:"notary_url,omitempty"`
}
