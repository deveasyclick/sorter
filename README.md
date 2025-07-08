# 🧮 Go Sorter Library

A flexible, extensible sorting framework in Go, designed to allow teams to easily plug in their own sorting strategies. Useful for sorting product catalogs or any list of domain objects using different algorithms such as by price, popularity, or custom rules.

---

## ✨ Features

- ✅ Strategy Pattern implementation for clean, pluggable sort logic
- ✅ Supports sorting by price, sales/view ratio, and more
- ✅ Easy to extend with new strategies
- ✅ Type-safe with Go generics
- ✅ Fully unit-testable

---

## 🏗️ Installation

```bash
go get github.com/deveasyclick/sorter
```

---

## 🧱 Adding a New Strategy

You can easily add your own strategy by implementing the interface:

``` go
type SortByName struct{}

func (s SortByName) Sort(products []types.Product) []types.Product {
  sort.Slice(products, func(i, j int) bool {
    return products[i].Name < products[j].Name
  })
  return products
}
```

Use it like this:

```go
sorter.SetStrategy(SortByName{})
sorted := sorter.Sort(products)
```


---

## 🧪 Testing

This project uses Go’s built-in testing package for unit tests.

✅ How to Run Tests

Run all tests in the project:

```bash
go test ./...
```

Run tests in a specific package (e.g. strategies):

```bash
go test ./strategies
```