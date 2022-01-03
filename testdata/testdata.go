package testdata

var testdataInt = 1
var (
	testdataInt2 = 2
)

const testdataString = "a"
const (
	testdataString2 = "aa"
)

func foo() string {
	var testdataVar string
	var testdataVar2 string = testdataVar
	var (
		testdataVar3 = testdataVar2
	)
	testdataVariableStringInFunc := testdataString2
	testdataVariableStringInFunc += "a"
	return testdataVariableStringInFunc + testdataVar3
}
func bar() int {
	const testdataConst = 2
	const testdataConst2 = testdataConst
	const (
		testdataConst3 = testdataConst2
	)
	testdataVariableIntInFunc := 0
	return testdataVariableIntInFunc + testdataConst3
}
