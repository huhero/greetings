# Saludos en Go

Este paquete proporciona una forma simple de obtener saludos.

## Instalación
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/huhero/greetings
```

## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
	"fmt"
	"log"

	"github.com/huhero/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Hugo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```
