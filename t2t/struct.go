package t2t

import (
	"reflect"
	"strings"
	"unicode"
)

type tagOptions struct {
	Tag string

	Inline    bool
	Omitempty bool
}

func (t *tagOptions) Parse(tag string) error {
	idx := strings.Index(tag, ",")
	if idx >= 0 {
		t.Tag = tag[:idx]
		t.parseOptions(tag[idx+1:])
	} else {
		t.Tag = tag
	}

	if !t.isValidTag(t.Tag) {
		t.Tag = ""
	}

	return nil
}

func (t *tagOptions) parseOptions(opt string) {
	if strings.Contains(opt, "inline") {
		t.Inline = true
	} else {
		t.Inline = false
	}

	if strings.Contains(opt, "omitempty") {
		t.Omitempty = true
	} else {
		t.Omitempty = false
	}
}

func (t *tagOptions) isValidTag(tag string) bool {
	if tag == "" {
		return false
	}
	for _, c := range tag {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
			// Backslash and quote chars are reserved, but
			// otherwise any punctuation chars are allowed
			// in a tag name.
		default:
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
				return false
			}
		}
	}
	return true
}

type structField struct {
	Name      string
	FieldName string

	PkgPath string

	Type       reflect.Type
	Anonymous  bool
	TagOptions tagOptions
}

func getStructFieldMap(rv reflect.Value, tagName string) (map[string]reflect.Value, error) {
	// TODO: cached
	result := map[string]reflect.Value{}
	var field structField
	var subFields map[string]reflect.Value
	var err error

	for i := 0; i < rv.NumField(); i++ {
		field, err = getStructField(rv.Type().Field(i), tagName)
		if err != nil {
			return nil, err
		}

		if field.Name == "-" {
			continue // Ignore "-" tag name
		}

		if len(field.PkgPath) != 0 {
			continue // Ignore unexported field
		}

		if field.TagOptions.Inline {
			subFields, err = getStructFieldMap(rv.Field(i), tagName)
			if err != nil {
				return nil, err
			}

			for k, v := range subFields {
				result[k] = v
			}
			continue
		}

		result[field.Name] = rv.Field(i)
	}

	return result, nil
}

func getStructField(f reflect.StructField, tagName string) (structField, error) {
	result := structField{}

	result.FieldName = f.Name
	result.Type = f.Type
	result.Anonymous = f.Anonymous
	result.PkgPath = f.PkgPath

	err := result.TagOptions.Parse(f.Tag.Get(tagName))
	if err != nil {
		return result, &structTagInvalidError{}
	}

	if len(result.TagOptions.Tag) != 0 {
		result.Name = result.TagOptions.Tag
	} else {
		result.Name = result.FieldName
	}

	return result, nil
}
