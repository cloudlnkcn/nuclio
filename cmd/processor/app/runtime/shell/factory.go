package shell

import (
	"github.com/nuclio/nuclio/cmd/processor/app/runtime"
	"github.com/nuclio/nuclio/pkg/logger"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type factory struct{}

func (f *factory) Create(parentLogger logger.Logger,
	configuration *viper.Viper) (runtime.Runtime, error) {

	newConfiguration, err := runtime.NewConfiguration(configuration)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create configuration")
	}

	return NewRuntime(parentLogger.GetChild("shell").(logger.Logger),
		&Configuration{
			Configuration: *newConfiguration,
			ScriptPath:    configuration.GetString("path"),
			ScriptArgs:    configuration.GetStringSlice("args"),
		})
}

// register factory
func init() {
	runtime.RegistrySingleton.Register("shell", &factory{})
}