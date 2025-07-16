package core

func CheckRequiredKeys(env map[string]string, required []string) []string {
	var missing []string

	for _, key := range required {
		if _, exists := env[key]; !exists {
			missing = append(missing, key)
		}
	}

	return missing
}
