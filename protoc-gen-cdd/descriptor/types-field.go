package descriptor

import (
	"fmt"
	"log"
	"strings"

	cddext "github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/cdd/api"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FieldDescriptorExt struct {
	*descriptorpb.FieldDescriptorProto
	Repository      *DescriptorRepository
	MessageExt      *MessageDescriptorExt
	DBField         *cddext.DBField
	ValidationRules []string
	DefaultValueExt string
}

func (FieldDescriptorExt) New(msgext *MessageDescriptorExt, field *descriptorpb.FieldDescriptorProto) *FieldDescriptorExt {
	fieldext := &FieldDescriptorExt{
		FieldDescriptorProto: field,
		Repository:           msgext.Repository,
		MessageExt:           msgext,
		DBField:              nil,
		ValidationRules:      []string{},
		DefaultValueExt:      "",
	}
	fieldext.DBField = parseExtDBField(field)
	if fieldext.DBField == nil {
		fieldext.DBField = &cddext.DBField{ColumnName: fieldext.GetJsonName(), PrimaryKey: false, ColumnType: ""}
	} else if fieldext.DBField.ColumnName == "" {
		fieldext.DBField.ColumnName = fieldext.GetJsonName()
	}

	validationRule := parseExtFieldValidation(field)
	if validationRule != "" {
		fieldext.ValidationRules = strings.Split(validationRule, "|")
	}

	fieldext.DefaultValueExt = parseExtFieldDefault(field)

	return fieldext
}

func (fieldext *FieldDescriptorExt) GetIdentifier() string {
	return fieldext.MessageExt.GetIdentifier() + "." + fieldext.GetName()
}

func parseExtDBField(field *descriptorpb.FieldDescriptorProto) *cddext.DBField {
	if field.Options == nil {
		return nil
	} else if !proto.HasExtension(field.Options, cddext.E_Dbfield) {
		return nil
	}

	ext := proto.GetExtension(field.Options, cddext.E_Dbfield)
	opts, ok := ext.(*cddext.DBField)
	if !ok {
		log.Println(fmt.Errorf("[parseExtDBField] extension is %T; want an DBField", ext))
		return nil
	}
	return opts
}

func parseExtFieldValidation(field *descriptorpb.FieldDescriptorProto) string {
	if field.Options == nil {
		return ""
	} else if !proto.HasExtension(field.Options, cddext.E_Validate) {
		return ""
	}

	ext := proto.GetExtension(field.Options, cddext.E_Validate)
	opts, ok := ext.(string)
	if !ok {
		log.Println(fmt.Errorf("[parseExtFieldValidation] extension is %T; want an string", ext))
		return ""
	}
	return opts
}

func parseExtFieldDefault(field *descriptorpb.FieldDescriptorProto) string {
	if field.Options == nil {
		return ""
	} else if !proto.HasExtension(field.Options, cddext.E_Default) {
		return ""
	}

	ext := proto.GetExtension(field.Options, cddext.E_Default)
	opts, ok := ext.(string)
	if !ok {
		log.Println(fmt.Errorf("[parseExtFieldDefault] extension is %T; want an string", ext))
		return ""
	}
	return opts
}
