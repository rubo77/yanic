package nodelist

import (
	"errors"

	"github.com/FreifunkBremen/yanic/output"
	"github.com/FreifunkBremen/yanic/runtime"
)

type Output struct {
	output.Output
	path  string
	nodes *runtime.Nodes
}

type Config map[string]interface{}

func (c Config) Enable() bool {
	if enable, ok := c["enable"]; ok {
		return enable.(bool)
	}
	return false
}

func (c Config) Path() string {
	if path, ok := c["path"]; ok {
		return path.(string)
	}
	return ""
}

func init() {
	output.RegisterAdapter("nodelist", Register)
}

func Register(nodes *runtime.Nodes, configuration interface{}) (output.Output, error) {
	var config Config
	config = configuration.(map[string]interface{})
	if !config.Enable() {
		return nil, nil
	}

	if path := config.Path(); path != "" {
		return &Output{
			path:  path,
			nodes: nodes,
		}, nil
	}
	return nil, errors.New("no path given")

}

func (o *Output) Save() {
	o.nodes.RLock()
	defer o.nodes.RUnlock()

	runtime.SaveJSON(transform(o.nodes), o.path)
}
