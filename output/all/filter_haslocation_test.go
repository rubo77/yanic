package all

import (
	"testing"

	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestFilterHasLocation(t *testing.T) {
	assert := assert.New(t)
	var config filterConfig

	config = map[string]interface{}{}

	assert.Nil(config.HasLocation())

	config["has_location"] = true
	assert.True(*config.HasLocation())

	config["has_location"] = false
	assert.False(*config.HasLocation())

	n := filterHasLocation(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{},
	}}, true)
	assert.NotNil(n)

	n = filterHasLocation(&runtime.Node{Nodeinfo: &data.NodeInfo{}}, true)
	assert.Nil(n)

	n = filterHasLocation(&runtime.Node{}, true)
	assert.Nil(n)

	n = filterHasLocation(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{},
	}}, false)
	assert.Nil(n)

	n = filterHasLocation(&runtime.Node{Nodeinfo: &data.NodeInfo{}}, false)
	assert.NotNil(n)

	n = filterHasLocation(&runtime.Node{}, false)
	assert.NotNil(n)
}
