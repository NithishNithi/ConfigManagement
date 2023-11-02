package Config

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

// Encoder interface defines methods for encoding and decoding data.
type Encoder interface {
	Extension() string
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// JSONEncoder implements the Encoder interface for JSON.
type JSONEncoder struct{}

// YAMLEncoder implements the Encoder interface for YAML.
type YAMLEncoder struct{}

// Extension returns the file extension for JSON.
func (je *JSONEncoder) Extension() string {
	return ".json"
}

// Marshal encodes data to JSON format.
func (je *JSONEncoder) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal decodes JSON data into a Go struct.
func (je *JSONEncoder) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Extension returns the file extension for YAML.
func (ye *YAMLEncoder) Extension() string {
	return ".yaml"
}

// Marshal encodes data to YAML format.
func (ye *YAMLEncoder) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Unmarshal decodes YAML data into a Go struct.
func (ye *YAMLEncoder) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
