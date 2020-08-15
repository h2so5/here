package here

// Shortcut here.Doc.
func D(raw string) string {
	return Doc(raw)
}

// Shortcut here.Docf.
func Df(raw string, args ...interface{}) string {
	return Docf(raw, args...)
}
