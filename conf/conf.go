package conf

import (
	"io/ioutil"

	"github.com/Depado/conftags"
	yaml "gopkg.in/yaml.v2"
)

// Conf stores all the configuration related to this instance
type Conf struct {
	Server        Server `yaml:"server"`
	OpenSireneAPI string `yaml:"opensirene_api"`
}

// C is the main exported configuration
var C Conf

// Parse will parse every nested fields with the env/defaults parser
// and set the values accordingly
func (c *Conf) Parse() error {
	return conftags.Parse(&c.Server)
}

// Load loads the configuration file into C
func Load(fp string) error {
	var err error
	var c []byte

	if c, err = ioutil.ReadFile(fp); err != nil {
		return err
	}
	if err = yaml.Unmarshal(c, &C); err != nil {
		return err
	}
	return C.Parse()
}
