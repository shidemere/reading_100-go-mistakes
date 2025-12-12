package client

import "github.com/teivah/100-go-mistakes/src/02-code-project-organization/6-interface-producer/store"

/*
	Вот пример создания интерфейса на стороне потребителя. 
	Я знаю что мне нужно только получение и я реализую это получение. 
	Даже несмотря на то что у меня реализация на которую я делаю интерфейс лежит в соседнем  пакете - я все равно
	должен его импортировать в этот пакет и поэтому всё таки стоит даже в таких мелочах следовать этому правилу.
*/
type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
