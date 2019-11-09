package options

// Alias hides the real type of the enum
// and users can use it to define the var for accepting enum
type Alias = string

type list struct {
	Python Alias
	Flask  Alias
	Django Alias
}

// Enum for public use
var Enum = &list{
	Python: "python",
	Flask:  "flask",
	Django: "django",
}
