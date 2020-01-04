package pkg

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

var convertInterfaceFailed = errors.New("failed to convert interface to []interface")

func ToInterfaceSlice(i interface{}) ([]interface{}, error) {
	if reflect.TypeOf(i).Kind() != reflect.Slice {
		return nil, convertInterfaceFailed
	}

	s := reflect.ValueOf(i)
	var slice []interface{}
	for i := 0; i < s.Len(); i++ {
		slice = append(slice, s.Index(i).Interface())
	}

	return slice, nil
}

// GetElementValue はrefrect.Valueを必ず値にして返す関数です。
func GetElementValue(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Ptr:
		return GetElementValue(v.Elem())
	default:
		return v
	}
}

// FindField は名前からフィールドを検索する関数です。
// 例: AssignmentItemからAssignmentItemAnswers[0].Answerを探す場合
// FindField(reflect.ValueOf(item), "AssignmentItemAnswers[0].Answer")
// => AssignmentItemAnswers[0].Answerのreflect.Valueが帰ってきます。
func FindField(v reflect.Value, fullName string) (reflect.Value, error) {
	if v.Kind() == reflect.Ptr {
		return FindField(v.Elem(), fullName)
	}

	dot := strings.Index(fullName, ".")
	left := strings.Index(fullName, "[")
	if dot == -1 && left == -1 {
		if v.Kind() != reflect.Struct {
			return reflect.ValueOf(nil), fmt.Errorf("error")
		}

		f := v.FieldByName(fullName)
		if !f.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("no field name '%s' in '%v'", fullName, v.Type())
		}

		return f, nil
	}

	if dot < left || left == -1 {
		f := v.FieldByName(fullName[:dot])
		if !f.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("no field name '%s' in '%v'", fullName[:dot], v.Type())
		}

		return FindField(f, fullName[dot+1:])
	} else {
		name := fullName[:left]
		right := strings.Index(fullName, "]")
		if right < 0 {
			return reflect.ValueOf(nil), fmt.Errorf("could not find ']': %s", fullName[left+1:])
		}

		f := v.FieldByName(name)
		if !f.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("no field name '%s' in '%v'", name, v.Type())
		}

		if f.Kind() != reflect.Slice {
			return reflect.ValueOf(nil), fmt.Errorf("%s is not type of slice. use %s instead of %s[]", name, name, name)
		}

		index := cast.ToInt(fullName[left+1 : right])

		if index >= f.Len() {
			return reflect.ValueOf(nil), errors.New("slice index out of range")
		}

		return FindField(f.Index(index), strings.TrimPrefix(fullName[right+1:], "."))
	}
}
