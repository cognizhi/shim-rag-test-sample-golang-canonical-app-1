# shim-rag-test-sample-golang-canonical-app-1

## Project Overview

This repository is a small Go CLI used to test canonical document handling in RAG ingestion. This branch removes the length conversion feature and keeps the arithmetic and named constants commands.

The repository intentionally keeps some byte-for-byte identical files in different paths. Those files represent the same canonical content appearing under different app, component, and test-data locations.

## Local Setup

### Build

```bash
make build
```

### Run Tests

```bash
make test
```

### Run Commands

```bash
go run ./cmd/canonicalcalc add 7 5
go run ./cmd/canonicalcalc div 10 2
go run ./cmd/canonicalcalc constant pi
go run ./cmd/canonicalcalc convert-length 1200 meter kilometer
```

## Commands On Main

| Command | Purpose |
|---|---|
| `add <left> <right>` | Add two numbers |
| `sub <left> <right>` | Subtract the second number from the first |
| `mul <left> <right>` | Multiply two numbers |
| `div <left> <right>` | Divide the first number by the second |
| `constant <name>` | Print a named constant such as `pi`, `e`, or `tau` |

## Canonical Fixture Paths

These paths contain intentionally identical content so ingestion can test canonical grouping:

| Canonical content | Duplicate paths |
|---|---|
| Operation request JSON | `testdata/canonical/operation-request.json`, `docs/components/calculator/operation-request.json`, `internal/calculator/testdata/operation-request.json` |
| Shared component note | `testdata/canonical/shared-component-note.md`, `docs/components/calculator/shared-component-note.md` |

## Branch Plan

The branch set is designed for RAG tests that compare added files and removed files across repository history:

| Branch | Intent |
|---|---|
| `main` | Stable calculator app baseline |
| `feature/http-api` | Adds an HTTP API command and server package |
| `feature/batch-runner` | Adds batch execution from CSV files |
| `remove/conversion-feature` | Removes the length conversion feature files |
| `remove/constants-feature` | Removes the named constants feature files |
