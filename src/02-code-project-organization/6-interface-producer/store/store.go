package store

/*
	Тут приводится пример плохого проектирования, потому что мы явно навязываем то, какую реализацию будет
	использовать потребитель этого интерфейса. 
	Иногда это полезно, как например в пункте про ограничение поведения, которое у меня было в прошлой ошибке.
	Однако если я не имею каких то требований по огранчению пользователя - я могу оставить это ему на откуп
	и дать ему возможность создать интерфейс на стороне потребителя
*/
type CustomerStorage interface {
	StoreCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	UpdateCustomer(customer Customer) error
	GetAllCustomers() ([]Customer, error)
	GetCustomersWithoutContract() ([]Customer, error)
	GetCustomersWithNegativeBalance() ([]Customer, error)
}

type Customer struct{}
