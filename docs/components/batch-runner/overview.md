# Batch Runner Component

The batch runner component adds a `batch` command that executes calculator operations from a CSV file.

Input files must use the header `operation,left,right`. Each following row produces one calculated result.

## Example

```bash
go run ./cmd/canonicalcalc batch examples/batch/operations.csv
```
