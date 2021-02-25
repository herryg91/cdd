package crudgenerator

import (
	"strings"

	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
)

func isMessagesNeedImportTime(mexts []*descriptor.MessageDescriptorExt) bool {
	for _, mext := range mexts {
		if isMessageNeedImportTime(mext) {
			return true
		}
	}

	return false
}
func isMessageNeedImportTime(mext *descriptor.MessageDescriptorExt) bool {
	needImport := !mext.DBSchema.DisableTimestampTracking || mext.DBSchema.EnableSoftDelete
	if !needImport {
		for _, fext := range mext.FieldExt {
			if getGoType(fext) == "time.Time" {
				needImport = true
				break
			}
		}
	}
	return needImport
}

func getPrimaryKey(fieldexts []*descriptor.FieldDescriptorExt) []*descriptor.FieldDescriptorExt {
	fieldpks := []*descriptor.FieldDescriptorExt{}
	for _, f := range fieldexts {
		if f.DBField.PrimaryKey {
			fieldpks = append(fieldpks, f)
		}
	}
	return fieldpks
}

func getPrimaryKeyAsString(fieldexts []*descriptor.FieldDescriptorExt, prefix, suffix, delimiter string, toLower bool, withGoType bool) string {
	fieldpks := getPrimaryKey(fieldexts)
	out := []string{}
	for _, pk := range fieldpks {
		pkName := pk.GetName()
		if toLower {
			pkName = strings.ToLower(pkName)
		}
		tmpOut := prefix + pkName
		if withGoType {
			tmpOut += " " + getGoStandartType(pk)
		}
		tmpOut += suffix
		out = append(out, tmpOut)
	}
	return strings.Join(out, delimiter)
}

func getPrimaryKeyAsQueryStmt(fieldexts []*descriptor.FieldDescriptorExt) string {
	fieldpks := getPrimaryKey(fieldexts)
	out := []string{}
	for _, pk := range fieldpks {
		out = append(out, pk.DBField.ColumnName+" = ?")
	}
	return strings.Join(out, " AND ")
}

func getGormTagAttribute(fieldext *descriptor.FieldDescriptorExt) string {
	result := ""
	if fieldext.DBField.PrimaryKey {
		result = "primary_key"
	}
	if fieldext.DBField.ColumnName != "" {
		if result != "" {
			result += ";"
		}
		result += "column:" + fieldext.DBField.ColumnName
	}
	if fieldext.DBField.ColumnType != "" {
		if result != "" {
			result += ";"
		}
		result += "type:" + fieldext.DBField.ColumnType
	}
	return result
}

func getGoType(fieldext *descriptor.FieldDescriptorExt) string {
	t := ""
	if fieldext.GetTypeName() == "" {
		t = strings.ToLower(strings.Replace(fieldext.GetType().String(), "TYPE_", "", -1))
		if t == "double" {
			t = "float64"
		}
	} else {
		switch fieldext.GetTypeName() {
		case ".google.protobuf.Timestamp":
			t = "time.Time"
		}
	}
	return t
}

func getGoStandartType(fieldext *descriptor.FieldDescriptorExt) string {
	t := getGoType(fieldext)
	switch t {
	case "float32":
		return "float64"
	case "float64":
		return t
	case "int":
		return t
	case "int8":
		return "int"
	case "int16":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "uint":
		return t
	case "uint16":
		return "uint"
	case "uint32":
		return "uint"
	case "uint64":
		return "uint"
	}
	return t
}
