package generator

type dbStruct struct {
	name   string
	fields []dbField
}

type dbField struct {
	name string
	kind string
}

func DbStruct(name string) *dbStruct {
	return &dbStruct{
		name:   name,
		fields: make([]dbField, 0),
	}
}

func (v *dbStruct) Field(dbName string, kind string) *dbStruct {
	v.fields = append(v.fields, dbField{
		name: dbName,
		kind: kind,
	})
	return v
}

func (v *dbStruct) IntoField() *Field {
	f := NewField(v.name, v.name, nil, nil)
	for _, field := range v.fields {
		f.withField(NewField(sqlToFieldName(field.name, true), field.kind, Tags().DB(field.name), nil))
	}
	return f
}
