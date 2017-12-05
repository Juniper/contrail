package utils

import "encoding/pem"

//EncodePEM encodes bytes for PEM format.
func EncodePEM(t string, p []byte) string {
	return string(pem.EncodeToMemory(&pem.Block{Type: t, Bytes: p}))
}

//DecodePEM decoes PEM encoded string to bytes.
func DecodePEM(p string) []byte {
	decoded, _ := pem.Decode([]byte(p))
	return decoded.Bytes
}
