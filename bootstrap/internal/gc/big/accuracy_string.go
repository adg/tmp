// Do not edit. Bootstrap copy of /Users/rsc/g/go/src/cmd/internal/gc/big/accuracy_string.go

// generated by stringer -type=Accuracy; DO NOT EDIT

package big

import "fmt"

const _Accuracy_name = "BelowExactAbove"

var _Accuracy_index = [...]uint8{0, 5, 10, 15}

func (i Accuracy) String() string {
	i -= -1
	if i < 0 || i+1 >= Accuracy(len(_Accuracy_index)) {
		return fmt.Sprintf("Accuracy(%d)", i+-1)
	}
	return _Accuracy_name[_Accuracy_index[i]:_Accuracy_index[i+1]]
}
