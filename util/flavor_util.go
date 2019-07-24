package util

//import "crypto"
import (
	"intel/isecl/lib/flavor"

	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type SignedImageFlavor struct {
	ImageFlavor flavor.Image `json:"flavor"`
	Signature   string       `json:"signature"`
}

//GetSignedFlavor is used to get SHA 384 signature for a provided input
func GetSignedFlavor(flavorString string, rsaPrivateKeyLocation string) (string, error) {
	var privateKey *rsa.PrivateKey
	var flavorInterface flavor.ImageFlavor
	if rsaPrivateKeyLocation == "" {
		log.Error("No RSA Key file path provided")
		return "", errors.New("No RSA Key file path provided")
	}

	priv, err := ioutil.ReadFile(rsaPrivateKeyLocation)
	if err != nil {
		log.Error("No RSA private key found")
		return "", err
	}

	privPem, _ := pem.Decode(priv)
	parsedKey, err := x509.ParsePKCS8PrivateKey(privPem.Bytes)
	if err != nil {
		log.Error("Cannot parse RSA private key from file")
		return "", err
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		log.Error("Unable to parse RSA private key")
		return "", err
	}
	hashEntity := sha512.New384()
	hashEntity.Write([]byte(flavorString))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA384, hashEntity.Sum(nil))
	signatureString := base64.StdEncoding.EncodeToString(signature)

	json.Unmarshal([]byte(flavorString), &flavorInterface)

	signedFlavor := &SignedImageFlavor{
		ImageFlavor: flavorInterface.Image,
		Signature:   signatureString,
	}

	signedFlavorJSON, err := json.Marshal(signedFlavor)
	if err != nil {
		return "", errors.New("Error while marshalling signed container image flavor: " + err.Error())
	}

	return string(signedFlavorJSON), nil
}
