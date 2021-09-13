package internal

import (
	"github.com/stretchr/testify/assert"
	asserterImpl "lisp_lambda-go/internal/asserter/impl"
	evaluatorImpl "lisp_lambda-go/internal/evaluator/impl"
	functionImpl "lisp_lambda-go/internal/function/impl"
	interpreterImpl "lisp_lambda-go/internal/interpreter/impl"
	parserImpl "lisp_lambda-go/internal/parser/impl"
	printerImpl "lisp_lambda-go/internal/printer/impl"
	tokenizerImpl "lisp_lambda-go/internal/tokenizer/impl"
	userDefinedImpl "lisp_lambda-go/internal/userdefined/impl"
	"testing"
)

var tests = []struct {
	input string
	expected string
}{
	//ATOMIC tests
	{"T", "T"},
	{"NIL", "NIL"},
	{"9873", "9873"},

	//ATOM tests
	{"(atom (cdr (cons 34 54)))" , " T"},
	{"(atom (null (int (cons 34 54))))" , " T"},
	{"(atom -4)" , " T"},
	{"(atom (atom 45))" , " T"},
	{"(atom (cons 2 9))" , " NIL"},
	{"(atom (atom (atom T)))" , " T"},
	{"(atom ('(7 10)))" , " NIL"},
	{"(atom (' ZZ))" , " T"},
	//CAR tests
	{"(car (cons 2 45))" , " 2"},
	{"(car (cons (car (cons 56 43)) T))" , " 56"},
	//CDR tests
	{"(cdr (cons 2 45))" , " 45"},
	{"(cdr (cons (cdr (cons 56 43)) T))" , " T"},
	{"(cdr ('(4 (cons 45 6))))" , " ((cons 45 6))"},
	//COND tests
	{"(cond ((< 45 3) 1) (T 2) (34 45))" , " 2"},
	{"(cond (NIL 54) (NIL 34) (34 78) (645 234))" , " 78"},
	{"(cond (T 23))" , " 23"},
	{"(cond (NIL 3) (T 2))" , " 2"},
	{"(cond (4 34) ((+ T 4) 5))" , " 34"},
	//CONS tests
	{"(cons (+ 2 3)(cons 8 (null 5)))" , " (5 8)"},
	{"(cons 2 (cons 3 (cons 4 5)))" , " (2 3 4 . 5)"},
	{"(cons (cons 1 2) (cons 3 4))" , " ((1 . 2) 3 . 4)"},
	{"(cons (= 1 2) (= 3 3))" , " (NIL . T)"},
	{"(cons (car ('(7 8))) (cdr ('(6 10))))" , " (7 10)"},
	//EQ tests
	{"(= 23 23)" , " T"},
	{"(= (= 34 34) (= (= 23 23) NIL))" , " NIL"},
	{"(= T T)" , " T"},
	{"(= 2 (+ 1 1))" , " T"},
	{"(= (' XYZ1) (' XYZ1))" , " T"},
	{"(= (' A) (' B))" , " NIL"},
	{"(= (> 34 2) T)" , " T"},
	{"(= (int NIL) (null T))" , " T"},
	{"(= (cons 23 ()) 9)" , " NIL"},
	{"(= 23 T (cons 23 1))" , " NIL"},
	{"(= 1)" , " T"},
	{"(= -133)" , " T"},
	//EVEN_P tests
	{"(evenp 0)" , " T"},
	{"(evenp -34)" , " T"},
	{"(evenp 23)" , " NIL"},
	{"(evenp 55556)" , " T"},
	//EXP tests
	{"(exp 1)" , " 2"},
	{"(exp 2)" , " 7"},
	{"(exp 3)" , " 20"},
	//EXPT tests
	{"(expt 1 3)" , " 1"},
	{"(expt 2 3)" , " 8"},
	{"(expt 3 2)" , " 9"},
	//GCD tests
	{"(gcd)" , " 0"},
	{"(gcd 4)" , " 4"},
	{"(gcd 91 49)" , " 7"},
	{"(gcd 91 -49)" , " 7"},
	{"(gcd 63 42 35)" , " 7"},
	{"(gcd 81 153)" , " 9"},
	//GREATER tests
	{"(> 98 1)" , " T"},
	{"(> (+ 3 2) 6)" , " NIL"},
	{"(> (cdr (cons (+ 12 12) 6)) (- 13 19))" , " T"},
	{"(> 23 45 98 34)" , " NIL"},
	{"(> 45 32 21 5)" , " T"},
	{"(> 69)" , " T"},
	//GREATER_EQ tests
	{"(>= 3)" , " T"},
	{"(>= 5 5 4)" , " T"},
	{"(>= 4 4 5)" , " NIL"},
	{"(>= 1 1 1 1 1 1)" , " T"},
	{"(>= 1 1 1 1 1 -134)" , " T"},
	//INT tests
	{"(int 3)" , " T"},
	{"(int (int 45))" , " NIL"},
	{"(int (+ 4 5))" , " T"},
	{"(int (' X23))" , " NIL"},
	{"(int (= 34 34))" , " NIL"},
	{"(int (cons 4 5))" , " NIL"},
	//LCM tests
	{"(lcm 14 35)" , " 70"},
	{"(lcm 1 2 3 4 5 6)" , " 60"},
	{"(lcm 0 5)" , " 0"},
	{"(lcm -34)" , " 34"},
	//LESS tests
	{"(< 1 19)" , " T"},
	{"(< 34)" , " T"},
	{"(< (- 34 3) 1)" , " NIL"},
	{"(< (* 12 32) (cond ((null (int T)) 3)))" , " NIL"},
	//LESS_EQ tests
	{"(<= 3)" , " T"},
	{"(<= 5 5 4)" , " NIL"},
	{"(<= 4 4 5)" , " T"},
	{"(<= 1 1 1 1 1 1)" , " T"},
	//MAX tests
	{"(max 6 12)" , " 12"},
	{"(max 3)" , " 3"},
	{"(max 2 3 0 7)" , " 7"},
	//MIN tests
	{"(min 6 12)" , " 6"},
	{"(min 3)" , " 3"},
	{"(min 2 3 0 7)" , " 0"},
	{"(min 2 3 0 -7)" , " -7"},
	//MINUS tests
	{"(- 1 -13)" , " 14"},
	{"(- 13 (- (- 30 23) 7))" , " 13"},
	{"(- 45)" , " -45"},
	{"(- 3 4 5)" , " -6"},
	//MINUS_P tests
	{"(minusp 0)" , " NIL"},
	{"(minusp -34)" , " T"},
	{"(minusp 23)" , " NIL"},
	//NOT_EQ tests
	{"(/= T T)" , " NIL"},
	{"(/= 2 (+ 2 1))" , " T"},
	{"(/= 45)" , " T"},
	{"(/= 45 -34 12 1)" , " T"},
	//NULL tests
	{"(null NIL)" , " T"},
	{"(null (null NIL))" , " NIL"},
	{"(null ('(A)))" , " NIL"},
	{"(null (= 2 (+ 1 1)))" , " NIL"},
	{"(null (int (int (+ 23 23))))" , " T"},
	//ODD_P tests
	{"(oddp 0)" , " NIL"},
	{"(oddp -35)" , " T"},
	{"(oddp 23)" , " T"},
	{"(oddp 98)" , " NIL"},
	//ONE_MINUS tests
	{"(1- 0)" , " -1"},
	{"(1- 45)" , " 44"},
	//ONE_PLUS tests
	{"(1+ 0)" , " 1"},
	{"(1+ 456)" , " 457"},
	//PLUS tests
	{"(+)" , " 0"},
	{"(+ 8 4)" , " 12"},
	{"(+ 4 5 -3)" , " 6"},
	{"(+ (+ 4 3) (+ (+ 1 2) 4))" , " 14"},
	{"(+ 1 2)" , " 3"},
	{"(+ (+ 3 5) (* 4 4))" , " 24"},
	{"(+ (car (cons 2 0)) (car (cons 9 3)))" , " 11"},
	{"(+ (+ 3 5) (car ('(7 8))))" , " 15"},
	{"(+ (cond ((> 3 45) 1) (T 2)) 2854)" , " 2856"},
	{"(+ (cond ((int T) 34) ((null 2) 1) ((int 12) 12) (T 3)) 0)" , " 12"},
	//PLUS_P tests
	{"(plusp 0)" , " NIL"},
	{"(plusp -34)" , " NIL"},
	{"(plusp 23)" , " T"},
	//QUOTE tests
	{"('(3 4 5))" , " (3 4 5)"},
	{"('(cons 34 92))" , " (cons 34 92)"},
	//TIMES tests
	{"(* (- 0 1) 45)" , " -45"},
	{"(* 3 -34)" , " -102"},
	{"(* (* 1 2) (* 2 (* 5 6)))" , " 120"},
	{"(* 23 4 2)" , " 184"},
	{"(*)" , " 1"},
	//ZERO_P tests
	{"(zerop 0)" , " T"},
	{"(zerop -0)" , " T"},
	{"(zerop 23)" , " NIL"},
}

