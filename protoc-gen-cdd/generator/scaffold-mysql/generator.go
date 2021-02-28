package scaffold_mysql

import (
	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
)

type ScaffoldMysqlGeneratorTemplate struct {
	name         string
	goModuleName string
	descriptor   *descriptor.Descriptor
}

func New(d *descriptor.Descriptor, goModuleName string) *ScaffoldMysqlGeneratorTemplate {
	result := &ScaffoldMysqlGeneratorTemplate{
		name:         "scaffold-mysql",
		descriptor:   d,
		goModuleName: goModuleName,
	}

	return result
}

func (t *ScaffoldMysqlGeneratorTemplate) Generate() ([]*generator.GeneratorResponseFile, error) {
	var files []*generator.GeneratorResponseFile
	for _, f := range t.descriptor.FileToGenerate {
		entities := []ScaffoldMysql{}
		for _, mext := range f.MessageExt {
			if !mext.Mysql.Scaffold {
				continue
			}
			entities = append(entities, ScaffoldMysql{mext, t.goModuleName})

			fileUseCaseErrors, err := applyTemplateUseCaseErrors(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseErrors)

			fileUseCaseIntf, err := applyTemplateUseCaseIntf(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseIntf)

			fileRepoMysql, err := applyTemplateRepoMysql(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileRepoMysql)

			fileUseCaseRepoImpl, err := applyTemplateUseCaseRepoImpl(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseRepoImpl)

			fileUseCaseImpl, err := applyTemplateUseCaseImpl(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseImpl)
		}
		if len(entities) > 0 {
			fileEntity, err := applyTemplateEntity(entities, f)
			if err != nil {
				return nil, err
			}
			files = append(files, fileEntity)
		}
	}

	return files, nil
}
