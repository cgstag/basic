package account

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRandomAccount(t *testing.T) {
	account, err := NewRandomAccount()
	if nil != err || account == nil {
		t.Error("Error creating Random Account")
	} else {
		fmt.Println(reflect.TypeOf(account.Balance.Consolidated))
		if reflect.TypeOf(account.Balance.Consolidated).String() != "float64" {
			t.Error("Expected Balance.Consolidated as a float64")
		}
		if account.AccountSegment.SegmentType != "Varejo" &&
			account.AccountSegment.SegmentType != "Uniclass" &&
			account.AccountSegment.SegmentType != "Personalité" {
			t.Error("Expected segment as a String")
		}
	}
}
