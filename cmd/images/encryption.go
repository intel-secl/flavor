package images

/**
 *
 * @author purvades
 */

type Encryption struct {
	EncryptionRequired     	bool
	KeyURL          		string
	InitializationVector   	[]byte
	Digest       			string
}

