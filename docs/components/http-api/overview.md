# HTTP API Component

The HTTP API component adds a `serve` command that exposes calculator operations over HTTP.

The `/calculate` endpoint accepts `op`, `left`, and `right` query parameters. It returns a JSON response with the selected operation and result.

## Example

```bash
go run ./cmd/canonicalcalc serve :8080
curl "http://localhost:8080/calculate?op=add&left=7&right=5"
```
