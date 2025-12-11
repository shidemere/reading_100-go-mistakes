package main

import (
	"log"
	"net/http"
)

func main() {
	_ = listing1()
	_ = listing2()
	_ = listing3()
	_ = listing4()
}

func listing1() error {
	var client *http.Client
	/*
	Происходит непреднамеренное затенение переменной client и тем не менее ошибки компиляции
	не случается потому что значение переменной выше используется для логгирования вне условия
	*/
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	_ = client
	return nil
}
// Есть два варианта решения этой проблемы.
// Первые реинициализация, где делается client = c
// Громоздкая хуйня как по мне
func listing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}
// Второй вариант это разделение создания и инициализации
// Но тогда создавать переменную ошибки надо тоже отдельно, так как нельзя разделить
// инициализацию и присвоение на два разных этапа раздельно для двух возвращаемых значений
// Этот вариант мне нравится больше 
func listing3() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

// Второй вариант позволяет упрощать flow в момент когда мы конструируем поток программы 
// в зависимости от каких то условий
func listing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
