package all

import (
	"github.com/FreifunkBremen/yanic/output"
	"github.com/FreifunkBremen/yanic/runtime"
)

type Output struct {
	output.Output
	list   map[int]output.Output
	filter map[int]filterConfig
}

func Register(configuration map[string]interface{}) (output.Output, error) {
	list := make(map[int]output.Output)
	filter := make(map[int]filterConfig)
	i := 1
	allOutputs := configuration
	for outputType, outputRegister := range output.Adapters {
		outputConfigs, ok := allOutputs[outputType].([]map[string]interface{})
		if !ok {
			continue
		}
		for _, config := range outputConfigs {
			output, err := outputRegister(config)
			if err != nil {
				return nil, err
			}
			if output == nil {
				continue
			}
			list[i] = output
			if c, ok := config["filter"]; ok {
				filter[i] = c.(map[string]interface{})
			}
			i++
		}
	}
	return &Output{list: list, filter: filter}, nil
}

func (o *Output) Save(nodes *runtime.Nodes) {
	for i, item := range o.list {
		filteredNodes := nodes
		if config, ok := o.filter[i]; ok {
			filteredNodes = config.filtering(nodes)
		}

		item.Save(filteredNodes)
	}
}
