package objects

// Builtins contains all known builtin functions.
var Builtins = []struct {
	Name string
	Func CallableFunc
}{
	{
		Name: "print",
		Func: builtinPrint,
	},
	{
		Name: "len",
		Func: builtinLen,
	},
	{
		Name: "copy",
		Func: builtinCopy,
	},
	{
		Name: "append",
		Func: builtinAppend,
	},
	{
		Name: "string",
		Func: builtinString,
	},
	{
		Name: "int",
		Func: builtinInt,
	},
	{
		Name: "bool",
		Func: builtinBool,
	},
	{
		Name: "float",
		Func: builtinFloat,
	},
	{
		Name: "char",
		Func: builtinChar,
	},
	{
		Name: "is_int",
		Func: builtinIsInt,
	},
	{
		Name: "is_float",
		Func: builtinIsFloat,
	},
	{
		Name: "is_string",
		Func: builtinIsString,
	},
	{
		Name: "is_bool",
		Func: builtinIsBool,
	},
	{
		Name: "is_char",
		Func: builtinIsChar,
	},
	{
		Name: "is_error",
		Func: builtinIsError,
	},
	{
		Name: "is_undefined",
		Func: builtinIsUndefined,
	},
}
