# Canonical Fixtures

This repository intentionally includes identical files in different paths. The goal is to test whether RAG ingestion can detect that the content is canonical even when the file path suggests different ownership or purpose.

## Identical Operation Request

- `testdata/canonical/operation-request.json`
- `docs/components/calculator/operation-request.json`
- `internal/calculator/testdata/operation-request.json`

## Identical Shared Component Note

- `testdata/canonical/shared-component-note.md`
- `docs/components/calculator/shared-component-note.md`
- `docs/components/conversion/shared-component-note.md`
