package presenter

type Presenter interface {
	Show() ([]byte, error)
	Bind(interface{}) error //modelar os dados que se quer enviar para utilizar no show
}