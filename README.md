# Mastering Memory Alignment in Go

This repository contains practical Go examples illustrating memory alignment concepts, including stack vs heap, struct ordering, atomic alignment, cache lines, and pointer escape analysis. All examples are standalone and designed to complement the article [Mastering Memory Alignment in Go: From Stack to Heap](https://medium.com/@sogol.hedayatmanesh/mastering-memory-alignment-in-go-from-stack-to-heap-4b4e10c188c4).

---

## Project Structure
```
go-memory-alignment
├── concurrency_alignment
│ ├── atomic_example.go 
│ ├── false_sharing.go 
│ └── unsafe_example.go 
├── stack_vs_heap
│ ├── escape_analysis.go 
│ ├── heap_example.go 
│ └── stack_example.go 
├── struct_alignment
│ ├── bad_order.go 
│ ├── good_order.go 
│ └── nested_struct.go 
├── tools
│ ├── inspect_size.go 
│ └── reflect_example.go 
└── go.mod 
```

- **stack_vs_heap**: Demonstrates stack vs heap allocation, lifetime, and escape analysis.
- **struct_ordering**: Shows how field ordering affects struct size and padding.
- **atomic_alignment**: Explains proper alignment for atomic operations.
- **cache_line**: Highlights cache line padding to avoid false sharing.
- **pointers_and_escapes**: Demonstrates pointer usage and how Go decides whether variables escape to the heap.

---

## How to Run the Examples

## Clone the repo
```bash
  git clone https://github.com/sogolHedaytaManesh/go-memory-alignment.git
```

Each folder contains a standalone Go file demonstrating a specific memory alignment concept.  
You can run them individually using `go run`.

### 1. Concurrency Alignment
```bash
  cd concurrency_alignment
  go run atomic_example.go
  go run false_sharing.go
  go run unsafe_example.go
```
### 2. Stack VS Heap
```bash
  cd stack_vs_heap
  go run escape_analysis.go
  go run heap_example.go
  go run stack_example.go
```
### 3. Struct Alignment
```bash
  cd struct_alignment
  go run bad_order.go
  go run good_order.go
  go run nested_struct.go
```
### 4. Tools
```bash
  cd tools
  go run inspect_size.go
  go run reflect_example.go
```