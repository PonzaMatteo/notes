package oop

// the type is exported
type Foo struct {
	// Capitalized field: it is exported
	Exported int
	// field is not visible out of main package
	private int
}

func IsPublic()    {}
func isNotPublic() {}
