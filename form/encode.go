package form

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
)

type File struct {
	Name    string
	Content io.Reader
}

type param interface {
	WriteTo(*multipart.Writer) error
}

type arg struct {
	name  string
	value string
}

func (a arg) WriteTo(w *multipart.Writer) error {
	fw, err := w.CreateFormField(a.name)
	if err != nil {
		return err
	}
	fmt.Fprint(fw, a.value)
	fw.Write([]byte{}) // fails without this?
	return nil
}

type file struct {
	name     string
	fileName string
	content  io.Reader
}

func (f file) WriteTo(w *multipart.Writer) error {
	fw, err := w.CreateFormFile(f.name, f.fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, f.content)
	if err != nil {
		return err
	}
	fw.Write([]byte{}) // fails without this?
	return nil
}

// Encode encodes the struct s to a multpart form returning the request
// body, the content type (including boundary) for the form or an error
func Encode(s interface{}) ([]byte, string, error) {
	return buildForm(buildParams(s))
}

func buildParams(s interface{}) (params []param) {
	v := reflect.ValueOf(s)
	t := v.Type()
	switch k := t.Kind(); k {
	case reflect.Ptr:
		return buildParams(v.Elem().Interface())
	case reflect.Struct:
		break
	default:
		return params
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		name, omitEmpty := fieldInfo(f)

		if name == "-" {
			continue
		}
		fv := v.Field(i)
		if omitEmpty && isEmpty(fv) {
			continue
		}
		if fv.CanInterface() {
			f, _ := fv.Interface().(*File)
			if f == nil {
				fs, ok := fv.Interface().(File)
				if ok {
					f = &fs
				}
			}
			if f != nil {
				params = append(params, file{
					name:     name,
					fileName: f.Name,
					content:  f.Content,
				})
				continue
			}
		}
		params = append(params, arg{
			name:  name,
			value: encodeValue(fv),
		})
	}
	return params
}

func buildForm(params []param) ([]byte, string, error) {
	var out bytes.Buffer
	writer := multipart.NewWriter(&out)
	for _, param := range params {
		if err := param.WriteTo(writer); err != nil {
			return nil, "", err
		}
	}
	writer.Close()
	return out.Bytes(), writer.FormDataContentType(), nil
}

func encodeValue(v reflect.Value) string {
	t := v.Type()
	switch k := t.Kind(); k {
	case reflect.Ptr:
		return encodeValue(v.Elem())
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'g', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		s := fmt.Sprintf("%g", v.Complex())
		return strings.TrimSuffix(strings.TrimPrefix(s, "("), ")")
	case reflect.String:
		return v.String()
	}
	panic(t.String() + " has unsupported kind " + t.Kind().String())
}

func fieldInfo(f reflect.StructField) (name string, omitEmpty bool) {
	if f.PkgPath != "" { // Skip private fields
		name = "-"
		return
	}
	name = f.Name
	tag := f.Tag.Get("form")
	if tag == "" {
		return
	}

	tagBits := strings.SplitN(tag, ",", 2)
	if tagBits[0] != "" {
		name = tagBits[0]
	}
	if len(tagBits) == 2 && tagBits[1] == "omitempty" {
		omitEmpty = true
	}
	return
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return v.Int() == 0
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
