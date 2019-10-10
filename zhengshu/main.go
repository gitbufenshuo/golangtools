package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	parsePemFile(os.Args[1])
}
func parsePemFile(path string) {
	certPEMBlock, _ := ioutil.ReadFile(path)

	certDERBlock, _ := pem.Decode(certPEMBlock)

	x509Cert, _ := x509.ParseCertificate(certDERBlock.Bytes)

	b, _ := json.Marshal(x509Cert)
	fmt.Println(string(b))
}
