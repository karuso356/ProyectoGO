package configuracion

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed configuracion.yaml
var config []byte

type DBconfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Nombre   string `yaml:"nombre"`
}

type Configuracion struct {
	Port      int      `yaml:"port"`
	BaseDatos DBconfig `yaml:"baseDatos"`
}

func Inicio() (*Configuracion, error) {
	var c Configuracion

	err := yaml.Unmarshal(config, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
