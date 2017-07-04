// copytagged.go - copy only tagged struct fields.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to copytagged, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package copytagged

import (
	"errors"
	"reflect"
	"strings"

	"github.com/oleiade/reflections"
)

func CopyTagged(dst, src interface{}, tagkey, tag string) error {
	t := reflect.TypeOf(src).Elem()
	dt := reflect.TypeOf(dst).Elem()
	if t.Kind() != reflect.Struct || t.Kind() != reflect.Struct {
		return errors.New("wrong type of inputs")
	}
	if t.Name() != dt.Name() {
		return errors.New("types of structs do not match")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		srctag := field.Tag.Get(tagkey)
		if strings.Contains(srctag, tag) {
			v, err := reflections.GetField(src, field.Name)
			if err != nil {
				return err
			}
			err = reflections.SetField(dst, field.Name, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
