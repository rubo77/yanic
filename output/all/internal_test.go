package all

import (
	"errors"
	"testing"

	"github.com/FreifunkBremen/yanic/output"
	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

type testOutput struct {
	output.Output
	CountSave int
}

func (c *testOutput) Save() {
	c.CountSave++
}

func TestStart(t *testing.T) {
	assert := assert.New(t)

	nodes := &runtime.Nodes{}

	globalOutput := &testOutput{}
	output.RegisterAdapter("a", func(nodes *runtime.Nodes, config interface{}) (output.Output, error) {
		return globalOutput, nil
	})
	output.RegisterAdapter("b", func(nodes *runtime.Nodes, config interface{}) (output.Output, error) {
		return globalOutput, nil
	})
	output.RegisterAdapter("c", func(nodes *runtime.Nodes, config interface{}) (output.Output, error) {
		return globalOutput, nil
	})
	output.RegisterAdapter("d", func(nodes *runtime.Nodes, config interface{}) (output.Output, error) {
		return nil, nil
	})
	output.RegisterAdapter("e", func(nodes *runtime.Nodes, config interface{}) (output.Output, error) {
		return nil, errors.New("blub")
	})
	allOutput, err := Register(nodes, map[string][]interface{}{
		"a": []interface{}{"a1", "a2"},
		"b": nil,
		"c": []interface{}{"c1"},
		"d": []interface{}{"d0"}, // fetch continue command in Connect
	})
	assert.NoError(err)

	assert.Equal(0, globalOutput.CountSave)
	allOutput.Save()
	assert.Equal(3, globalOutput.CountSave)

	_, err = Register(nodes, map[string][]interface{}{
		"e": []interface{}{"give me an error"},
	})
	assert.Error(err)
}
