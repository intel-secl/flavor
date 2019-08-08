package flavor

/**
 *
 * @author arijitgh
 */

// SignedImageFlavor struct defines the image flavor and
// its corresponding signature
type SignedImageFlavor struct {
	ImageFlavor Image  `json:"flavor"`
	Signature   string `json:"signature"`
}
