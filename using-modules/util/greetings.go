package greet

func SayHello(name string) string {
	if name == "" {
		return "Hello, World!"
	}

	return "Hello, " + name + "!"
}
