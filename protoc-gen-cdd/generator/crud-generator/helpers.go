package crudgenerator

import (
	"strings"

	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
)

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
			tmpOut += " " + pk.GetGoStandartType()
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
			if fext.GetGoType() == "time.Time" {
				needImport = true
				break
			}
		}
	}
	return needImport
}
