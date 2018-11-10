package images

/**
 *
 * @author purvades
 */

type Encryption struct {
	EncryptionRequired     	bool		`json:"encryption_required"`
	KeyURL          		string		`json:"key_url"`
	InitializationVector   	[]byte		`json:"initialization_vector"`
	Digest       			string		`json:"digest"`
}

