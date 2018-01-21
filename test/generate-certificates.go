// https://www.socketloop.com/tutorials/golang-create-x509-certificate-private-and-public-keys

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"github.com/davidwalter0/go-cfg"
)

// App application configuration struct
type App struct {
	DNSNames     []string `json:"dnsnames" doc:"comma separated list of dns names" default:"example.com"`
	Country      []string `json:"country" doc:"comma separated list of country abbrev" default:"US"`
	Organization []string `json:"org" doc:"comma separated list of organizations" default:"example.com"`
	Cert         string   `json:"cert" doc:"cert file path" default:"certs/cert.pem"`
	Key          string   `json:"key"  doc:"key file path" default:"certs/server.pem"`
}

var app *App = &App{}
var err error

func main() {
	if err = cfg.Parse(app); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if true {
		var j = []byte{}
		if j, err = json.MarshalIndent(app, "", "  "); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, string(j))
		// os.Exit(1)
	}
	// ok, lets populate the certificate with some data
	// not all fields in Certificate will be populated
	// see Certificate structure at
	// http://golang.org/pkg/crypto/x509/#Certificate
	template := &x509.Certificate{
		IsCA: true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1),
		DNSNames:              app.DNSNames,
		Subject: pkix.Name{
			Country:      app.Country,
			Organization: app.Organization,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(5, 5, 5),
		// see http://golang.org/pkg/crypto/x509/#KeyUsage
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	// generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err)
	}

	publickey := &privatekey.PublicKey

	// create a self-signed certificate. template = parent
	var parent = template
	cert, err := x509.CreateCertificate(rand.Reader, template, parent, publickey, privatekey)

	if err != nil {
		fmt.Println(err)
	}

	if false {
		// save private key
		pkey := x509.MarshalPKCS1PrivateKey(privatekey)
		ioutil.WriteFile("private.key", pkey, 0777)
		fmt.Println("private key saved to private.key")

		// save public key
		pubkey, _ := x509.MarshalPKIXPublicKey(publickey)
		ioutil.WriteFile("public.key", pubkey, 0777)
		fmt.Println("public key saved to public.key")
	}

	// save cert
	//	ioutil.WriteFile("cert.pem", cert, 0777)
	//	fmt.Println("certificate saved to cert.pem")
	certOut, err := os.Create(app.Cert)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create fail for %s for %s", app.Cert, err)
		os.Exit(1)
	}
	// pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
	certOut.Close()

	// // these are the files save with encoding/gob style
	// privkeyfile, _ := os.Create("server.key")
	// privkeyencoder := gob.NewEncoder(privkeyfile)
	// privkeyencoder.Encode(privatekey)
	// privkeyfile.Close()

	if false {
		// these are the files save with encoding/gob style
		privkeyfile, _ := os.Create("server-raw.key")
		privkeyencoder := gob.NewEncoder(privkeyfile)
		privkeyencoder.Encode(privatekey)
		privkeyfile.Close()

		pubkeyfile, _ := os.Create("public-raw.key")
		pubkeyencoder := gob.NewEncoder(pubkeyfile)
		pubkeyencoder.Encode(publickey)
		pubkeyfile.Close()
	}

	// this will create plain text PEM file.
	pemfile, err := os.Create(app.Key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create fail for %s for %s", app.Key, err)
		os.Exit(1)
	}

	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}
	pem.Encode(pemfile, pemkey)
	pemfile.Close()

}
