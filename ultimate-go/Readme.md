### Directory Hierarchy
put any kinds of example in the examples directory
```
ulimate-go
├── Readme.md
└── examples
```

Details regarding the **"Ultimate Go Programming"** video lessons including it's table of contents can be found [here](https://www.oreilly.com/library/view/ultimate-go-programming/9780134757476/).

# Note Section
## Lesson 1: Design Guidelines

### 1.1: Prepare Your Mind
- Software Engineering is about designing and architecting.
- Back off of being impressed with large amounts of code.
- We don't need high levels of abstraction for decoupling.
- The code runs on the hardware and the hardware is what determines how fast your code runs. Programming languages like Java, Ruby, etc are based on the concept of virtual machine but Go is based on the real machine.
- Think in terms of Cost Vs Benefit! NOT Good vs Bad! Think about what kind of benefit we're adding for what cost!
- Loves the two quotes: "Technology changes quickly" but "But our minds changes slowly."
- Start thinking about what we're doing!

### 1.2: Productivity Vs Performance
- Does performance matter? Yes. Is it your number 1 priority? No. Because if it was, people would be looking into Assembly languages since Assembly languages give you full performance. So, nowadays, instead of performance, productivity is the number 1 priority. So it's Performance Vs Productivity. GO allows us to have both Performance as well as Productivity.
- Go is just a tool but we ourselves should develop the sound engineering practices to build reliable software.

### 1.3: Correctness Vs Performance
- Software engineers usually write code and then decide to optimize the code for performance. This is wrong since until and unless you don't benchmark and profile it, you're just guessing.
- First optimize the code for correctness, then benchmark the code and profile them, then look at if the code is fast enough and finally optimize the code for performance. 
- Simple straightforward code is just eaiser to work with and less likely to have problems.
- We all need to write code where we can maintain a mental model of what's going on.
- When we're making decisions around simplicity, readability and minimizing, when we write code with consistency in mind, when we write code that is being designed and architected and written for maintainability first, then we're optimizing for correctness.

### 1.4: Code Reviews
- Integrity must be first. It is built in inside Go. Integrity means we write code that is reliable, we need to take reliability seriously.
- Every code that we write does one of 3 things: allocate memory, read memory and write memory. Integrity is about making sure all 3 of these operations are accurate, consistent and efficient. Every code we write is also a data transformation.
- For maintaining integrity at a macro level, we should:
  - Write less code: Writing less code means less bugs
  - Error handling: It **MUST** be a part of the code you're writing.
- Every code we write should be built upon
  1. Integrity
  2. Readability
    - Readable code means an average developer on your team should be able to understand what each piece of code is doing logically.
    - Readable code also mean that every peace of code does not hide its cost and impact on the machine.
  3. Simplicity
    - It's about hiding complexity.
    - Go's garbage collector hides the complexity of memory management. 
  4. Performance
    - Never guess about performance.
    - Profile before you decide something is performance critical.
    - It's about being mechanically sympathetic. Go allows us to be sympathetic.

## Lesson 2: Language Syntax
## Lesson 3: Data Structures
## Lesson 4: Decoupling
## Lesson 5: Composition
## Lesson 6: Error Handling
## Lesson 7: Packaging
## Lesson 8: Goroutines
## Lesson 9: Data Races
## Lesson 10: Channels
## Lesson 11: Concurrency Patterns
## Lesson 12: Testing
## Lesson 13: Profiling