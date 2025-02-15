package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"os"
)

func EncryptFile(fileData []byte, pubKey *rsa.PublicKey) ([]byte, error) {
    encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, fileData, nil)
    return encryptedData, err
}

func SaveEncryptedFile(filename string, data []byte) error {
    return os.WriteFile("/path/to/storage/"+filename, data, 0644)
}

func DecryptFile(cipherData []byte, privKey *rsa.PrivateKey) ([]byte, error) {
    decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, cipherData, nil)
    return decryptedData, err
}



