# logrus

[logrus](https://github.com/sirupsen/logrus) logger implementation for __fits-monitor-sdk__ [meta logger](https://github.com/fits-devops/fits-monitor-sdk/tree/master/logger).

## Usage

```go
import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/fits-devops/fits-monitor-sdk/logger"
)

func ExampleWithOutput() {
	logger.DefaultLogger = NewLogger(logger.WithOutput(os.Stdout))
	logger.Infof("testing: %s", "Infof")
}

func ExampleWithLogger() {
	l := logrus.New() // *logrus.Logger
	logger.DefaultLogger = NewLogger(WithLogger(l))
	logger.Infof("testing: %s", "Infof")
}
```

