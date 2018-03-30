package dockerlang

const (
	ADDITION_OPERATOR       = "+"
	SUBTRACTION_OPERATOR    = "†"
	MULTIPLICATION_OPERATOR = "*"
	DIVISION_OPERATOR       = "‡"
	MODULO_OPERATOR         = "%"
	VARIABLE_INITIALIZATION = "≡"
	VARIABLE_ASSIGNMENT     = "="
	EXIT_OPERATOR           = "ꙮ"
	NOOP                    = "NOOP"

	R_PAREN_PUNCTION    = "("
	L_PAREN_PUNCUTATION = ")"
)

var (
	OP_TO_ARITY = map[string]int{
		ADDITION_OPERATOR:       2,
		SUBTRACTION_OPERATOR:    2,
		MULTIPLICATION_OPERATOR: 2,
		DIVISION_OPERATOR:       2,
		MODULO_OPERATOR:         2,
		VARIABLE_INITIALIZATION: 1,
		VARIABLE_ASSIGNMENT:     1,
		EXIT_OPERATOR:           1,
		NOOP:                    1,
	}
)

// all the language-defined tokens in dockerlang
type Symbols struct {
	Operators   []string
	Keywords    []string
	Punctuation []string
}

func PopulateSymbols() *Symbols {
	return &Symbols{
		Operators: []string{
			ADDITION_OPERATOR,
			SUBTRACTION_OPERATOR,
			MULTIPLICATION_OPERATOR,
			DIVISION_OPERATOR,
			MODULO_OPERATOR,
			EXIT_OPERATOR,
			NOOP,
		},
		Keywords: []string{},
		Punctuation: []string{
			R_PAREN_PUNCTION,
			L_PAREN_PUNCUTATION,
		},
	}
}
