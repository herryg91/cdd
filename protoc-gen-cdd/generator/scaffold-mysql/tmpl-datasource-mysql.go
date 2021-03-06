package scaffold_mysql

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	tmplRepoMysql = template.Must(template.New("repositories-mysql").Funcs(template.FuncMap{}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.FileExt.GetName}}

	package {{.Mysql.TableName}}
	
	import (
		{{ if .NeedImportTime }} "time" {{ end }}
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
		"{{.GoModuleName}}/entity"
	)

	type MysqlDatasource struct {
		db        *gorm.DB
		tableName string
	}

	func NewMysqlDatasource(db *gorm.DB) *MysqlDatasource {
		return &MysqlDatasource{db, "{{.Mysql.TableName}}"}
	}

	func (r *MysqlDatasource) GetByPrimaryKey({{ .GetPrimaryKeyAsString "" "" "," true true }}) (*entity.{{.GetName}}, error) {
		result := &entity.{{.GetName}}{}
		err := r.db.Table(r.tableName).Where("{{ .GetPrimaryKeyAsQueryStmt }}", {{ .GetPrimaryKeyAsString "" "" "," true false }}).Scan(&result).Error
		return result, err
	}

	func (r *MysqlDatasource) GetAll() ([]*entity.{{.GetName}}, error) {
		result := []*entity.{{.GetName}}{}
		err := r.db.Table(r.tableName).Find(&result).Error
		return result, err
	}

	func (r *MysqlDatasource) Create(in entity.{{.GetName}}) (*entity.{{.GetName}}, error) {
		{{ if not .Mysql.DisableTimestampTracking}}
		in.CreatedAt = time.Now()
		in.UpdatedAt = time.Now() {{ end }}

		err := r.db.Table(r.tableName).Create(&in).Error
		if err != nil {
			return nil, err
		}
		return &in, nil
	}

	func (r *MysqlDatasource) Update(in entity.{{.GetName}}) (*entity.{{.GetName}}, error) {
		in.UpdatedAt = time.Now()
		err := r.db.Table(r.tableName).Where("{{ .GetPrimaryKeyAsQueryStmt }}", {{ .GetPrimaryKeyAsString "in." "" "," false false }}).Updates(&in).Error
		if err != nil {
			return nil, err
		}
		return &in, nil
	}

	func (r *MysqlDatasource) Delete({{ .GetPrimaryKeyAsString "" "" "," true true }}) error {
		return r.db.Table(r.tableName).Delete(&entity.{{.GetName}}{}, "{{ .GetPrimaryKeyAsQueryStmt }}", {{ .GetPrimaryKeyAsString "" "" "," true false }}).Error
	}
	`))
)

func applyTemplateRepoMysql(sm ScaffoldMysql) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)
	var tmplData = struct {
		ScaffoldMysql
	}{
		sm,
	}

	if err := tmplRepoMysql.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     "drivers/datasource/mysql/" + sm.Mysql.TableName + "/" + sm.Mysql.TableName + ".cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}
