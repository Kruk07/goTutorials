package handlers

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"example.com/go_basics/go/api"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetCertificates(c echo.Context) error {
	var certs []api.Certificate
	fmt.Println("GetCertificates")
	// CA cert
	caPath := "certs/ca.pem"
	if cert, err := loadCertificate(caPath, "CA", "CA Authority"); err == nil {
		certs = append(certs, cert)
	} else {
		fmt.Println("CA cert error:", err)
	}

	// Movie certs
	movieDir := "certs/movies"
	if err := loadCertsFromDir(movieDir, "Movie", "CA Authority", &certs); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	} else {
		fmt.Println("CA movies cert error:", err)
	}

	// Character certs
	charDir := "certs/characters"
	if err := loadCertsFromDir(charDir, "Character", "Movie Authority", &certs); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	} else {
		fmt.Println("CA characters cert error:", err)
	}

	return c.JSON(http.StatusOK, certs)
}

func loadCertsFromDir(dir, certType, issuer string, out *[]api.Certificate) error {
	fmt.Println("Scanning dir:", dir)
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".pem" {
			return nil
		}
		cert, err := loadCertificate(path, certType, issuer)
		if err == nil {
			*out = append(*out, cert)
		}
		return nil
	})
}

func loadCertificate(path, certType, issuer string) (api.Certificate, error) {
	fmt.Println("Parsing cert:", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return api.Certificate{}, fmt.Errorf("read error: %w", err)
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return api.Certificate{}, fmt.Errorf("PEM decode failed for %s", path)
	}
	parsed, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return api.Certificate{}, fmt.Errorf("parse error: %w", err)
	}
	return api.Certificate{
		Id:       parsed.SerialNumber.String(),
		Type:     api.CertificateType(certType),
		IssuedTo: parsed.Subject.CommonName,
		IssuedBy: issuer,
		IssuedAt: parsed.NotBefore,
	}, nil
}
