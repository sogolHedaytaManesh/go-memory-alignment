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

Each folder contains a standalone Go file demonstrating a specific memory alignment concept.  
You can run them individually using `go run`.

### 1. Stack vs Heap
```bash
  cd stack_vs_heap
  go run stack_vs_heap.go
```
### 2. Struct Ordering & Padding
```bash
  cd struct_ordering
  go run struct_ordering.go
```
### 3. Atomic Alignment
```bash
  cd atomic_alignment
  go run atomic_alignment.go
```
### 4. Cache Line & False Sharing
```bash
  cd cache_line
  go run cache_line.go
```
### 5. Pointers & Escape Analysis
```bash
  cd pointers_and_escapes
  go run pointers_and_escapes.go
```