package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtract(t *testing.T) {
	text := `package pkg

	import (
		"context"
		"log"
		"time"
	
		"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	)
	
	func Hello(ctx context.Context) {
		span, ctx := tracer.StartSpanFromContext(ctx, "hello")
		defer span.Finish()
		time.Sleep(time.Second * 5)
		log.Print("hello")
	}
`

	names, lines := extract(text)
	require.Equal(t, []string{"hello"}, names)
	require.Equal(t, []int{12}, lines)
}
