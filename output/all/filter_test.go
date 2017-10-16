package all

import (
	"testing"

	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	nodes := &runtime.Nodes{}
	config := filterConfig{
		"no_owner": false,
	}
	nodes = config.filtering(nodes)
	assert.Len(nodes.List, 0)
}
