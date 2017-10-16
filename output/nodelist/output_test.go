package nodelist

import (
	"os"
	"testing"

	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	assert := assert.New(t)
	out, err := Register(&runtime.Nodes{}, map[string]interface{}{})
	assert.NoError(err)
	assert.Nil(out)

	out, err = Register(&runtime.Nodes{}, map[string]interface{}{
		"enable": true,
	})
	assert.Error(err)
	assert.Nil(out)

	out, err = Register(&runtime.Nodes{}, map[string]interface{}{
		"enable": true,
		"path":   "/tmp/nodelist.json",
	})
	os.Remove("/tmp/nodelist.json")
	assert.NoError(err)
	assert.NotNil(out)

	out.Save()
	_, err = os.Stat("/tmp/nodelist.json")
	assert.NoError(err)
}
