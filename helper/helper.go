package helper

// Int return pointer to int
func Int(i int) *int {
	return &i
}

// String return pointer to string
func String(s string) *string {
	return &s
}
