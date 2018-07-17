package utils

// Check checks and panics if the error is not nil
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
