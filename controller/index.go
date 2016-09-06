// Package controller can send commands to the model to update the model's
// state. It can also send commands to its associated view to change the view's
// presentation of the model (e.g. by scrolling through a document).
package controller

import (
	"fmt"
	"net/http"

	"github.com/gouvinb/go-microservice/shared"
)

// Index displays the default home page
func Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := shared.Instance(r)

	msg := "{ \"message\": \"if you see this json, it's because the micro service is OP\"}"
	fmt.Fprint(w, msg)

	sess.Save(r, w)
}
