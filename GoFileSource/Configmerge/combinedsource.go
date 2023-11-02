package Config

// CombinedFileSource represents a combination of multiple configuration sources.
type CombinedFileSource struct {
	Sources []*FileSource
}

// ReadConfig reads and merges configurations from multiple sources.
func (cfs *CombinedFileSource) ReadConfig() (map[string]interface{}, error) {
	configs := make([]map[string]interface{}, len(cfs.Sources))
	for i, source := range cfs.Sources {
		config, err := source.ReadConfig()
		if err != nil {
			return nil, err
		}
		configs[i] = config
	}
	return mergeConfigs(configs...), nil
}

// mergeConfigs merges multiple configurations, with values from the first map taking precedence.
func mergeConfigs(configs ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, config := range configs {
		for key, value := range config {
			if _, ok := merged[key]; !ok {
				merged[key] = value
			}
		}
	}
	return merged
}
