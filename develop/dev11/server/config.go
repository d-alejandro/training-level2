package server

func GetConfigs() map[string]any {
	return map[string]any{
		"http": map[string]string{
			"host": "localhost",
			"port": "8080",
		},
	}
}
