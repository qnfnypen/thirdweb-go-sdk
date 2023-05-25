package merkle

import (
	"bytes"
	"crypto/sha256"
)

// sha256Digest is the reusable digest for defaultHashFunc.
var sha256Digest = sha256.New()

// defaultHashFunc is used when no user hash function is specified.
// It implements SHA256 hash function.
func defaultHashFunc(data []byte) ([]byte, error) {
	defer sha256Digest.Reset()
	sha256Digest.Write(data)
	return sha256Digest.Sum(nil), nil
}

// defaultHashFuncParal is used by parallel algorithms when no user hash function is specified.
// It implements SHA256 hash function.
// When implementing hash functions for paralleled algorithms, please make sure it is concurrent safe.
func defaultHashFuncParal(data []byte) ([]byte, error) {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(nil), nil
}

func concatHash(b1 []byte, b2 []byte) []byte {
	return append(b1, b2...)
}

func concatSortHash(b1 []byte, b2 []byte) []byte {
	if bytes.Compare(b1, b2) < 0 {
		return append(b1, b2...)
	}
	return append(b2, b1...)
}
