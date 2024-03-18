package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello SayHello
}

type SayHelloImpl struct {
	Age string
}

func (s SayHelloImpl) Hello(name string) string {
	return "Hello " + name + ", My age is " + s.Age
}

// provider

func NewSayHelloImpl(Age string) *SayHelloImpl {
	return &SayHelloImpl{
		Age: Age,
	}
}

func NewHelloService(sayHello SayHello) *HelloService {
	// parameter sayHello-nya adalah  interface
	return &HelloService{
		SayHello: sayHello,
	}
}

// ketika tipe datanya berupa interface, maka yang dikirim harus berupa implementasi dari interface tersebut
// var sayhello = NewHelloService(NewSayHelloImpl("23"))
 