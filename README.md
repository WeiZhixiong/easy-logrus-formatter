## Easy Logrus Formatter
Provide a user-friendly formatter for [Logrus](https://github.com/sirupsen/logrus)
Some inspiration taken from [logrus-easy-formatter](https://github.com/t-tomalak/logrus-easy-formatter)

## Default output
When format options are not provided `Formatter` will output
```bash
2023-01-01T01:01:01+00:00||INFO||Log message
```

## Getting started

### Getting easy-logrus-formatter
```bash
go get -u github.com/WeiZhixiong/easy-logrus-formatter
```

### Sample Usage
```go
package main

import (
	formatter "github.com/WeiZhixiong/easy-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&formatter.Formatter{})
	log.Info("Log message")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
```
Above sample will produce:
```bash
2023-01-01T01:01:01+00:00||INFO||Log message
2023-01-01T01:01:01+00:00||INFO||A group of walrus emerges from the ocean||animal=walrus||size=10
```