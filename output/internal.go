package output

import (
	"time"

	"github.com/FreifunkBremen/yanic/runtime"
)

var quit chan struct{}

// Start workers of database
// WARNING: Do not override this function
//  you should use New()
func Start(output Output, nodes *runtime.Nodes, config *runtime.Config) {
	quit = make(chan struct{})
	go saveWorker(output, nodes, config.Nodes.SaveInterval.Duration)
}

func Close() {
	if quit != nil {
		close(quit)
	}
}

// save periodically to output
func saveWorker(output Output, nodes *runtime.Nodes, saveInterval time.Duration) {
	ticker := time.NewTicker(saveInterval)
	for {
		select {
		case <-ticker.C:
			output.Save(nodes)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
