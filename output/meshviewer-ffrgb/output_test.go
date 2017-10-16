package meshviewerFFRGB

import (
	"os"
	"testing"

	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	assert := assert.New(t)
	out, err := Register(map[string]interface{}{})
	assert.NoError(err)
	assert.Nil(out)

	out, err = Register(map[string]interface{}{
		"enable": true,
	})
	assert.Error(err)
	assert.Nil(out)

	out, err = Register(map[string]interface{}{
		"enable": true,
		"path":   "/tmp/meshviewer.json",
	})
	os.Remove("/tmp/meshviewer.json")
	assert.NoError(err)
	assert.NotNil(out)

	out.Save(&runtime.Nodes{})
	_, err = os.Stat("/tmp/meshviewer.json")
	assert.NoError(err)
}
