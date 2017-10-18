package all

import (
	"testing"

	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
	"github.com/stretchr/testify/assert"
)

func TestFilterInArea(t *testing.T) {
	assert := assert.New(t)
	var config filterConfig
	areaConfig := map[string]interface{}{
		"latitude_min":  3.0,
		"latitude_max":  5.0,
		"longitude_min": 10.0,
		"longitude_max": 12.0,
	}
	config = map[string]interface{}{}

	area := config.InArea()
	assert.Nil(area)

	config["in_area"] = areaConfig
	area = config.InArea()
	assert.NotNil(area)

	// drop area without nodeinfo
	n := filterLocationInArea(&runtime.Node{}, area)
	assert.Nil(n)

	// keep without location
	n = filterLocationInArea(&runtime.Node{Nodeinfo: &data.NodeInfo{}}, area)
	assert.NotNil(n)

	// zeros not in area (0, 0)
	n = filterLocationInArea(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{},
	}}, area)
	assert.Nil(n)

	// in area
	n = filterLocationInArea(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{Latitude: 4.0, Longtitude: 11.0},
	}}, area)
	assert.NotNil(n)

	n = filterLocationInArea(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{Latitude: 4.0, Longtitude: 13.0},
	}}, area)
	assert.Nil(n)

	n = filterLocationInArea(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{Latitude: 6.0, Longtitude: 11.0},
	}}, area)
	assert.Nil(n)

	n = filterLocationInArea(&runtime.Node{Nodeinfo: &data.NodeInfo{
		Location: &data.Location{Latitude: 1.0, Longtitude: 2.0},
	}}, area)
	assert.Nil(n)
}
