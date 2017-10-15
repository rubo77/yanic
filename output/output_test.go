package output

import (
	"testing"

	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	assert := assert.New(t)
	assert.Len(Adapters, 0)

	RegisterAdapter("blub", func(nodes *runtime.Nodes, config interface{}) (Output, error) {
		return nil, nil
	})

	assert.Len(Adapters, 1)
}
