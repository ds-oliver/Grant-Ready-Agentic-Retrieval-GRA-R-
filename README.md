# Grant-Ready Agentic Retrieval (GRA-R)

## Architecture

GRA-R uses a hybrid architecture: Go services handle ingestion, platform concerns, and API boundaries, while Python agents (LangGraph/LangChain) orchestrate the agentic workflows. This split keeps ingestion and platform layers fast, strongly typed, and production ready, while allowing the agent runtime to iterate quickly on reasoning and prompt logic.

## Distributed Resilience

The platform implements defensive patterns to avoid duplicate work and concurrency issues:

- **Idempotency middleware** (`internal/platform/http/idempotency.go`) checks incoming `X-Request-ID` values and returns cached responses for repeated requests, preventing double-spending or repeated side effects.
- **Distributed locks** (`internal/platform/lock/lock.go`) provide an acquire/release interface (mocked in memory today) to prevent race conditions when multiple workers attempt to process the same resource.

## Security

The PII sanitization layer (`internal/platform/security/sanitize.go`) redacts common sensitive data such as email addresses and SSN-like patterns before they are logged or stored.

## Run Instructions

### Go

```bash
go test ./...
```

### Python agent

```bash
python agents/agentic_core.py
```
