package wren_test

import (
	"fmt"

	wren "bitbucket.org/eridenmk/dbl-konstruct/go-wren"
)

func Example_callWren() {
	const program = `
		class Bird {
			static fly(where) {
				return "Flying to %(where)!"
			}
		}
	`

	vm := wren.NewVM()
	if err := vm.Interpret("main", program); err != nil {
		panic(err)
	}

	response, err := vm.Variable("Bird").Call("fly(_)", "Chicago")
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
