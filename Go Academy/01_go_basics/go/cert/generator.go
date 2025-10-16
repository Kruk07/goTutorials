package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

func main() {
	mode := flag.String("mode", "", "Mode: ca | movie | character")
	name := flag.String("name", "", "Name of movie or character")
	caPath := flag.String("ca", "", "Path to CA cert (for movie)")
	moviePath := flag.String("movie", "", "Path to movie cert (for character)")
	outPath := flag.String("out", "", "Output path for cert")
	flag.Parse()

	switch *mode {
	case "ca":
		err := generateCACert(*outPath)
		exitOnError(err)
	case "movie":
		err := generateMovieCert(*name, *caPath, *outPath)
		exitOnError(err)
	case "character":
		err := generateCharacterCert(*name, *moviePath, *outPath)
		exitOnError(err)
	default:
		fmt.Println("Invalid mode. Use -mode=ca|movie|character")
		os.Exit(1)
	}
}

func generateCACert(out string) error {
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "My CA"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	return writeCertAndKey(out, certBytes, priv)
}

func generateMovieCert(name, caPath, out string) error {
	caCert, caKey, err := loadCertAndKey(caPath)
	if err != nil {
		return err
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject:      pkix.Name{CommonName: name},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(5, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, caCert, &priv.PublicKey, caKey)
	if err != nil {
		return err
	}

	return writeCertAndKey(out, certBytes, priv)
}

func generateCharacterCert(name, moviePath, out string) error {
	movieCert, movieKey, err := loadCertAndKey(moviePath)
	if err != nil {
		return err
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject:      pkix.Name{CommonName: name},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(2, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, movieCert, &priv.PublicKey, movieKey)
	if err != nil {
		return err
	}

	return writeCertAndKey(out, certBytes, priv)
}

func writeCertAndKey(out string, cert []byte, key *rsa.PrivateKey) error {
	certOut := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert})
	keyOut := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})

	dir := filepath.Dir(out)
	os.MkdirAll(dir, 0755)

	if err := os.WriteFile(out, certOut, 0644); err != nil {
		return err
	}
	keyPath := out[:len(out)-len(filepath.Ext(out))] + ".key"
	return os.WriteFile(keyPath, keyOut, 0600)
}

func loadCertAndKey(path string) (*x509.Certificate, *rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, nil, errors.New("invalid cert")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, err
	}

	keyPath := path[:len(path)-len(filepath.Ext(path))] + ".key"
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, nil, err
	}
	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil || keyBlock.Type != "RSA PRIVATE KEY" {
		return nil, nil, errors.New("invalid key")
	}
	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
