package rollrus

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/heroku/rollbar"
)

var levelToRollbar = map[log.Level]string{
	log.ErrorLevel: rollbar.ERR,
	log.FatalLevel: rollbar.CRIT,
	log.PanicLevel: rollbar.CRIT,
}

type Hook struct {
	Client rollbar.Client
}

func (r *Hook) Fire(entry *log.Entry) error {
	if level, exists := levelToRollbar[entry.Level]; !exists {
		r.Client.Error(rollbar.ERR, fmt.Errorf(entry.Message))
	} else {
		r.Client.Error(level, fmt.Errorf(entry.Message))
	}

	return nil
}

func (r *Hook) Levels() []log.Level {
	return []log.Level{
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	}
}