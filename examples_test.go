package wren_test

import wren "bitbucket.org/eridenmk/dbl-konstruct/go-wren"

func ExampleWrenVM_Interpret() {
	const program = `
		System.print("Hello from Wren!")
	`

	vm := wren.NewVM()
	if err := vm.Interpret("main", program); err != nil {
		panic(err)
	}
}
