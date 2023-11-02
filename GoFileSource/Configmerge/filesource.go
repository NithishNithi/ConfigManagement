package Config

import (
	"io/ioutil"

	"github.com/warthog618/config/blob/decoder/ini"
)

// FileSource represents a configuration source from a file.
type FileSource struct {
	Path    string
	Encoder Encoder
	Decoder ini.Decoder
}

// ReadConfig reads and decodes the configuration data from the file using the specified Encoder.
func (fs *FileSource) ReadConfig() (map[string]interface{}, error) {
	fileContent, err := ioutil.ReadFile(fs.Path)
	if err != nil {
		return nil, err
	}

	var configData map[string]interface{}
	if err := fs.Encoder.Unmarshal(fileContent, &configData); err != nil {
		return nil, err
	}

	return configData, nil
}
