package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello SayHello // boleh diganti dengan SayHello saja karena kalau nama sama tidak perlu ditulis type-nya
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}
