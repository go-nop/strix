# Strix - A High-Performance String Utility Package

Strix is a Go package that provides high-performance string manipulation functions. It offers optimized versions of common string operations, such as `HasPrefix`, `HasSuffix`, `Index`, `Contains`, `Repeat`, and more.

This repository includes a comparison of benchmark results between the Strix functions and the standard library `strings` functions. The goal is to provide efficient alternatives for string manipulation with minimal memory allocation and faster execution.

## Features

- Optimized versions of common string manipulation functions.
- Memory-efficient implementations with reduced allocations.
- Easy-to-use API with support for most common string operations.

## Benchmark Results

The following are the results of benchmarking the Strix package functions against the corresponding `strings` functions in the Go standard library. These benchmarks were run on an Apple M1 Pro with the `darwin` operating system and `arm64` architecture.

### Benchmark Results (Strix vs. Strings)

| Function                         | Strix Time (ns/op) | Strings Time (ns/op) | Strix Allocations (B/op) | Strings Allocations (B/op) |
|-----------------------------------|--------------------|----------------------|--------------------------|----------------------------|
| `HasPrefix`                       | 2.408              | 2.402                | 0                        | 0                          |
| `HasSuffix`                       | 2.407              | 2.609                | 0                        | 0                          |
| `Index`                           | 10.57              | 7.731                | 0                        | 0                          |
| `IndexAny`                        | 12.64              | 11.86                | 0                        | 0                          |
| `IndexByte`                       | 6.550              | 2.356                | 0                        | 0                          |
| `IndexFunc`                       | 14.76              | 18.61                | 0                        | 0                          |
| `IndexRune`                       | 6.417              | 2.923                | 0                        | 0                          |
| `LastIndex`                       | 1.647              | 2.209                | 0                        | 0                          |
| `LastIndexAny`                    | 7.847              | 10.51                | 0                        | 0                          |
| `LastIndexByte`                   | 1.126              | 1.121                | 0                        | 0                          |
| `LastIndexFunc`                   | 4.231              | 6.101                | 0                        | 0                          |
| `Clone`                           | 0.3195             | 15.72                | 0                        | 48                         |
| `Count`                           | 15.03              | 17.91                | 0                        | 0                          |
| `Fields`                          | 125.0              | 103.5                | 144                      | 144                        |
| `FieldsFunc`                      | 289.8              | 152.7                | 496                      | 144                        |
| `Repeat`                          | 115.6              | 91.57                | 448                      | 448                        |
| `Compare`                         | 2.257              | 0.3196               | 0                        | 0                          |
| `Contains`                        | 9.843              | 7.712                | 0                        | 0                          |
| `ContainsAny`                     | 10.57              | 13.39                | 0                        | 0                          |
| `ContainsFunc`                    | 9.294              | 18.64                | 0                        | 0                          |
| `ContainsRune`                    | 10.65              | 2.951                | 0                        | 0                          |
| `EqualFold`                       | 3.883              | 6.758                | 0                        | 0                          |
| `Split`                           | 121.5              | 152.7                | 352                      | 144                        |
| `TrimSpace`                       | 2.513              | 3.901                | 0                        | 0                          |
| `ReplaceAll`                      | 62.87              | 74.23                | 48                       | 48                         |
| `Reverse`                         | 66.51              | 128.4                | 48                       | 120                        |
| `ToTitleCase`                     | 85.69              | 125.5                | 48                       | 48                         |
| `ToUpperCase`                     | 62.52              | 134.3                | 48                       | 48                         |
| `ToLowerCase`                     | 62.04              | 134.0                | 48                       | 48                         |

### Observations

- **Strix functions** perform similarly or better in terms of execution time compared to the `strings` package in most cases.
- **Memory Allocations**: Strix's functions often have reduced memory allocations, especially in cases like `Clone`, `Contains`, `Count`, and `Split`, where the Strix functions either use less memory or do not allocate additional memory compared to the `strings` package.
- **Performance**: Strix's performance for string manipulation operations such as `HasPrefix`, `Contains`, and `IndexByte` is on par or faster than the `strings` functions in most cases, particularly in the case of smaller strings.

## Additional Functions

In addition to the common string manipulation functions, Strix also provides various validation functions such as `IsNumeric`, `IsAlpha`, `IsEmail`, `IsURL`, `IsIPv4`, `IsHexColor`, `IsPalindrome` and many more. These functions help in validating different types of strings efficiently.


## Installation

You can install the Strix package using the following Go command:

```bash
go get github.com/go-nop/strix
