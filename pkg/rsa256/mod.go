package rsa256

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

const privateKey = "private_key.pem"
const publicKey = "public_key.pem"
const basePath = "./private/pem"

const defaultPrivateKeyPath = basePath + "/" + privateKey
const defaultPublicKeyPath = basePath + "/" + publicKey

const defaultValidKey = "123abc"

func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		return nil, err
	}

	return io.ReadAll(file)
}

func getPrivateKey() ([]byte, error) {
	return readFile(defaultPrivateKeyPath)
}

func GetPublicKey() ([]byte, error) {
	return readFile(defaultPublicKeyPath)
}

func CreateRsaPem() error {
	publicKeyPath := defaultPublicKeyPath
	privateKeyPath := defaultPrivateKeyPath

	encode, err := Encrypt(defaultValidKey)
	decode, err := Decrypt(encode)
	if err == nil && decode == defaultValidKey {
		fmt.Println("RSA private_key.pem and private_key.pem already exist")
		return nil
	}

	err = os.MkdirAll(basePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Create folder fail:%v", err)
	}

	// 1. Generate a new 2048-bit RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("Error generating RSA key pair: %v", err)
	}

	// 2. Convert the private key to PKCS1 ASN.1 DER format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	// 3. Create a PEM block with the private key bytes and type
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	// 4. Save the private key to a file
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return fmt.Errorf("Error creating private_key.pem: %v", err)
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		return fmt.Errorf("Error saving private key: %v", err)
	}

	// 5. Extract the public key from the key pair
	publicKey := privateKey.Public().(*rsa.PublicKey)

	// 6. Convert the public key to PKIX ASN.1 DER format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("Error encoding public key: %v", err)
	}

	// 7. Create a PEM block with the public key bytes and type
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// 8. Save the public key to a file
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return fmt.Errorf("Error creating public_key.pem: %v", err)
	}
	defer publicKeyFile.Close()

	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		return fmt.Errorf("Error saving public key: %v", err)
	}

	fmt.Println("RSA key pair generated and saved to private_key.pem and public_key.pem")
	return nil
}

func Decrypt(data string) (string, error) {
	privateKeyPEM, err := getPrivateKey()
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Failed to load Private Key")
		return "", fmt.Errorf("Failed to load Private Key")
	}

	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error ParsePKCS1PrivateKey: ", err)
		return "", fmt.Errorf("Error ParsePKCS1PrivateKey: %v", err)
	}

	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("Failed to decode base64 data: %v", err)
	}

	decryptedBytes, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPrivateKey,
		decodedData,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("Error DecryptOAEP: %v", err)
	}

	return string(decryptedBytes), nil
}

func Encrypt(data string) (string, error) {
	publicKeyPEM, err := GetPublicKey()
	if err != nil {
		return "", fmt.Errorf("Error reading file: %v", err)
	}

	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return "", fmt.Errorf("Failed to load Public Key")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to ParsePKIXPublicKey: %v", err)
	}
	rsaPublicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("Failed to parse Public Key")
	}

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		rsaPublicKey,
		[]byte(data),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("Failed to encrypt data: %v", err)
	}

	encryptedData := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encryptedData, nil
}
