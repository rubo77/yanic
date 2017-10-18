package template

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	assert := assert.New(t)

	nodes := &runtime.Nodes{
		List: map[string]*runtime.Node{
			"a": &runtime.Node{
				Online: true,
				Statistics: &data.Statistics{
					Clients: data.Clients{
						Total:  5,
						Wifi24: 1,
						Wifi5:  2,
						Wifi:   4,
					},
				},
			},
		},
	}

	// no panic
	assert.Panics(func() {
		Register(map[string]interface{}{
			"template_path": "/dev/not-exists",
		})
	})

	// test Marshal function
	out, err := Register(map[string]interface{}{
		"template_path": "testdata/stats-json.tmpl",
		"result_path":   "/tmp/stats.json",
	})
	os.Remove("/tmp/stats.json")
	assert.NoError(err)
	assert.NotNil(out)

	out.Save(nodes)
	_, err = os.Stat("/tmp/stats.json")
	assert.NoError(err)

	// test content
	out, err = Register(map[string]interface{}{
		"template_path": "testdata/stats.tmpl",
		"result_path":   "/tmp/stats.txt",
	})
	os.Remove("/tmp/stats.txt")
	assert.NoError(err)
	assert.NotNil(out)

	out.Save(nodes)
	b, err := ioutil.ReadFile("/tmp/stats.txt")
	assert.NoError(err)
	file := string(b)
	assert.Contains(file, "0:1:5:4:1:2")

	// could not open files to write
	out, err = Register(map[string]interface{}{
		"template_path": "testdata/stats.tmpl",
		"result_path":   "/dev/stats.txt",
	})
	assert.Panics(func() {
		out.Save(nodes)
	})
}
