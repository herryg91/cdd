package usecase_mysql

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplUseCaseRepoImpl = template.Must(template.New("usecase-repo-impl").Funcs(template.FuncMap{}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}
	package {{.GetCrudPackageName}}
	
	import (
		"fmt"
		"gorm.io/gorm"
		"{{.GoModuleName}}/entity"
		{{.Mysql.TableName}}_ds "{{.GoModuleName}}/drivers/datasource/mysql/{{.Mysql.TableName}}"
		"errors"
	)
	type repository struct {
		db     *gorm.DB
		ds *{{.Mysql.TableName}}_ds.MysqlDatasource
	}
	func NewRepository(db *gorm.DB, ds *{{.Mysql.TableName}}_ds.MysqlDatasource) Repository {
		return &repository{db, ds}
	}
	func (r *repository) GetByPrimaryKey({{ .GetPrimaryKeyAsString "" "" "," true true }}) (*entity.{{.GetName}}, error) {
		out, err := r.ds.GetByPrimaryKey({{ .GetPrimaryKeyAsString "" "" "," true false }})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrRecordNotFound
			}
			return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}
		return out.To{{.GetName}}Entity(), err
	}
	func (r *repository) GetAll() ([]*entity.{{.GetName}}, error) {
		datas, err := r.ds.GetAll()
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}
		out := []*entity.{{.GetName}}{}
		for _, data := range datas{
			out = append(out, data.To{{.GetName}}Entity())
		}
		return out, err
	}
	func (r *repository) Create(in entity.{{.GetName}}) (*entity.{{.GetName}}, error) {
		out, err := r.ds.Create(*{{.Mysql.TableName}}_ds.{{.GetName}}Model{}.From{{.GetName}}Entity(in))
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}
		return out.To{{.GetName}}Entity(), err
	}
	func (r *repository) Update(in entity.{{.GetName}}) (*entity.{{.GetName}}, error) {
		out, err := r.ds.Update(*{{.Mysql.TableName}}_ds.{{.GetName}}Model{}.From{{.GetName}}Entity(in))
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}
		return out.To{{.GetName}}Entity(), err
	}
	func (r *repository) Delete({{ .GetPrimaryKeyAsString "" "" "," true true }}) error {
		err := r.ds.Delete({{ .GetPrimaryKeyAsString "" "" "," true false }})
		if err != nil {
			return fmt.Errorf("%w: %s", ErrDatabaseError, err.Error())
		}
		return err
	}

	/*
	// Add this in drivers/datasource/mysql/{{.Mysql.TableName}}/{any_file}.go
	
	func (model *{{.GetName}}Model) To{{.GetName}}Entity() *entity.{{.GetName}} {
		return &entity.{{.GetName}}{ 
			{{- range $fext := .FieldExt}}
			{{$fext.GetName}}:model.{{$fext.GetName}},
			{{- end}}
		}
	}
	func ({{.GetName}}Model) From{{.GetName}}Entity(in entity.{{.GetName}}) *{{.GetName}}Model {
		return &{{.GetName}}Model{ 
			{{- range $fext := .FieldExt}}
			{{$fext.GetName}}:in.{{$fext.GetName}},
			{{- end}}
		}
	}
	*/
	`))
)

func applyTemplateUseCaseRepoImpl(sm ScaffoldMysql) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		ScaffoldMysql
	}{
		sm,
	}

	if err := tmplUseCaseRepoImpl.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     "usecase/" + sm.GetCrudPackageName() + "/repository-impl-mysql.cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}
