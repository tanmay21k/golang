## What is `iota`?

`iota` is a predeclared identifier in Go that is used inside `const` declarations to create auto-incrementing constant values.

## Good resource : [Constant & Iota](https://webreference.com/go/basics/constants/)

## Important Rules

### 1. `iota` can only be used inside `const` declarations

Valid:

```go
const (
    A = iota
)
```

Invalid:

```go
var x = iota // Compilation error
```

---

### 2. `iota` increments for each constant specification

iota starts with 0

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
    D        // 3
)
```

---

### 3. The blank identifier (`_`) still increments `iota`

The value is discarded, but `iota` continues to increase.

```go
const (
    Sunday = iota + 1 // 1
    _                 // 2 (discarded)
    Monday            // 3
)
```

Output:

```text
Sunday = 1
Monday = 3
```

---

### 4. `iota` is scoped to a single `const` block

Each new `const` declaration resets `iota` back to `0`.

```go
const A = iota // 0

const B = iota // 0
```

Output:

```text
A = 0
B = 0
```

---

## Example: [iota.go](./iota.go)
