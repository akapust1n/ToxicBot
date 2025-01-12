package bulling

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	FilePath       string        `envconfig:"BULLINGS_FILE" required:"true"`
	ThresholdCount int           `envconfig:"BULLINGS_THRESHOLD_COUNT" default:"5"`
	ThresholdTime  time.Duration `envconfig:"BULLINGS_THRESHOLD_TIME" default:"1m"`
	Cooldown       time.Duration `envconfig:"BULLINGS_COOLDOWN" default:"1h"`
}

func (b *Bulling) parseConfig() error {
	if err := envconfig.Process("", &b.cfg); err != nil {
		envconfig.Usage("", b.cfg)
		return err
	}

	return nil
}
