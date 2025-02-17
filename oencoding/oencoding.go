package oencoding

import "github.com/mainvec/ugo/registry"

var (
	encodingsRegistry = registry.NewRegistry[Encoding]()
)

// Encoding is an interface that defines the methods that an encoding type must implement.
// No encoding type is provided by this package.
type Encoding interface {
	Encode(any) ([]byte, error)
	Decode([]byte, any) error
	MimeType() string
}

func RegisterEncoding(name string, enc Encoding) {
	encodingsRegistry.Register(name, enc)
}

func LookupEncoding(name string) (Encoding, bool) {
	return encodingsRegistry.Lookup(name)
}

func ListEncodings() []string {
	return encodingsRegistry.List()
}
