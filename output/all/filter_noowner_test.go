package all

import (
	"testing"

	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestFilterNoOwner(t *testing.T) {
	assert := assert.New(t)
	var config filterConfig

	config = map[string]interface{}{}

	assert.True(config.NoOwner())

	config["no_owner"] = true
	assert.True(config.NoOwner())

	config["no_owner"] = false
	assert.False(config.NoOwner())

	n := filterNoOwner(&runtime.Node{})
	assert.NotNil(n)

	n = filterNoOwner(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Owner: &data.Owner{
			Contact: "blub",
		},
	}})
	assert.NotNil(n)
	assert.Nil(n.Nodeinfo.Owner)
}
