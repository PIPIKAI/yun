// package enter
package tracker

import (
	"github.com/pipikai/yun/core/tracker/svc"
)

// Run
func Run() {
	tracker := svc.NewSvc()
	tracker.Server()
}
