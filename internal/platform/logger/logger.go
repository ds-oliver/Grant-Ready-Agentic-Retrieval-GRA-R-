package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type contextKey string

const traceIDKey contextKey = "trace_id"

// WithTraceID returns a new context with the provided trace ID.
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// LogWithContext emits a JSON log line containing basic fields and metadata.
func LogWithContext(ctx context.Context, msg string, meta map[string]interface{}) {
	payload := map[string]interface{}{
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"level":     "info",
		"message":   msg,
		"trace_id":  traceIDFromContext(ctx),
	}

	for key, value := range meta {
		payload[key] = value
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintf(os.Stdout, "{\"timestamp\":%q,\"level\":\"error\",\"message\":%q}\n", payload["timestamp"], "failed to marshal log payload")
		return
	}

	fmt.Fprintln(os.Stdout, string(encoded))
}

func traceIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return "trace-mock"
	}
	if traceID, ok := ctx.Value(traceIDKey).(string); ok && traceID != "" {
		return traceID
	}
	return "trace-mock"
}
