# Leaf Common 

## Functions
Context Functions
```go
func DoSkipNoticeError(ctx *context.Context)
func DontSkipNoticeError(ctx *context.Context)
func SkipNoticeError(ctx context.Context) bool
```
Converter Functions
```go
func ConvertReflectValueToString(value reflect.Value) string
func ConvertStringToUint64(value string, defaultVal ...uint64) uint64
func ConvertUint64ToString(value uint64) string
```
Index Functions
```go
func IndexString(data []string, target string) 
func IndexInt(data []int, target int) int
func IndexInt64(data []int64, target int64) int
func IndexUint64(data []uint64, target uint64) int
func IndexFloat64(data []float64, target float64) int
```
Paging Functions
```go
func CalculateTotalPages(dataTotalCount int, limit int) int
```

## Model
Pagination Model
```go
type (
    PagingParams struct { ... }
    BasePagingResponse struct { ... }
    BaseSimplePagingResponse struct { ... }
    PagingResponse struct { ... }
    SimplePagingResponse struct { ... }
)
```

## Types
```go
type (
    NullBool struct { ... }
    NullFloat64 struct { ... }
    NullInt32 struct { ... }
    NullInt64 struct { ... }
    NullString struct { ... }
    NullTime struct { ... }
    NullUInt64 struct { ... }
)
```

