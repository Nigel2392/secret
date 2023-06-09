package secret

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

var KEY Key = New("THIS-IS-NOT-SAFE-FOR-PRODUCTION-" + strconv.FormatInt(time.Now().Unix(), 10))

type Key interface {
	// Get the bytes crypto interface.
	Bytes() *_byteCrypto
	// Get the html safe crypto interface.
	HTMLSafe() *_htmlSafe
	// Get a new aes key from the secret key.
	KeyFromSecret() *[32]byte
	// Generate a sha256 hash from the secret key.
	Sha256() string
	// Verify a sha256 hash with the secret key.
	VerifySha256(hash string) bool
	// Sign a string with the secret key.
	Sign(data string) string
	// Verify a string with the secret key.
	Verify(data, signature string) bool
	// Raw returns the raw secret key.
	Raw() string
}

// Secret key struct.
type secretKey struct {
	SECRET_KEY string
	secret_key *[32]byte
	bytes      *_byteCrypto
	htmlSafe   *_htmlSafe
}

// Create a new secret key.
func New(key string) Key {
	var k = &secretKey{
		SECRET_KEY: key,
	}
	k.bytes = &_byteCrypto{k: &k}
	k.htmlSafe = &_htmlSafe{k: &k}
	return k
}

// Get the bytes crypto.
func (c *secretKey) Bytes() *_byteCrypto {
	return c.bytes
}

// Get the html safe crypto.
func (c *secretKey) HTMLSafe() *_htmlSafe {
	return c.htmlSafe
}

// Create a new aes key from the secret key.
func (c *secretKey) KeyFromSecret() *[32]byte {
	if c.secret_key != nil {
		return c.secret_key
	}
	if c.SECRET_KEY == "" {
		panic("No secret key set")
	}
	var key [32]byte
	if len(c.SECRET_KEY) < 32 {
		for i := 0; i < 32; i++ {
			key[i] = c.SECRET_KEY[i%len(c.SECRET_KEY)]
		}
	} else {
		// Compact the key to 32 bytes.
		for i := 0; i < len(c.SECRET_KEY); i++ {
			key[i%32] += byte(int(c.SECRET_KEY[i]) - i)
		}
	}
	c.secret_key = &key
	return &key
}

// Public key from the secret key.
func (c *secretKey) Sha256() string {
	var hash = sha256.Sum256([]byte(c.SECRET_KEY))
	return hex.EncodeToString(hash[:])
}

// Verify a sha256 hash with the secret key.
func (c *secretKey) VerifySha256(hash string) bool {
	if c.SECRET_KEY == "" {
		panic("No secret key set")
	}
	return c.Sha256() == hash
}

// Sign a string with the secret key.
func (c *secretKey) Sign(data string) string {
	if c.SECRET_KEY == "" {
		panic("No secret key set")
	}
	return Hash([]byte(data), []byte(c.SECRET_KEY))
}

// Verify a string with the secret key.
func (c *secretKey) Verify(data, signature string) bool {
	if c.SECRET_KEY == "" {
		panic("No secret key set")
	}
	return Compare([]byte(data), []byte(c.SECRET_KEY), signature)
}

// Raw returns the raw secret key.
func (c *secretKey) Raw() string {
	return c.SECRET_KEY
}
