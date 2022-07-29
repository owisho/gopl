package expand

func expand(s string, f func(string) string) string {
	return f(s)
}
