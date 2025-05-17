package lexutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var lexTypesMap map[string]reflect.Type

func init() {
	lexTypesMap = make(map[string]reflect.Type)
	RegisterType("blob", &LexBlob{})
}

func RegisterType(id string, val interface{}) {

	t := reflect.TypeOf(val)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if _, ok := lexTypesMap[id]; ok {
		panic(fmt.Sprintf("already registered type for %q", id))
	}

	lexTypesMap[id] = t
}

func NewFromType(typ string) (interface{}, error) {
	t, ok := lexTypesMap[typ]
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnrecognizedType, typ)
	}
	v := reflect.New(t)
	return v.Interface(), nil
}

func JsonDecodeValue(b []byte) (any, error) {
	tstr, err := TypeExtract(b)
	if err != nil {
		return nil, err
	}

	t, ok := lexTypesMap[tstr]
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnrecognizedType, tstr)
	}

	val := reflect.New(t)

	ival := val.Interface()
	if err := json.Unmarshal(b, ival); err != nil {
		return nil, err
	}

	return ival, nil
}

var ErrUnrecognizedType = fmt.Errorf("unrecognized lexicon type")

type LexiconTypeDecoder struct {
	Val interface{}
}

func (ltd *LexiconTypeDecoder) UnmarshalJSON(b []byte) error {
	val, err := JsonDecodeValue(b)
	if err != nil {
		return err
	}

	ltd.Val = val

	return nil
}

func (ltd *LexiconTypeDecoder) MarshalJSON() ([]byte, error) {
	if ltd == nil || ltd.Val == nil {
		return nil, fmt.Errorf("LexiconTypeDecoder MarshalJSON called on a nil")
	}
	v := reflect.ValueOf(ltd.Val)
	t := v.Type()
	sf, ok := t.Elem().FieldByName("LexiconTypeID")
	if !ok {
		return nil, fmt.Errorf("lexicon type decoder can only handle record fields")
	}

	tag, ok := sf.Tag.Lookup("cborgen")
	if !ok {
		return nil, fmt.Errorf("lexicon type decoder can only handle record fields with const $type")
	}

	parts := strings.Split(tag, ",")

	var cval string
	for _, p := range parts {
		if strings.HasPrefix(p, "const=") {
			cval = strings.TrimPrefix(p, "const=")
			break
		}
	}
	if cval == "" {
		return nil, fmt.Errorf("must have const $type field")
	}

	v.Elem().FieldByName("LexiconTypeID").SetString(cval)

	return json.Marshal(ltd.Val)
}
