package json

import (
	"encoding/json"

	"github.com/mainvec/ugo/oencoding"
)

var _ oencoding.Encoding = (*JSONEncoding)(nil)

type JSONEncoding struct {
}

func init() {
	jsonEncoding := &JSONEncoding{}
	oencoding.RegisterEncoding("application/json", jsonEncoding)
	oencoding.RegisterEncoding("json", jsonEncoding)
}

func (j *JSONEncoding) Encode(o any) ([]byte, error) {
	return json.Marshal(o)
}

func (j *JSONEncoding) Decode(data []byte, o any) error {
	return json.Unmarshal(data, o)
}

func (j *JSONEncoding) MimeType() string {
	return "application/json"
}
