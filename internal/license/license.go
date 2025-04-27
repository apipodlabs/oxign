package license

type EncryptionAlgorithm int
type SignAlgorithm int

const (
	Ed25519 EncryptionAlgorithm = iota
)

var EncryptionAlgorithmName = map[EncryptionAlgorithm]string{
	Ed25519: "Ed25519",
}

type PublicLicense struct {
	Key string
	Alg string
}

type PrivateLicense struct {
}
