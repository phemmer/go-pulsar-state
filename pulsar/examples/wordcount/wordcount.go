package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/apache/pulsar/pulsar-function-go/logutil"

	"github.com/phemmer/go-pulsar-state/pulsar"
)

func main() {
	pulsar.Start(WordCount)
}

func WordCount(ctx context.Context, input []byte) ([]byte, error) {
	fs, _ := pulsar.FromContext(ctx)
	if fs == nil {
		return nil, fmt.Errorf("missing function context")
	}

	words := strings.Split(string(input), " ")
	for _, word := range words {
		count, err := fs.IncrCounter(ctx, word, 1)
		if err != nil {
			return nil, fmt.Errorf("error incrementing '%s': %s", word, err)
		}
		logutil.Infof("The value for %s is %d", word, count)
	}

	return append(input, '!'), nil
}
