# Textio - Go Messaging Service

Textio is a high-performance backend messaging service . This project demonstrates my transition from dynamic languages to the statically-typed, concurrent world of Go.

## Project Overview

Textio serves as a robust engine for managing SMS and internal messaging logic. The project focuses on idiomatic Go patterns, performance-oriented memory management, and safe concurrency.

## Key Technical Concepts Applied

### 1. Concurrency & Synchronization
- **Channels**: Used to orchestrate data flow between goroutines, adhering to the proverb: *"Don't communicate by sharing memory, share memory by communicating."*
- **Mutexes**: Implemented `sync.Mutex` and `sync.RWMutex` to serialize access to shared state and prevent race conditions.

### 2. Type-Safe Enums (iota)
- **Status Management**: Instead of using raw strings or integers, I used custom types and the `iota` keyword to create type-safe enumerations.
- **Implementation**: You can find my implementation of message and user states in [`iota.go`](./iota.go).

### 3. Generics (Type Parameters)
- Built reusable functions and data structures using **Generics** to eliminate code duplication while maintaining strict compile-time type safety.
- Applied **Type Constraints** to ensure generic logic only runs on valid data types.

### 4. Interfaces & Decoupling
- Utilized Go's implicit interfaces to create flexible, decoupled components.
- Followed the principle that "Clear is better than clever" by keeping interface definitions small and purposeful.

### 5. Error Handling
- Treated **errors as values**. By explicitly checking and handling errors, I ensured the service remains stable and avoids unexpected panics.

### 6. Memory Efficiency
- Managed state using **Pointers** and **Pointer Receivers** to avoid expensive copies of large data structures, ensuring the backend remains performant under load.

## Philosophy

This project was built following the **Go Proverbs**. It prioritizes:
- **Simplicity**: Avoiding "clever" code in favor of readable logic.
- **Robustness**: Handling every error state gracefully.
- **Performance**: Leveraging Go's compiled nature and efficient concurrency primitives.

## Getting Started

### Prerequisites
- Go 1.20+

### Installation & Usage
1. Clone the repository and build the binary:
   ```bash
   go build -o textio
