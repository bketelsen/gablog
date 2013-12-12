package main

// START1 OMIT
func Process(prefix string, spec interface{}) error {
	// START2 OMIT
	s := reflect.ValueOf(spec).Elem()
	if s.Kind() != reflect.Struct {
		return ErrInvalidSpecification
	}
	// END2 OMIT
	typeOfSpec := s.Type()
	// START3 OMIT
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.CanSet() {
			fieldName := typeOfSpec.Field(i).Name
			key := fmt.Sprintf("%s_%s", prefix, fieldName)
			value := os.Getenv(strings.ToUpper(key))
			// END3 OMIT
			if value == "" {
				continue
			}
			// START4 OMIT
			switch f.Kind() {
			case reflect.String:
				f.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(value, 0, f.Type().Bits())
				if err != nil {
					return &ParseError{
						FieldName: fieldName,
						TypeName:  f.Kind().String(),
						Value:     value,
					}
				}
				f.SetInt(intValue)
			}
			// END4 OMIT
		}
	}
}
// END1 OMIT
