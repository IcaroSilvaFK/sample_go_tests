package functions

import (
	"github.com/IcaroSilvaFK/sample_tests_go/functions/utils"
)

func Sum(numA, numB uint) uint {
	return numA + numB
}

func IsPalatial(str string) bool {

	return str == utils.Reverse(str)

}
