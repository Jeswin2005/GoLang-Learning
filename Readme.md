# Go Learning Progress

## 16-10-2025
### Topics Covered
- Data types  
- Operators  
- If–Else  
- Switch  
- For loop  
- Strings  
- Functions  
- Arrays and slices  
- Pointers  
- Map  
- Interface  

---

## 17-10-2025
### Topics Covered
- File operation  
- JSON handling  
- XML handling  
- Call by value and call by reference  
- How interface works  
  - Implicit Implementation:  
    In Go, any struct that has the required methods of an interface implicitly implements that interface without explicit declaration.  
- Instead of class and object  
  - In Go, structs are used instead of classes.  
  - Method functions act as class methods.  
  - Struct instances act as objects.  
- Basic API Implementation  
  - Implemented GET, POST, PUT, and DELETE methods.

---

## 21-10-2025
## Go Internal Execution and Build Toolchain
### Go Compiler (cmd/compile)
The Go compiler transforms source code into machine code through multiple phases:
- Parsing: Converts Go source code into a syntax tree by identifying tokens, keywords, and identifiers.
- Typing: Performs syntax and type checking for variables, functions, and return types.
- Noding (IR Construction): Converts code into Go’s internal Intermediate Representation (IR).
  - Inlining: Small function calls are replaced directly with their code to reduce call overhead.
  - Escape Analysis: Determines whether a variable should live on the fast stack or escape to the slower heap.
  - Devirtualization: Replaces interface calls with direct calls when possible.
- Walk (Code Simplification):
  - Desugaring: Rewrites high-level constructs (like switch) into simpler jumps.
  - Converts maps and channels into runtime calls.
- SSA (Static Single Assignment): Converts IR into SSA form where variables are assigned only once — improving optimization and removing redundant code.
- Machine Code Generation: SSA is converted to CPU-specific assembly and optimized for efficient register usage.
- Export Data Generation: Produces export data (type info and public function signatures) for other packages to avoid recompilation.

### Assembler (cmd/asm)
- Takes assembly files (.s) and converts them into object files (.o) containing machine code and metadata.
- Ensures consistent memory and subroutine management across CPU architectures.

### Linker (cmd/link)
- Combines all object files (.o) from the compiler and assembler with the Go runtime to create a single executable.
- Produces .exe on Windows or a binary (ELF) on Linux/macOS.

### Go Runtime
Go has its own runtime, unlike C or C++, which depend heavily on the OS. It handles:
- Memory allocation
- Garbage collection
- Goroutine scheduling (GMP model)
G - Goroutine - Lightweight thread of execution
M - Machine - Actual OS thread
P - Processor - Logical Processor managing execution queue

### Program vs Process vs Thread
- Program: It is a passive, static entity that does nothing on its own. For example, the Microsoft Word installation file (.exe) is a program.
- Process: When you launch that program by double-clicking the icon, the operating system loads the program's code into RAM (main memory) and creates an active instance of it. This active instance is a process, and it includes the program code, its own allocated memory space (heap and stack), registers, and other resources required for execution.
- Thread: Within a process, a thread is a single, smaller unit of execution. A single process can have one or many threads. All threads within the same process share the process's resources, such as its memory space. This makes threads "lightweight" compared to processes, with less overhead for creation and switching. 

### Garbage Collector
- Concurrent: Runs alongside user code.

- Mark-and-sweep: Marks live objects, then sweeps away unreferenced ones.

- Pacer: Dynamically adjusts GC frequency based on allocation rate.

### go run vs go build
- go run main.go
  - Compiles, runs and deletes temporary binary used for quick testing.
- go build main.go
  - compiles and saves binary used for repeated executiion without recompiling.

### executing pre compiled code
- In windows by double clicking the .exe file or typing myapp.exe in cmd line runs the exe file
- In linux,
  - chmod +x myapp (make file executable)
  - ./myapp (executes the file) 

### Build commands
- go build -o bin/myapp main.go (Specify output name/location)
- go build -a (Force rebuild all packages)
- go build -p 8 (Parallel compilation (default = CPU count))
- go build -trimpath (Remove file paths from binary)
- go build -gcflags='-m' (Show optimization decisions)
- go build -gcflags='-m=2' (Verbose escape analysis)
- go build -gcflags='-l' (Disable inlining)
- go build -gcflags='-N' (Disable optimizations)
- go build -gcflags='-S' (Print assembly listing)
- go build -gcflags='-e' (More error checking)

### Cross Compilation
- GOOS=linux GOARCH=amd64 go build -o myapp-linux (build for linux from any os)
- GOOS=windows GOARCH=amd64 go build -o myapp.exe (build for Windows from any os)
- GOOS=darwin GOARCH=arm64 go build -o myapp-macos-arm (build for macOS from any os)