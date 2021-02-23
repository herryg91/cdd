package crudgenerator

import (
	"github.com/herryg91/cdd/protoc-gen-cdd/descriptor"
	"github.com/herryg91/cdd/protoc-gen-cdd/generator"
)

type CRUDGeneratorTemplate struct {
	name       string
	pkgPath    string
	descriptor *descriptor.Descriptor
}

func New(d *descriptor.Descriptor, pkgPath string) *CRUDGeneratorTemplate {
	result := &CRUDGeneratorTemplate{
		name:       "crud",
		descriptor: d,
		pkgPath:    pkgPath,
	}

	return result
}

func (t *CRUDGeneratorTemplate) Generate() ([]*generator.GeneratorResponseFile, error) {
	var files []*generator.GeneratorResponseFile
	for _, f := range t.descriptor.FileToGenerate {
		mextForEntity := []*descriptor.MessageDescriptorExt{}
		for _, mext := range f.MessageExt {
			if mext.DBSchema.IsDbSchema && mext.DBSchema.TableName != "" {
				mextForEntity = append(mextForEntity, mext)

				fileUseCaseErrors, err := applyTemplateUseCaseErrors(mext)
				if err != nil {
					return nil, err
				}
				files = append(files, fileUseCaseErrors)

				fileUseCaseIntf, err := applyTemplateUseCaseIntf(mext, t.pkgPath)
				if err != nil {
					return nil, err
				}
				files = append(files, fileUseCaseIntf)

				fileUseCaseRepoImpl, err := applyTemplateUseCaseRepoImpl(mext, t.pkgPath)
				if err != nil {
					return nil, err
				}
				files = append(files, fileUseCaseRepoImpl)

				fileUseCaseImpl, err := applyTemplateUseCaseImpl(mext, t.pkgPath)
				if err != nil {
					return nil, err
				}
				files = append(files, fileUseCaseImpl)

			}
		}
		fileEntity, err := applyTemplateEntity(mextForEntity, f.GetPackage(), f.GetName())
		if err != nil {
			return nil, err
		}
		files = append(files, fileEntity)
	}

	return files, nil
}
