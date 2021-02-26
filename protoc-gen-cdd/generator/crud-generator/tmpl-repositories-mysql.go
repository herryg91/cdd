package crudgenerator

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"

	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplRepoMysql = template.Must(template.New("repositories-mysql").Funcs(template.FuncMap{
		"GetPrimaryKeyAsString":    getPrimaryKeyAsString,
		"GetPrimaryKeyAsQueryStmt": getPrimaryKeyAsQueryStmt,
	}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}

	package {{.PackageName}}
	
	import (
		{{ if .IsImportTime }} "time" {{ end }}
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
		"{{.PackagePath}}/entity"
	)

	type MysqlRepository struct {
		db        *gorm.DB
		tableName string
	}

	func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
		return &MysqlRepository{db, "{{.DBSchema.TableName}}"}
	}

	func (r *MysqlRepository) GetByPrimaryKey({{GetPrimaryKeyAsString .FieldExt "" "" "," true true}}) (entity.{{.GetName}}, error) {
		result := entity.{{.GetName}}{}
		err := r.db.Table(r.tableName).Where("{{GetPrimaryKeyAsQueryStmt .FieldExt}}", {{GetPrimaryKeyAsString .FieldExt "" "" "," true false}}).Scan(&result).Error
		return result, err
	}

	func (r *MysqlRepository) GetAll() ([]entity.{{.GetName}}, error) {
		result := []entity.{{.GetName}}{}
		err := r.db.Table(r.tableName).Find(&result).Error
		return result, err
	}

	func (r *MysqlRepository) Create(in entity.{{.GetName}}) (entity.{{.GetName}}, error) {
		{{ if not .DBSchema.DisableTimestampTracking}}
		in.CreatedAt = time.Now()
		in.UpdatedAt = time.Now() {{ end }}

		err := r.db.Table(r.tableName).Create(&in).Error
		return in, err
	}

	func (r *MysqlRepository) Update(in entity.{{.GetName}}) (entity.{{.GetName}}, error) {
		in.UpdatedAt = time.Now()
		err := r.db.Table(r.tableName).Updates(&in).Error
		return in, err
	}

	func (r *MysqlRepository) Delete({{GetPrimaryKeyAsString .FieldExt "" "" "," true true}}) error {
		return r.db.Table(r.tableName).Delete(&entity.{{.GetName}}{}, "{{GetPrimaryKeyAsQueryStmt .FieldExt}}", {{GetPrimaryKeyAsString .FieldExt "" "" "," true false}}).Error
	}
	`))
)

func applyTemplateRepoMysql(mext *descriptor.MessageDescriptorExt, pkgpath string) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	packageName := mext.DBSchema.TableName //strcase.ToKebab(strings.ToLower("crud-" + mext.DBSchema.TableName))
	var tmplData = struct {
		*descriptor.MessageDescriptorExt
		PackageName  string
		PackagePath  string
		IsImportTime bool
	}{
		mext,
		strings.Replace(packageName, "-", "_", -1),
		pkgpath,
		isMessageNeedImportTime(mext),
	}

	if err := tmplRepoMysql.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     "repositories/" + packageName + "/mysql.cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}