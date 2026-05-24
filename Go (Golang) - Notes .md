# ***Go (Golang) — Interview Revision Guide*** {#go-(golang)-—-interview-revision-guide}

*Based on: Learn GO Fast: Full Tutorial | Covers all fundamentals \+ REST API patterns*

---

## ***Table of Contents*** {#table-of-contents}

[Go (Golang) — Interview Revision Guide](#go-\(golang\)-—-interview-revision-guide)

[Table of Contents](#table-of-contents)

[1\. Why Go?](#1.-why-go?)

[2\. Program Structure](#2.-program-structure)

[3\. Variables & Types](#3.-variables-&-types)

[Declaration styles](#declaration-styles)

[Basic Types](#basic-types)

[Type Conversion (explicit only)](#type-conversion-\(explicit-only\))

[Constants](#constants)

[4\. Control Flow](#4.-control-flow)

[if / else](#if-/-else)

[for (Go's only loop)](#for-\(go's-only-loop\))

[switch](#switch)

[defer](#defer)

[5\. Functions](#5.-functions)

[6\. Arrays, Slices & Maps](#6.-arrays,-slices-&-maps)

[Arrays (fixed size)](#arrays-\(fixed-size\))

[Slices (dynamic, most common)](#slices-\(dynamic,-most-common\))

[Maps](#maps)

[7\. Structs & Methods](#7.-structs-&-methods)

[Methods](#methods)

[Embedding (composition over inheritance)](#embedding-\(composition-over-inheritance\))

[8\. Interfaces](#8.-interfaces)

[Empty Interface](#empty-interface)

[Type Assertion](#type-assertion)

[Stringer Interface (fmt)](#stringer-interface-\(fmt\))

[9\. Pointers](#9.-pointers)

[10\. Error Handling](#10.-error-handling)

[Idiomatic Error Pattern](#idiomatic-error-pattern)

[panic & recover](#panic-&-recover)

[11\. Goroutines & Channels (Concurrency)](#11.-goroutines-&-channels-\(concurrency\))

[Goroutines](#goroutines)

[Channels](#channels)

[select](#select)

[sync.WaitGroup](#sync.waitgroup)

[sync.Mutex](#sync.mutex)

[12\. Packages & Modules](#12.-packages-&-modules)

[Visibility](#visibility)

[13\. Building a REST API (Patterns)](#13.-building-a-rest-api-\(patterns\))

[Using net/http (stdlib)](#using-net/http-\(stdlib\))

[JSON Struct Tags](#json-struct-tags)

[Common HTTP Pattern](#common-http-pattern)

[14\. Common Interview Questions](#14.-common-interview-questions)

[Q: What is a goroutine? How is it different from a thread?](#q:-what-is-a-goroutine?-how-is-it-different-from-a-thread?)

[Q: What is the difference between a slice and an array?](#q:-what-is-the-difference-between-a-slice-and-an-array?)

[Q: How does Go handle inheritance?](#q:-how-does-go-handle-inheritance?)

[Q: What is nil in Go?](#q:-what-is-nil-in-go?)

[Q: When would you use a pointer receiver vs value receiver?](#q:-when-would-you-use-a-pointer-receiver-vs-value-receiver?)

[Q: What is the select statement?](#q:-what-is-the-select-statement?)

[Q: How does error handling work in Go?](#q:-how-does-error-handling-work-in-go?)

[Q: What is defer?](#q:-what-is-defer?)

[Q: What is the difference between make and new?](#q:-what-is-the-difference-between-make-and-new?)

[Q: Can Go return multiple values? How is it used?](#q:-can-go-return-multiple-values?-how-is-it-used?)

[Q: What are Go interfaces and how is satisfaction determined?](#q:-what-are-go-interfaces-and-how-is-satisfaction-determined?)

[Q: What is a channel and when would you use it?](#q:-what-is-a-channel-and-when-would-you-use-it?)

[Q: How do you avoid race conditions in Go?](#q:-how-do-you-avoid-race-conditions-in-go?)

[15\. Garbage Collection & Memory Management](#15.-garbage-collection-&-memory-management)

[The Stack vs The Heap](#the-stack-vs-the-heap)

[Escape Analysis](#escape-analysis)

[Go's Garbage Collector — The Tri-Color Mark-and-Sweep](#go's-garbage-collector-—-the-tri-color-mark-and-sweep)

[The Three Colors](#the-three-colors)

[GC Phases (simplified)](#gc-phases-\(simplified\))

[The Write Barrier](#the-write-barrier)

[GC Tuning — GOGC](#gc-tuning-—-gogc)

[Memory Allocation Internals](#memory-allocation-internals)

[sync.Pool — Reusing Allocations](#sync.pool-—-reusing-allocations)

[Finalizers](#finalizers)

[Common Memory Mistakes & How to Fix Them](#common-memory-mistakes-&-how-to-fix-them)

[1\. Goroutine Leaks (most common\!)](#1.-goroutine-leaks-\(most-common!\))

[2\. Slice Memory Retention](#2.-slice-memory-retention)

[3\. Map Growth Without Shrink](#3.-map-growth-without-shrink)

[4\. String ↔ Byte Slice Conversions](#4.-string-↔-byte-slice-conversions)

[Profiling Memory — pprof](#profiling-memory-—-pprof)

[runtime Package Snippets](#runtime-package-snippets)

[GC & Memory Interview Questions](#gc-&-memory-interview-questions)

[Quick Reference Cheatsheet](#quick-reference-cheatsheet)

---

## ***1\. Why Go?*** {#1.-why-go?}

* *Compiled, statically typed, garbage-collected*  
* *Built-in concurrency (goroutines \+ channels)*  
* *Fast compilation, simple syntax, opinionated formatting (`gofmt`)*  
* *No classes — uses structs \+ interfaces instead*  
* *Developed at Google; used by Docker, Kubernetes, Terraform*

---

## ***2\. Program Structure*** {#2.-program-structure}

*package main          // Every executable must have package main*

*import (*

    *"fmt"*

    *"math"*

*)*

*func main() {         // Entry point*

    *fmt.Println("Hello, Go\!")*

*}*

***Key points:***

* *`package main` \+ `func main()` \= executable*  
* *Other packages \= libraries*  
* *Unused imports → **compile error***  
* *`_` used to blank-import for side effects: `import _ "pkg"`*

---

## ***3\. Variables & Types*** {#3.-variables-&-types}

### ***Declaration styles*** {#declaration-styles}

*// Explicit type*

*var name string \= "Alice"*

*// Type inferred*

*var age \= 30*

*// Short declaration (inside functions only)*

*city := "Delhi"*

*// Multiple*

*var x, y int \= 1, 2*

*// Zero values (default when not initialised)*

*var i int     // 0*

*var f float64 // 0.0*

*var b bool    // false*

*var s string  // ""*

### ***Basic Types*** {#basic-types}

| *Type* | *Description* | *Example* |
| ----- | ----- | ----- |
| *`int`, `int8/16/32/64`* | *Integer* | *`42`* |
| *`uint`* | *Unsigned int* | *`10`* |
| *`float32`, `float64`* | *Decimal* | *`3.14`* |
| *`bool`* | *Boolean* | *`true`* |
| *`string`* | *UTF-8 text* | *`"hello"`* |
| *`byte`* | *Alias for uint8* | *`'A'`* |
| *`rune`* | *Alias for int32 (Unicode)* | *`'🎉'`* |

### ***Type Conversion (explicit only)*** {#type-conversion-(explicit-only)}

*var x int \= 42*

*var f float64 \= float64(x)   // Must be explicit — no implicit conversion*

### ***Constants*** {#constants}

*const Pi \= 3.14159*

*const (*

    *StatusOK    \= 200*

    *StatusNotFound \= 404*

*)*

*// iota — auto-incrementing constant*

*type Direction int*

*const (*

    *North Direction \= iota  // 0*

    *East                    // 1*

    *South                   // 2*

    *West                    // 3*

*)*

---

## ***4\. Control Flow*** {#4.-control-flow}

### ***if / else*** {#if-/-else}

*if x \> 10 {*

    *fmt.Println("big")*

*} else if x \> 5 {*

    *fmt.Println("medium")*

*} else {*

    *fmt.Println("small")*

*}*

*// Init statement in if*

*if err := doSomething(); err \!= nil {*

    *log.Fatal(err)*

*}*

### ***for (Go's only loop)*** {#for-(go's-only-loop)}

*// Classic*

*for i := 0; i \< 5; i++ { }*

*// While-style*

*for x \< 100 { x \*= 2 }*

*// Infinite*

*for { }*

*// Range over slice*

*for i, v := range slice { }*

*// Range over map*

*for k, v := range myMap { }*

*// Range — index only*

*for i := range slice { }*

*// Range — value only*

*for \_, v := range slice { }*

### ***switch*** {#switch}

*switch day {*

*case "Mon", "Tue":*

    *fmt.Println("Weekday")*

*case "Sat", "Sun":*

    *fmt.Println("Weekend")*

*default:*

    *fmt.Println("Unknown")*

*}*

*// No condition — like if-else chain*

*switch {*

*case x \< 0:*

    *fmt.Println("negative")*

*case x \== 0:*

    *fmt.Println("zero")*

*}*

***Note:** Go switch does NOT fall through by default. Use `fallthrough` keyword to force it.*

### ***defer*** {#defer}

*func readFile() {*

    *f, \_ := os.Open("file.txt")*

    *defer f.Close()   // Runs when function returns (LIFO order)*

    *// ... work with f*

*}*

---

## ***5\. Functions*** {#5.-functions}

*// Basic*

*func add(a int, b int) int {*

    *return a \+ b*

*}*

*// Same-type params shorthand*

*func add(a, b int) int { return a \+ b }*

*// Multiple return values (idiomatic Go)*

*func divide(a, b float64) (float64, error) {*

    *if b \== 0 {*

        *return 0, errors.New("division by zero")*

    *}*

    *return a / b, nil*

*}*

*// Named return values*

*func minMax(nums \[\]int) (min, max int) {*

    *min, max \= nums\[0\], nums\[0\]*

    *for \_, n := range nums {*

        *if n \< min { min \= n }*

        *if n \> max { max \= n }*

    *}*

    *return // "naked return"*

*}*

*// Variadic*

*func sum(nums ...int) int {*

    *total := 0*

    *for \_, n := range nums { total \+= n }*

    *return total*

*}*

*sum(1, 2, 3\)*

*sum(nums...)  // Unpack a slice*

*// First-class functions*

*fn := func(x int) int { return x \* x }*

*apply := func(f func(int) int, v int) int { return f(v) }*

---

## ***6\. Arrays, Slices & Maps*** {#6.-arrays,-slices-&-maps}

### ***Arrays (fixed size)*** {#arrays-(fixed-size)}

*var arr \[3\]int           // \[0 0 0\]*

*arr2 := \[3\]string{"a", "b", "c"}*

*arr3 := \[...\]int{1, 2, 3}  // Compiler counts length*

*Arrays are **value types** — copying an array copies all elements.*

### ***Slices (dynamic, most common)*** {#slices-(dynamic,-most-common)}

*// From array*

*arr := \[5\]int{1, 2, 3, 4, 5}*

*s := arr\[1:3\]   // \[2 3\] — shares underlying array*

*// Make*

*s := make(\[\]int, 3\)       // len=3, cap=3*

*s := make(\[\]int, 3, 10\)   // len=3, cap=10*

*// Literal*

*s := \[\]int{1, 2, 3}*

*// Append*

*s \= append(s, 4, 5\)*

*s \= append(s, other...)   // Append another slice*

*// Copy*

*dst := make(\[\]int, len(src))*

*copy(dst, src)*

***Slice internals:** A slice header holds `(pointer, length, capacity)`. Slices share the underlying array until `append` triggers a reallocation.*

### ***Maps*** {#maps}

*// Declaration*

*m := map\[string\]int{"a": 1, "b": 2}*

*m := make(map\[string\]int)*

*// Set / Get*

*m\["key"\] \= 42*

*val := m\["key"\]*

*// Check existence (comma ok idiom)*

*val, ok := m\["key"\]*

*if \!ok { fmt.Println("not found") }*

*// Delete*

*delete(m, "key")*

*// Iterate*

*for k, v := range m { }*

*Maps are **reference types**. Zero value of a map is `nil` — writing to nil map panics.*

---

## ***7\. Structs & Methods*** {#7.-structs-&-methods}

*// Define*

*type Person struct {*

    *Name string*

    *Age  int*

*}*

*// Instantiate*

*p := Person{Name: "Alice", Age: 30}*

*p2 := Person{"Bob", 25}   // Positional (fragile, avoid)*

*var p3 Person              // Zero value*

*// Access*

*fmt.Println(p.Name)*

*// Pointer to struct*

*ptr := \&Person{Name: "Carol"}*

*ptr.Age \= 40   // Go auto-dereferences*

*// Anonymous struct*

*config := struct {*

    *Host string*

    *Port int*

*}{"localhost", 8080}*

### ***Methods*** {#methods}

*// Value receiver — copy of struct*

*func (p Person) Greet() string {*

    *return "Hi, I'm " \+ p.Name*

*}*

*// Pointer receiver — can mutate original*

*func (p \*Person) Birthday() {*

    *p.Age++*

*}*

*p := Person{Name: "Alice", Age: 30}*

*p.Birthday()           // Go auto-takes address*

*fmt.Println(p.Greet())*

***Rule of thumb:** Use pointer receivers when you need to mutate the struct, or when the struct is large (avoids copying).*

### ***Embedding (composition over inheritance)*** {#embedding-(composition-over-inheritance)}

*type Animal struct {*

    *Name string*

*}*

*func (a Animal) Speak() string { return a.Name \+ " speaks" }*

*type Dog struct {*

    *Animal         // Embedded — promotes methods*

    *Breed string*

*}*

*d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Lab"}*

*fmt.Println(d.Speak())  // Promoted method*

---

## ***8\. Interfaces*** {#8.-interfaces}

*type Shape interface {*

    *Area() float64*

    *Perimeter() float64*

*}*

*type Rectangle struct { Width, Height float64 }*

*func (r Rectangle) Area() float64      { return r.Width \* r.Height }*

*func (r Rectangle) Perimeter() float64 { return 2 \* (r.Width \+ r.Height) }*

*// Rectangle implicitly satisfies Shape — no "implements" keyword*

*var s Shape \= Rectangle{Width: 5, Height: 3}*

*fmt.Println(s.Area())*

### ***Empty Interface*** {#empty-interface}

*var anything interface{}    // Holds any value*

*anything \= 42*

*anything \= "hello"*

*// Modern alias (Go 1.18+)*

*var v any \= 3.14*

### ***Type Assertion*** {#type-assertion}

*var i interface{} \= "hello"*

*s, ok := i.(string)    // ok \= true, s \= "hello"*

*n, ok := i.(int)       // ok \= false, n \= 0*

*// Type switch*

*switch v := i.(type) {*

*case string:*

    *fmt.Println("string:", v)*

*case int:*

    *fmt.Println("int:", v)*

*default:*

    *fmt.Println("unknown")*

*}*

### ***Stringer Interface (fmt)*** {#stringer-interface-(fmt)}

*func (p Person) String() string {*

    *return fmt.Sprintf("%s (%d)", p.Name, p.Age)*

*}*

*// fmt.Println(p) now uses your String() method*

---

## ***9\. Pointers*** {#9.-pointers}

*x := 42*

*p := \&x        // p is \*int (pointer to x)*

*fmt.Println(\*p) // Dereference: 42*

*\*p \= 100        // Mutates x*

*fmt.Println(x)  // 100*

*// new() — allocates zeroed value, returns pointer*

*p2 := new(int)  // \*int pointing to 0*

***Why pointers?***

* *Avoid copying large structs*  
* *Mutate values in functions*  
* *Share a single instance*

*Go has **no pointer arithmetic** (unlike C).*

---

## ***10\. Error Handling*** {#10.-error-handling}

*Go does not use exceptions. Errors are **values**.*

*// errors.New*

*err := errors.New("something went wrong")*

*// fmt.Errorf — with context*

*err := fmt.Errorf("open file: %w", originalErr)*

*// Custom error type*

*type ValidationError struct {*

    *Field   string*

    *Message string*

*}*

*func (e \*ValidationError) Error() string {*

    *return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)*

*}*

*// Checking error types*

*var ve \*ValidationError*

*if errors.As(err, \&ve) {*

    *fmt.Println("field:", ve.Field)*

*}*

*// Unwrap / Is*

*if errors.Is(err, os.ErrNotExist) {*

    *fmt.Println("file not found")*

*}*

### ***Idiomatic Error Pattern*** {#idiomatic-error-pattern}

*result, err := doSomething()*

*if err \!= nil {*

    *return fmt.Errorf("context: %w", err)*

*}*

*// Use result*

### ***panic & recover*** {#panic-&-recover}

*// panic — for unrecoverable errors (avoid for normal flow)*

*panic("something went horribly wrong")*

*// recover — catch a panic (used in middleware, frameworks)*

*func safeDiv(a, b int) (result int, err error) {*

    *defer func() {*

        *if r := recover(); r \!= nil {*

            *err \= fmt.Errorf("recovered: %v", r)*

        *}*

    *}()*

    *return a / b, nil*

*}*

---

## ***11\. Goroutines & Channels (Concurrency)*** {#11.-goroutines-&-channels-(concurrency)}

### ***Goroutines*** {#goroutines}

*go func() {*

    *fmt.Println("running concurrently")*

*}()*

*go doWork(arg)   // Any function call prefixed with go*

*Goroutines are lightweight threads managed by the Go runtime (not OS threads).*

### ***Channels*** {#channels}

*// Unbuffered — send blocks until receiver is ready*

*ch := make(chan int)*

*go func() { ch \<- 42 }()*

*val := \<-ch*

*// Buffered — send only blocks when buffer full*

*ch := make(chan string, 3\)*

*ch \<- "a"*

*ch \<- "b"*

*ch \<- "c"*

*// Close & range*

*close(ch)*

*for msg := range ch { fmt.Println(msg) }*

*// Check if closed*

*val, ok := \<-ch   // ok \= false if closed and empty*

### ***select*** {#select}

*select {*

*case msg := \<-ch1:*

    *fmt.Println("ch1:", msg)*

*case msg := \<-ch2:*

    *fmt.Println("ch2:", msg)*

*case \<-time.After(1 \* time.Second):*

    *fmt.Println("timeout")*

*default:*

    *fmt.Println("no channel ready")*

*}*

### ***sync.WaitGroup*** {#sync.waitgroup}

*var wg sync.WaitGroup*

*for i := 0; i \< 5; i++ {*

    *wg.Add(1)*

    *go func(id int) {*

        *defer wg.Done()*

        *fmt.Println("worker", id)*

    *}(i)*

*}*

*wg.Wait()*

### ***sync.Mutex*** {#sync.mutex}

*var mu sync.Mutex*

*var counter int*

*mu.Lock()*

*counter++*

*mu.Unlock()*

*// Or with defer*

*mu.Lock()*

*defer mu.Unlock()*

---

## ***12\. Packages & Modules*** {#12.-packages-&-modules}

*\# Initialise a module*

*go mod init github.com/yourname/project*

*\# Add a dependency*

*go get github.com/some/package*

*\# Tidy dependencies*

*go mod tidy*

*\# Build*

*go build ./...*

*\# Run*

*go run main.go*

*\# Test*

*go test ./...*

*\# Format*

*gofmt \-w .*

***File: `go.mod`***

*module github.com/yourname/project*

*go 1.21*

*require (*

    *github.com/gin-gonic/gin v1.9.1*

*)*

### ***Visibility*** {#visibility}

* ***Exported** (public): starts with **uppercase** — `MyFunc`, `MyStruct`*  
* ***Unexported** (private): starts with **lowercase** — `myHelper`, `internalType`*

---

## ***13\. Building a REST API (Patterns)*** {#13.-building-a-rest-api-(patterns)}

### ***Using `net/http` (stdlib)*** {#using-net/http-(stdlib)}

*package main*

*import (*

    *"encoding/json"*

    *"net/http"*

*)*

*type User struct {*

    *ID   int    \`json:"id"\`*

    *Name string \`json:"name"\`*

*}*

*func getUsers(w http.ResponseWriter, r \*http.Request) {*

    *users := \[\]User{{1, "Alice"}, {2, "Bob"}}*

    *w.Header().Set("Content-Type", "application/json")*

    *json.NewEncoder(w).Encode(users)*

*}*

*func main() {*

    *http.HandleFunc("/users", getUsers)*

    *http.ListenAndServe(":8080", nil)*

*}*

### ***JSON Struct Tags*** {#json-struct-tags}

*type Product struct {*

    *ID       int     \`json:"id"\`*

    *Name     string  \`json:"name"\`*

    *Price    float64 \`json:"price,omitempty"\`  // omit if zero*

    *internal string  // unexported — excluded from JSON*

*}*

*// Decode request body*

*var p Product*

*json.NewDecoder(r.Body).Decode(\&p)*

*// Encode response*

*json.NewEncoder(w).Encode(p)*

### ***Common HTTP Pattern*** {#common-http-pattern}

*func handler(w http.ResponseWriter, r \*http.Request) {*

    *switch r.Method {*

    *case http.MethodGet:*

        *// handle GET*

    *case http.MethodPost:*

        *// handle POST*

    *default:*

        *http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)*

    *}*

*}*

---

## ***14\. Common Interview Questions*** {#14.-common-interview-questions}

### ***Q: What is a goroutine? How is it different from a thread?*** {#q:-what-is-a-goroutine?-how-is-it-different-from-a-thread?}

*Goroutines are user-space threads managed by Go's scheduler (M:N threading model). They start with \~2KB stack (vs \~1MB for OS threads) and can scale to millions. The scheduler multiplexes goroutines onto OS threads.*

### ***Q: What is the difference between a slice and an array?*** {#q:-what-is-the-difference-between-a-slice-and-an-array?}

*Arrays have fixed size and are value types (copying copies data). Slices are dynamic, reference the underlying array, and carry `(ptr, len, cap)`. Slices are far more commonly used.*

### ***Q: How does Go handle inheritance?*** {#q:-how-does-go-handle-inheritance?}

*Go uses **composition, not inheritance**. You embed structs to promote fields/methods and satisfy interfaces to achieve polymorphism. No `extends` or `implements` keywords.*

### ***Q: What is `nil` in Go?*** {#q:-what-is-nil-in-go?}

*`nil` is the zero value for pointers, slices, maps, channels, functions, and interfaces. It is NOT a type — its meaning depends on context.*

### ***Q: When would you use a pointer receiver vs value receiver?*** {#q:-when-would-you-use-a-pointer-receiver-vs-value-receiver?}

*Use a pointer receiver when: (1) the method needs to mutate the struct, (2) the struct is large (avoid copying), (3) consistency with other methods on the type. Use value receiver for small read-only structs.*

### ***Q: What is the `select` statement?*** {#q:-what-is-the-select-statement?}

*`select` is like a `switch` for channels — it blocks until one of its cases can proceed. If multiple cases are ready, one is chosen randomly. A `default` case makes it non-blocking.*

### ***Q: How does error handling work in Go?*** {#q:-how-does-error-handling-work-in-go?}

*Errors are values implementing the `error` interface (`Error() string`). Functions return `(result, error)` and callers check `if err != nil`. Use `errors.Is()` to match specific errors and `errors.As()` to extract error types.*

### ***Q: What is `defer`?*** {#q:-what-is-defer?}

*`defer` schedules a function call to run just before the enclosing function returns. Multiple defers execute in LIFO order. Commonly used for cleanup (close files, unlock mutexes).*

### ***Q: What is the difference between `make` and `new`?*** {#q:-what-is-the-difference-between-make-and-new?}

*`new(T)` allocates a zeroed value of type T and returns a pointer (`*T`). `make(T, ...)` is only for slices, maps, and channels — it allocates and **initialises** the internal data structure, returning `T` (not a pointer).*

### ***Q: Can Go return multiple values? How is it used?*** {#q:-can-go-return-multiple-values?-how-is-it-used?}

*Yes. The most common pattern is returning `(value, error)`. Named return values allow a "naked return" but should be used sparingly for clarity.*

### ***Q: What are Go interfaces and how is satisfaction determined?*** {#q:-what-are-go-interfaces-and-how-is-satisfaction-determined?}

*An interface is a set of method signatures. A type **implicitly satisfies** an interface by implementing all its methods — no explicit declaration needed. This enables duck typing with static type safety.*

### ***Q: What is a channel and when would you use it?*** {#q:-what-is-a-channel-and-when-would-you-use-it?}

*A channel is a typed conduit for communication between goroutines. Use channels to send data or signals between goroutines. Prefer channels over shared memory; use `sync.Mutex` only when sharing state is simpler.*

### ***Q: How do you avoid race conditions in Go?*** {#q:-how-do-you-avoid-race-conditions-in-go?}

*Use channels for communication, `sync.Mutex` / `sync.RWMutex` for shared state, or `sync/atomic` for simple counters. Run `go test -race` to detect race conditions.*

---

## ***15\. Garbage Collection & Memory Management*** {#15.-garbage-collection-&-memory-management}

*Go manages memory for you — but knowing how it works makes you a better Go developer and impresses interviewers.*

---

### ***The Stack vs The Heap*** {#the-stack-vs-the-heap}

*Go allocates memory in two places:*

|  | *Stack* | *Heap* |
| ----- | ----- | ----- |
| ***What lives here*** | *Local variables, function args* | *Long-lived / shared objects* |
| ***Who manages it*** | *Compiler (automatic)* | *Garbage Collector* |
| ***Speed*** | *Very fast (just move a pointer)* | *Slower (GC overhead)* |
| ***Size*** | *Limited (\~1MB default, goroutine stack starts at \~2KB)* | *Large (limited by RAM)* |

*func add(a, b int) int {*

    *result := a \+ b   // result lives on the STACK — gone when function returns*

    *return result*

*}*

*func newUser(name string) \*User {*

    *u := User{Name: name}  // u ESCAPES to the HEAP — pointer outlives the function*

    *return \&u*

*}*

***Key insight:** You don't choose stack vs heap — the **compiler decides** via escape analysis.*

---

### ***Escape Analysis*** {#escape-analysis}

*The compiler checks at compile time: "Does this variable's lifetime exceed its function?"*

* *If **yes** → allocates on the **heap** (GC will clean it up)*  
* *If **no** → allocates on the **stack** (free at function return)*

*// See escape analysis output:*

*// go build \-gcflags="-m" ./...*

*func stackAlloc() int {*

    *x := 42        // x does NOT escape → stack*

    *return x*

*}*

*func heapAlloc() \*int {*

    *x := 42        // x ESCAPES to heap (pointer returned)*

    *return \&x*

*}*

*func interfaceEscape(v interface{}) {*

    *// Storing in interface{} often causes escape to heap*

    *// because the compiler can't know the concrete type's size*

    *fmt.Println(v)  // v may escape here*

*}*

***Practical tip:** Avoid unnecessary pointers and `interface{}` in hot paths — they force heap allocation.*

---

### ***Go's Garbage Collector — The Tri-Color Mark-and-Sweep*** {#go's-garbage-collector-—-the-tri-color-mark-and-sweep}

*Go uses a **concurrent, tri-color mark-and-sweep GC** (introduced in Go 1.5, refined since).*

#### ***The Three Colors*** {#the-three-colors}

*Every heap object is conceptually one of:*

* *⬜ **White** — not yet visited; candidates for collection*  
* *🔘 **Grey** — discovered but children not yet scanned*  
* *⬛ **Black** — fully scanned; reachable, will NOT be collected*

#### ***GC Phases (simplified)*** {#gc-phases-(simplified)}

*1\. STW (Stop-the-World) — very brief*

   *└─ Mark setup: enable write barrier, snapshot goroutine stacks*

*2\. Concurrent Mark (runs alongside your code)*

   *└─ Start from roots (globals, stack vars, registers)*

   *└─ Turn roots Grey → scan them → turn Black*

   *└─ Repeat until no Grey objects remain*

*3\. STW — brief*

   *└─ Mark termination: flush remaining work*

*4\. Concurrent Sweep*

   *└─ Reclaim all White (unreachable) objects*

   *└─ Runs concurrently — memory is reused incrementally*

***Go's GC goal:** Keep STW pauses under **1 millisecond**. In practice, pauses are often 100–500 microseconds.*

---

### ***The Write Barrier*** {#the-write-barrier}

*A problem with concurrent GC: your program might create a new reference to a white object while the GC is running, causing it to be incorrectly swept.*

*The **write barrier** is a small piece of code injected by the compiler around every pointer write. It tells the GC "hey, this pointer just changed" so the GC can re-shade objects appropriately.*

*// You write this:*

*obj.Field \= newPtr*

*// Compiler generates something like:*

*writeBarrier(obj, \&obj.Field, newPtr)  // notify GC*

*obj.Field \= newPtr*

*The write barrier is active only during the mark phase. It has a small but measurable CPU cost — one reason to minimise pointer-heavy data structures in performance-critical code.*

---

### ***GC Tuning — `GOGC`*** {#gc-tuning-—-gogc}

*Go exposes one main GC knob: `GOGC` (default \= **100**).*

*GOGC \= 100  →  GC runs when heap size doubles from last collection*

*GOGC \= 200  →  GC runs when heap grows 3× (less frequent, more memory used)*

*GOGC \= 50   →  GC runs more frequently (lower latency, more CPU used)*

*GOGC \= off  →  GC disabled entirely (dangerous — use only in tests/benchmarks)*

*import "runtime/debug"*

*// At runtime:*

*debug.SetGCPercent(200)  // same as GOGC=200*

*// Force a GC cycle manually (rarely needed):*

*runtime.GC()*

***Go 1.19+** introduced `GOMEMLIMIT` — a hard cap on total memory:*

*GOMEMLIMIT=512MiB ./myapp   \# GC will be more aggressive before hitting this limit*

---

### ***Memory Allocation Internals*** {#memory-allocation-internals}

*Go's allocator (inspired by tcmalloc) uses size classes for efficiency:*

*Tiny objects  (\< 16 bytes)  → packed together into a single block*

*Small objects (≤ 32KB)      → allocated from per-P (processor) caches (mcache)*

*Large objects (\> 32KB)      → allocated directly from the heap (mheap)*

***Why this matters:** Allocating many tiny objects (e.g. in a tight loop) is fast because it mostly touches thread-local caches — no locking needed.*

---

### ***sync.Pool — Reusing Allocations*** {#sync.pool-—-reusing-allocations}

*`sync.Pool` lets you reuse objects to reduce GC pressure:*

*var pool \= sync.Pool{*

    *New: func() interface{} {*

        *return make(\[\]byte, 1024\)  // create a fresh buffer if pool is empty*

    *},*

*}*

*func handleRequest() {*

    *buf := pool.Get().(\[\]byte)   // grab from pool (or create new)*

    *defer pool.Put(buf)          // return to pool when done*

    *// use buf...*

*}*

***Gotcha:** `sync.Pool` objects can be collected by the GC between GC cycles. Don't store anything you can't recreate. Never use it as a cache with TTL expectations.*

---

### ***Finalizers*** {#finalizers}

*Go allows you to attach a function that runs before an object is garbage collected:*

*type Resource struct { fd int }*

*func NewResource(fd int) \*Resource {*

    *r := \&Resource{fd: fd}*

    *runtime.SetFinalizer(r, func(r \*Resource) {*

        *syscall.Close(r.fd)  // cleanup when GC collects r*

    *})*

    *return r*

*}*

***Avoid finalizers in production code.** They are non-deterministic (you don't know when GC runs), can cause issues in test suites, and prevent objects from being collected in the same GC cycle as their dependents. Use `defer` and explicit `Close()` instead.*

---

### ***Common Memory Mistakes & How to Fix Them*** {#common-memory-mistakes-&-how-to-fix-them}

#### ***1\. Goroutine Leaks (most common\!)*** {#1.-goroutine-leaks-(most-common!)}

*// ❌ BAD — goroutine blocks forever if no one reads ch*

*func leak() {*

    *ch := make(chan int)*

    *go func() {*

        *ch \<- expensiveWork()  // blocks forever if caller exits*

    *}()*

    *// function returns, nobody reads ch — goroutine leaks*

*}*

*// ✅ GOOD — use context for cancellation*

*func noLeak(ctx context.Context) (int, error) {*

    *ch := make(chan int, 1\)  // buffered so goroutine doesn't block*

    *go func() {*

        *ch \<- expensiveWork()*

    *}()*

    *select {*

    *case result := \<-ch:*

        *return result, nil*

    *case \<-ctx.Done():*

        *return 0, ctx.Err()*

    *}*

*}*

#### ***2\. Slice Memory Retention*** {#2.-slice-memory-retention}

*// ❌ BAD — small slice holds reference to huge underlying array*

*func getFirst(data \[\]byte) \[\]byte {*

    *return data\[:10\]   // data (e.g. 10MB) never freed\!*

*}*

*// ✅ GOOD — copy the data you need*

*func getFirst(data \[\]byte) \[\]byte {*

    *result := make(\[\]byte, 10\)*

    *copy(result, data\[:10\])*

    *return result*

*}*

#### ***3\. Map Growth Without Shrink*** {#3.-map-growth-without-shrink}

*// Maps grow but NEVER shrink in Go*

*// If you add 1M keys then delete them, the map still holds that memory*

*// ✅ Fix: replace the map with a new one after large deletions*

*m \= make(map\[string\]int)  // old map gets GC'd*

#### ***4\. String ↔ Byte Slice Conversions*** {#4.-string-↔-byte-slice-conversions}

*// Each conversion allocates — avoid in hot paths*

*b := \[\]byte(s)   // allocates*

*s := string(b)   // allocates*

*// ✅ Use strings.Builder for building strings in loops*

*var sb strings.Builder*

*for \_, word := range words {*

    *sb.WriteString(word)*

*}*

*result := sb.String()*

---

### ***Profiling Memory — `pprof`*** {#profiling-memory-—-pprof}

*import \_ "net/http/pprof"   // side-effect import adds /debug/pprof endpoints*

*func main() {*

    *go http.ListenAndServe(":6060", nil)  // expose pprof*

    *// ... your app*

*}*

*\# Capture a heap profile*

*go tool pprof http://localhost:6060/debug/pprof/heap*

*\# Inside pprof shell:*

*(pprof) top10          \# top 10 memory consumers*

*(pprof) list myFunc    \# annotated source for myFunc*

*(pprof) web            \# open flame graph in browser*

*\# Allocation profile (shows where allocations happen)*

*go tool pprof http://localhost:6060/debug/pprof/allocs*

*\# Run benchmarks with memory stats*

*go test \-bench=. \-benchmem ./...*

---

### ***`runtime` Package Snippets*** {#runtime-package-snippets}

*import "runtime"*

*// Print GC stats*

*var stats runtime.MemStats*

*runtime.ReadMemStats(\&stats)*

*fmt.Printf("Heap in use:  %v MB\\n", stats.HeapInuse/1024/1024)*

*fmt.Printf("Total alloc:  %v MB\\n", stats.TotalAlloc/1024/1024)*

*fmt.Printf("GC cycles:    %v\\n",   stats.NumGC)*

*fmt.Printf("Last GC pause:%v\\n",   time.Duration(stats.PauseNs\[(stats.NumGC+255)%256\]))*

*// Number of goroutines (watch for leaks)*

*fmt.Println("Goroutines:", runtime.NumGoroutine())*

*// GOMAXPROCS — how many OS threads run goroutines (default \= num CPUs)*

*runtime.GOMAXPROCS(4)*

---

### ***GC & Memory Interview Questions*** {#gc-&-memory-interview-questions}

***Q: Does Go have a garbage collector?** Yes. Go uses a concurrent tri-color mark-and-sweep GC. It runs mostly in parallel with your program, targeting sub-millisecond stop-the-world pauses.*

***Q: What is escape analysis?** A compile-time analysis that decides whether a variable lives on the stack or the heap. If a variable's lifetime exceeds its function (e.g. its address is returned), it "escapes" to the heap where the GC manages it.*

***Q: How do you reduce GC pressure?** Reuse objects with `sync.Pool`, avoid unnecessary heap allocations (prefer value types over pointers where sensible), use `strings.Builder` instead of string concatenation, and profile with `pprof` to find hot allocation sites.*

***Q: What is GOGC?** An environment variable (default 100\) that controls GC frequency. It means "run GC when heap has grown X% since last collection." Higher \= less frequent GC, more memory used. Lower \= more frequent GC, lower peak memory.*

***Q: What is a goroutine leak?** A goroutine that is blocked indefinitely (usually waiting on a channel or lock that will never be signalled). Leaking goroutines consume memory and CPU. Detect with `runtime.NumGoroutine()` and fix with context cancellation or buffered channels.*

***Q: Why can't you rely on finalizers in Go?** Finalizers run non-deterministically — you cannot predict when the GC will collect an object. They also prevent the object from being collected in the same cycle as objects it references. Use `defer` and explicit `Close()` methods instead.*

***Q: What is `sync.Pool` used for?** Temporarily pooling objects to avoid repeated allocation/GC cycles. Ideal for buffers and temporary objects used in request handlers. Note: pooled objects may be collected at any GC cycle, so they're not a persistent cache.*

---

## ***Quick Reference Cheatsheet*** {#quick-reference-cheatsheet}

*Variables:    var x int \= 5  |  x := 5*

*Zero values:  0, 0.0, false, ""*

*Loops:        for { }  |  for i := 0; i \< n; i++  |  for i, v := range s { }*

*Functions:    func f(a, b int) (int, error) { }*

*Structs:      type S struct { Field Type }*

*Methods:      func (s \*S) Method() { }*

*Interfaces:   type I interface { Method() }*

*Goroutine:    go func() { }()*

*Channel:      ch := make(chan T)  |  ch \<- val  |  val := \<-ch*

*Error:        if err \!= nil { return fmt.Errorf("context: %w", err) }*

*Defer:        defer cleanup()*

*GC tune:      GOGC=100 (default) | GOMEMLIMIT=512MiB*

*Pool:         sync.Pool{New: func() any { return \&MyObj{} }}*

*Mem profile:  go tool pprof http://localhost:6060/debug/pprof/heap*

*Goroutines:   runtime.NumGoroutine()*

---

*Good luck with your interviews\! 🚀*

