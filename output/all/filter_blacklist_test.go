package all

import (
	"testing"

	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestFilterBlacklist(t *testing.T) {
	assert := assert.New(t)
	var config filterConfig

	config = map[string]interface{}{}

	assert.Nil(config.Blacklist())

	config["blacklist"] = []interface{}{"a", "c"}
	list := *config.Blacklist()
	assert.Len(list, 2)

	n := filterBlacklist(&runtime.Node{Nodeinfo: &data.NodeInfo{NodeID: "a"}}, list)
	assert.Nil(n)

	n = filterBlacklist(&runtime.Node{Nodeinfo: &data.NodeInfo{}}, list)
	assert.NotNil(n)

	n = filterBlacklist(&runtime.Node{}, list)
	assert.NotNil(n)

}