func TestRegression(t *testing.T) {
	asserterSingleton := asserterImpl.NewSingleton()

	tokenizerSingleton := tokenizerImpl.NewSingleton()
	parserSingleton := parserImpl.NewSingleton()
	printerSingleton := printerImpl.NewSingleton()

	functionSingleton := functionImpl.NewSingleton()
	userDefinedSingleton := userDefinedImpl.NewSingleton(
		functionSingleton.GetFunctionMap(),
		asserterSingleton.GetFunctionLengthAsserter(),
	)
	evaluatorSingleton := evaluatorImpl.NewSingleton(
		functionSingleton.GetFunctionMap(),
	)
	interpreterSingleton := interpreterImpl.NewSingleton(
		tokenizerSingleton.GetTokenizer(),
		parserSingleton.GetParser(),
		userDefinedSingleton.GetFunctionGenerator(),
		asserterSingleton.GetExpressionListLengthAsserter(),
		evaluatorSingleton.GetProgramEvaluator(),
		printerSingleton.GetPrinter(),
	)

	interpreter := interpreterSingleton.GetInterpreter()


	for _, e  := range tests {

		actual, actualError := interpreter.Interpret(e.input)
		assert.Equal(t, e.expected, actual)
		assert.Nil(t, actualError)
	}
}
