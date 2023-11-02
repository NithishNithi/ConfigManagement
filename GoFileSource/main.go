package main

import (
	Config "GoFilesource/Configmerge"
	"fmt"
)

func main() {
	// Create instances of JSONEncoder and YAMLEncoder for handling JSON and YAML formats.
	jsonEncoder := &Config.JSONEncoder{}
	yamlEncoder := &Config.YAMLEncoder{}

	// iniDecoder := ini.NewDecoder()

	// Create FileSource instances for JSON and YAML files with respective paths and encoders.
	jsonSource := &Config.FileSource{Path: "./files/config.json", Encoder: jsonEncoder}
	yamlSource := &Config.FileSource{Path: "./files/config.yaml", Encoder: yamlEncoder}
	// iniSource := &Config.FileSource{Path: "./files/config.ini", Decoder: iniDecoder}

	// Create a CombinedFileSource with both FileSource instances.
	combinedSource := &Config.CombinedFileSource{Sources: []*Config.FileSource{jsonSource, yamlSource}}

	// Read and merge configurations from the combined source.
	config, err := combinedSource.ReadConfig()
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		return
	}

	// Print the merged configuration.
	fmt.Println("Merged Configuration:")
	for key, value := range config {
		fmt.Printf("%s: %v\n", key, value)
	}
}
