package renderer

func getFloat(values map[string]any, key string) float64 {
	v, exists := values[key]
	if !exists {
		return 0
	}

	f, ok := v.(float64)
	if !ok {
		return 0
	}

	return f
}
