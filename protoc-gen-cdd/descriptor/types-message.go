package descriptor

import (
	"fmt"
	"log"
	"strings"

	cddext "github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/cdd/api"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type MessageDescriptorExt struct {
	*descriptorpb.DescriptorProto
	Repository *DescriptorRepository
	FileExt    *FileDescriptorExt
	NestedPath []string
	FieldExt   []*FieldDescriptorExt
	DBSchema   *cddext.DBSchema
}

func (MessageDescriptorExt) New(fext *FileDescriptorExt, msg *descriptorpb.DescriptorProto, nestedPath []string) *MessageDescriptorExt {
	msgext := &MessageDescriptorExt{
		DescriptorProto: msg,
		Repository:      fext.Repository,
		FileExt:         fext,
		NestedPath:      nestedPath,
		FieldExt:        []*FieldDescriptorExt{},
		DBSchema:        nil,
	}
	for _, field := range msgext.Field {
		fieldext := FieldDescriptorExt{}.New(msgext, field)
		msgext.FieldExt = append(msgext.FieldExt, fieldext)
	}
	msgext.DBSchema = parseExtDBSchema(msg)
	if msgext.DBSchema == nil {
		msgext.DBSchema = &cddext.DBSchema{IsDbSchema: false}
	}
	return msgext
}

func (msgext *MessageDescriptorExt) GetIdentifier() string {
	components := []string{""}
	if msgext.FileExt.Package != nil {
		components = append(components, msgext.FileExt.GetPackage())
	}
	components = append(components, msgext.NestedPath...)
	components = append(components, msgext.GetName())
	return strings.Join(components, ".")
}

func parseExtDBSchema(message *descriptorpb.DescriptorProto) *cddext.DBSchema {
	if message.Options == nil {
		return nil
	} else if !proto.HasExtension(message.Options, cddext.E_Dbschema) {
		return nil
	}

	ext := proto.GetExtension(message.Options, cddext.E_Dbschema)
	opts, ok := ext.(*cddext.DBSchema)
	if !ok {
		log.Println(fmt.Errorf("[parseExtDBSchema] extension is %T; want an DBSchema", ext))
		return nil
	}
	return opts
}
