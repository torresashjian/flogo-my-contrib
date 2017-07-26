package log

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-tendermint")

const (
	ivMessage = "message"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

type TendermintActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &TendermintActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *TendermintActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval
func (a *TendermintActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	message, _ := context.GetInput(ivMessage).(string)

	msg := message

	activityLog.Info(msg)

	return true, nil
}
