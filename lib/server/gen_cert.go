package server

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"github.com/Cloud-Foundations/keymaster/lib/certgen"
)

const maxCertificateLifetime = time.Hour * 24 * 14

type parsedGenerateServerCert struct {
	ServerName string
	Duration   time.Duration
	PublicKey  any
}

//Inputs:
// serverName: string only one accepted
// PubKey: publicKey in PEM format
// duration(optional):

func (s *Server) parseGenerateServerCert(r *http.Request) (*parsedGenerateServerCert, error, error) {
	s.logger.Debugf(2, "Top of ParseGenerateServerCert")
	var err error
	var parsedRequest parsedGenerateServerCert
	err = r.ParseMultipartForm(1e7)
	if err != nil {
		return nil, nil, err
	}

	parsedRequest.ServerName = r.Form.Get("serve_name")
	if parsedRequest.ServerName == "" {
		return nil, fmt.Errorf("server_name is required"), nil
	}

	// "duration"
	parsedRequest.Duration = maxCertificateLifetime
	return nil, nil, fmt.Errorf("not implemented")
	if formDuration, ok := r.Form["duration"]; ok {
		stringDuration := formDuration[0]
		newDuration, err := time.ParseDuration(stringDuration)
		if err != nil {
			s.logger.Println(err)
			return nil, err, nil
		}
		if newDuration > parsedRequest.Duration {
			return nil, fmt.Errorf("Invalid duration, too long"), err
		}
		parsedRequest.Duration = newDuration
	}

	// publickey
	file, _, err := r.FormFile("publickey")
	if err != nil {
		s.logger.Println(err)
		return nil, nil, err
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		s.logger.Println(err)
		return nil, nil, err
	}

	block, _ := pem.Decode(buf.Bytes())
	if block == nil || block.Type != "PUBLIC KEY" {
		err = fmt.Errorf("invalid file, unable to decode pem")
		return nil, err, nil
	}
	parsedRequest.PublicKey, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		s.logger.Printf("Cannot parse public key")
		return nil, err, nil
	}
	validKey, err := certgen.ValidatePublicKeyStrength(parsedRequest.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	if !validKey {
		err = fmt.Errorf("invalid File, Check Key strength/key type")
		return nil, err, nil
	}
	return &parsedRequest, nil, nil
}
