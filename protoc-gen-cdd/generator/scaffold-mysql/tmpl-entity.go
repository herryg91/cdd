package scaffold_mysql

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplEntityHeader = template.Must(template.New("entity-header").Funcs(template.FuncMap{}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}

	package entity
	
	{{ if .NeedImportTime }}
	import (
		"time"
	)
	{{ end }}
	
	`))

	tmplEntityBody = template.Must(template.New("entity-body").Funcs(template.FuncMap{
		"getGormTagAttribute": getGormTagAttribute,
		"getGoStandartType":   getGoStandartType,
	}).Parse(`
	// Mysql Table: {{.Mysql.TableName}}
	type {{.GetName}} struct {
		{{- range $fext := .FieldExt}}
		{{$fext.GetName}} {{ if $fext.MysqlField.Nullable }}*{{ end }}{{getGoStandartType $fext}} ` + "`" + `gorm:"{{getGormTagAttribute $fext}}"` + "`" + `
		{{- end}}
		
		{{if not .Mysql.DisableTimestampTracking}}
		CreatedAt time.Time ` + "`" + `gorm:"column:created_at"` + "`" + `
		UpdatedAt time.Time ` + "`" + `gorm:"column:updated_at"` + "`" + `
		{{end}}
	}
	`))
)

func applyTemplateEntity(sm []ScaffoldMysql, fext *descriptor.FileDescriptorExt) (*generator.GeneratorResponseFile, error) {
	header, err := applyTemplateEntityHeader(sm, fext)
	if err != nil {
		return nil, err
	}
	bodies := []string{}
	for _, s := range sm {
		body, err := applyTemplateEntityBody(s)
		if err != nil {
			return nil, err
		}
		bodies = append(bodies, body)
	}

	formatted, err := format.Source([]byte(header + fmt.Sprintln() + strings.Join(bodies, fmt.Sprintln())))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     "entity/" + fext.GetPackage() + ".cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}

func applyTemplateEntityHeader(sm []ScaffoldMysql, fext *descriptor.FileDescriptorExt) (string, error) {
	w := bytes.NewBuffer(nil)

	needImportTime := false
	for _, s := range sm {
		needImportTime = needImportTime || s.NeedImportTime()
	}

	var tmplData = struct {
		FileExt        *descriptor.FileDescriptorExt
		NeedImportTime bool
	}{
		fext,
		needImportTime,
	}

	if err := tmplEntityHeader.Execute(w, tmplData); err != nil {
		return "", err
	}

	return w.String(), nil
}

func applyTemplateEntityBody(sm ScaffoldMysql) (string, error) {
	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		ScaffoldMysql
	}{
		sm,
	}

	if err := tmplEntityBody.Execute(w, tmplData); err != nil {
		return "", err
	}
	return w.String(), nil
}
