package helper

var version = "1.0.0" // hanya bisa di akses file ini
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
} // hanya bisa di akses di file ini

func SayHello(name string) string {
	return "Hello " + name
}
