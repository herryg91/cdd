package serviceYaml

type ServiceYAML struct {
	Version      string     `yaml:"version"`
	Language     string     `yaml:"language"`
	ServiceName  string     `yaml:"name"`
	Contract     Contract   `yaml:"contract"`
	Dependencies Dependency `yaml:"dependencies"`
}

type Contract struct {
	Config     ContractConfig `yaml:"config"`
	ProtoFiles []string       `yaml:"proto-files"`
}

type ContractConfig struct {
	OutputGrst          string `yaml:"output-grst"`
	OutputScaffoldMysql string `yaml:"output-scaffold-mysql"`
	OutputDependency    string `yaml:"output-dependency"`
}

type Dependency struct {
	Services []string `yaml:"services"`
}
