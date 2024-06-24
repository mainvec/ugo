package json

import (
	"encoding/json"

	"github.com/workoak/wogo/oencoding"
)

var _ oencoding.Encoding = (*JSONEncoding)(nil)

type JSONEncoding struct {
}

func init() {
	oencoding.RegisterEncoding("json", &JSONEncoding{})
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
