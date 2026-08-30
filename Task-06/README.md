# Task 06 - Pirate King's Scheduler

## Introduction

For this task, I had to build a CPU scheduling simulator in Go.

The idea of the task is to take a set of processes that are waiting for CPU execution and simulate how different CPU scheduling algorithms will execute them.

I implemented three scheduling algorithms:

1. First Come First Serve (FCFS)
2. Shortest Job First (SJF - Non-Preemptive)
3. Round Robin (RR)

The program runs completely in the terminal. The user enters the process details, selects a scheduling algorithm, and the program shows the execution order, Gantt chart/timeline, process results, waiting time, turnaround time, average waiting time, and average turnaround time.

I kept the program simple enough for me to understand and explain, but I also tried to make it properly structured instead of putting the whole program inside the `main()` function.

---

## What I Used

- Programming language: Go (Golang)
- Go version used: Go 1.26.7
- Input/output: Terminal
- Go packages used:
  - `bufio`
  - `fmt`
  - `os`
  - `sort`
  - `strconv`
  - `strings`

I used Go because it is the required language for this task.

---

## What I Did First

I first made sure that Go was installed on my Ubuntu system.

I checked it using:

```bash
go version
```

After confirming that Go was installed, I created the Task 06 project directory and initialized it as a Go module.

The basic setup I used was:

```bash
cd ~/Documents/amfoss-tasks
mkdir Task06
cd Task06
go mod init pirate-scheduler
touch main.go
```

Then I opened the Go file and wrote the scheduler program.

I also used:

```bash
gofmt -w main.go
```

to format the Go code properly.

I checked that the program compiled using:

```bash
go build
```

and ran it using:

```bash
go run main.go
```

---

# Task Requirements

The task required the program to accept:

- Process ID
- Arrival Time
- Burst Time
- Time Quantum for Round Robin

The program also had to support:

- FCFS
- SJF Non-Preemptive
- Round Robin
- Gantt chart or timeline
- Waiting Time
- Turnaround Time
- Average Waiting Time
- Average Turnaround Time
- Terminal based output

The program implements all of these requirements.

---

# Program Structure

I divided the program into different parts based on what each part is responsible for.

I did this because it makes the code easier to understand and also makes it easier to check each scheduling algorithm separately.

The main parts are:

- `Process` structure
- `Segment` structure
- Input handling
- Gantt chart handling
- Time calculation
- FCFS function
- SJF function
- Round Robin function
- Result printing
- Main function

---

# Process Structure

I created a `Process` structure to keep all the information related to one process together.

```go
type Process struct {
    ID             string
    ArrivalTime    int
    BurstTime      int
    RemainingTime  int
    CompletionTime int
    WaitingTime    int
    TurnaroundTime int
    Order          int
}
```

I used the following fields:

### ID

This stores the process name, such as `P1`, `P2`, or `P3`.

### ArrivalTime

This stores the time at which the process becomes available for execution.

For example, if:

```text
P2 Arrival Time = 1
```

it means P2 becomes available at time 1.

### BurstTime

This is the total CPU time required by the process.

For example:

```text
P1 Burst Time = 5
```

means P1 needs 5 units of CPU time.

### RemainingTime

This is mainly needed for Round Robin.

Since Round Robin can stop a process before its complete burst time is finished, I need to remember how much CPU time is still left for that process.

### CompletionTime

This stores the time at which the process completely finishes.

### WaitingTime

This stores how long the process waited before getting its required CPU execution.

### TurnaroundTime

This stores the total time taken from the process arrival until its completion.

### Order

I used this to remember the original input order of processes.

It is useful when two processes have the same arrival time or when tie-breaking is needed.

---

# Segment Structure

I also created another structure:

```go
type Segment struct {
    PID   string
    Start int
    End   int
}
```

This is used for the Gantt chart.

For example, if P1 executes from time 0 to 5, a segment can store:

```text
PID = P1
Start = 0
End = 5
```

The different segments are later printed as the execution timeline.

---

# Input Handling

The program first asks for the number of processes.

Then for every process, it asks for:

```text
Process ID
Arrival Time
Burst Time
```

I used a separate `readInt()` function for integer input.

The reason for using a separate function is that I wanted the program to handle invalid numeric input instead of immediately failing.

For example, if a user enters something that is not an integer where an integer is expected, the program asks again.

I also check that:

- Number of processes is greater than 0
- Arrival time is not negative
- Burst time is greater than 0
- Time quantum is greater than 0
- Process ID is not empty

This makes the input handling a little safer.

---

# FCFS - First Come First Serve

FCFS means **First Come First Serve**.

The process that arrives first gets the CPU first.

For example:

```text
P1 arrives at 0
P2 arrives at 1
P3 arrives at 2
```

The order will be:

```text
P1 -> P2 -> P3
```

The important point about FCFS is that once a process starts executing, it completes before the next process starts.

In my program, I first sort the processes according to arrival time.

If two processes have the same arrival time, I use their original input order.

Then I go through the processes one by one.

If the CPU is currently idle because the next process has not arrived yet, the current time is moved to that process's arrival time.

After that, the burst time is added to the current time and the resulting time becomes the completion time.

The execution is also added to the Gantt chart segments.

---

# SJF - Shortest Job First

The second algorithm is **Shortest Job First**.

The version required in this task is **Non-Preemptive SJF**.

The basic idea is:

> Among the processes that have already arrived, choose the process with the shortest burst time.

For example, suppose the currently available processes are:

```text
P2 -> Burst Time 3
P3 -> Burst Time 2
```

P3 is selected because its burst time is smaller.

Non-preemptive means that once the selected process starts, it is allowed to finish completely before another process is selected.

In my program, I repeatedly check all processes that:

- Have not completed
- Have already arrived

From those processes, I select the one with the smallest burst time.

I also handle ties using arrival time and original input order.

If there is no process available at the current time, I move the current time to the next process arrival time.

This prevents the program from getting stuck when there is an idle period.

---

# Round Robin

The third algorithm is **Round Robin**.

Round Robin gives each process a fixed amount of CPU time called the **Time Quantum**.

For example, if the time quantum is:

```text
2
```

a process can execute for at most 2 time units during one turn.

If it still has remaining work after that, it goes back into the queue.

A simple example is:

```text
P1 -> P2 -> P3 -> P1 -> P2 -> ...
```

The actual order depends on arrival times and remaining burst times.

Round Robin is different from FCFS and non-preemptive SJF because a process can be stopped before its full burst time is completed.

This is why I added the `RemainingTime` field to the `Process` structure.

I also use a queue for Round Robin.

The basic flow is:

1. Add processes that have arrived to the queue.
2. Take the process at the front of the queue.
3. Run it for the time quantum or until it finishes.
4. Add newly arrived processes.
5. If the current process still has remaining time, put it at the back of the queue.
6. If it has finished, record its completion time.
7. Continue until all processes are completed.

---

# Gantt Chart / Timeline

The task asks for the execution order to be displayed using a simple Gantt chart or timeline.

I created a `Segment` structure and an `addSegment()` function for this.

For example, the output can look like:

```text
| P1    | P2    | P3    |
0       5       8       10
```

For Round Robin, it can show multiple executions of the same process:

```text
| P1 | P2 | P3 | P1 | P2 | P1 |
0    2    4    6    8    9    10
```

The `addSegment()` function also combines consecutive segments belonging to the same process. This keeps the Gantt chart cleaner.

---

# Completion Time

Completion Time means the time at which a process finishes completely.

For example:

```text
P1 starts at 0
P1 needs 5 units
```

so:

```text
Completion Time = 5
```

The program stores this in the `CompletionTime` field.

---

# Turnaround Time

I calculate turnaround time using:

```text
Turnaround Time = Completion Time - Arrival Time
```

For example:

```text
Arrival Time = 1
Completion Time = 8
```

then:

```text
Turnaround Time = 8 - 1
                = 7
```

The program calculates this for every process.

---

# Waiting Time

I calculate waiting time using:

```text
Waiting Time = Turnaround Time - Burst Time
```

For example:

```text
Turnaround Time = 7
Burst Time = 3
```

then:

```text
Waiting Time = 7 - 3
             = 4
```

The program calculates this after the scheduling algorithm finishes.

---

# Average Waiting Time

After calculating waiting time for every process, I add all the waiting times and divide them by the number of processes.

For example:

```text
Waiting times:
P1 = 0
P2 = 4
P3 = 6
```

Then:

```text
Average Waiting Time = (0 + 4 + 6) / 3
                     = 3.33
```

The program prints the result to two decimal places.

---

# Average Turnaround Time

The same approach is used for turnaround time.

For example:

```text
Turnaround times:
P1 = 5
P2 = 7
P3 = 8
```

Then:

```text
Average Turnaround Time = (5 + 7 + 8) / 3
                        = 6.67
```

---

# Why I Used Separate Functions

I did not put all the logic directly inside `main()`.

Instead, I separated the work into functions such as:

```text
fcfs()
sjf()
roundRobin()
calculateTimes()
printResults()
printGanttChart()
readInt()
addSegment()
```

Each function has a specific purpose.

For example:

- `fcfs()` handles FCFS scheduling.
- `sjf()` handles non-preemptive SJF.
- `roundRobin()` handles Round Robin.
- `calculateTimes()` calculates waiting and turnaround times.
- `printGanttChart()` displays the execution timeline.
- `printResults()` displays the final process table and averages.
- `readInt()` handles integer input and basic validation.
- `addSegment()` stores information needed for the Gantt chart.

This makes the program easier to follow and also makes it easier to test each algorithm.

---

# Go Concepts I Used

While doing this task, I used several basic and intermediate Go concepts.

## 1. Structs

I used structs for representing a process and a Gantt chart segment.

For example:

```go
type Process struct {
    ID          string
    ArrivalTime int
    BurstTime   int
}
```

This helped me keep related values together.

---

## 2. Slices

I used slices to store the list of processes and Gantt chart segments.

For example:

```go
processes := make([]Process, numberOfProcesses)
```

Slices are useful here because the number of processes is entered by the user.

---

## 3. Functions

I divided the program into functions instead of writing everything in one place.

This made the code more organized.

---

## 4. Loops

I used `for` loops to:

- Read process information
- Search for available processes
- Execute scheduling logic
- Calculate results
- Print the process table

---

## 5. Conditions

I used `if` statements for decisions such as:

- Whether a process has arrived
- Whether a process has completed
- Whether the CPU is idle
- Whether a process has remaining time
- Whether the selected algorithm is FCFS, SJF, or Round Robin

---

## 6. Switch Statement

The program uses a `switch` statement after the user selects the algorithm.

The options are:

```text
1. FCFS
2. SJF
3. Round Robin
```

The selected option calls the corresponding function.

---

## 7. Sorting

I used Go's `sort` package to arrange processes according to arrival time.

This is mainly useful for FCFS and Round Robin.

For SJF, I search through the available processes and select the shortest burst time.

---

## 8. Queue

Round Robin needs a queue because processes take turns using the CPU.

I used a slice as a simple queue.

The first available process is taken from the front, and if it is not finished, it is added back to the end.

---

## 9. Pointers

The `addSegment()` function receives a pointer to the slice of Gantt chart segments.

This allows the function to modify the original slice.

---

## 10. Error Handling for Input

I used `strconv.Atoi()` to convert string input into integers.

If the conversion fails, the program asks the user to enter the value again.

This was useful for making the program more reliable.

---

# Testing I Did

I tested all three scheduling algorithms using the same set of processes.

The test input was:

```text
Number of processes: 3

P1
Arrival Time: 0
Burst Time: 5

P2
Arrival Time: 1
Burst Time: 3

P3
Arrival Time: 2
Burst Time: 2
```

---

## FCFS Test

I selected:

```text
1
```

The execution order was:

```text
P1 -> P2 -> P3
```

The timeline was:

```text
0 -> 5 -> 8 -> 10
```

The results were:

```text
P1: CT = 5,  WT = 0, TAT = 5
P2: CT = 8,  WT = 4, TAT = 7
P3: CT = 10, WT = 6, TAT = 8
```

Average waiting time:

```text
3.33
```

Average turnaround time:

```text
6.67
```

---

## SJF Test

I selected:

```text
2
```

The execution order was:

```text
P1 -> P3 -> P2
```

The timeline was:

```text
0 -> 5 -> 7 -> 10
```

The results were:

```text
P1: CT = 5,  WT = 0, TAT = 5
P2: CT = 10, WT = 6, TAT = 9
P3: CT = 7,  WT = 3, TAT = 5
```

Average waiting time:

```text
3.00
```

Average turnaround time:

```text
6.33
```

---

## Round Robin Test

For Round Robin, I selected:

```text
3
```

and used:

```text
Time Quantum = 2
```

The execution order was:

```text
P1 -> P2 -> P3 -> P1 -> P2 -> P1
```

The timeline was:

```text
0 -> 2 -> 4 -> 6 -> 8 -> 9 -> 10
```

The results were:

```text
P1: CT = 10, WT = 5, TAT = 10
P2: CT = 9,  WT = 5, TAT = 8
P3: CT = 6,  WT = 2, TAT = 4
```

Average waiting time:

```text
4.00
```

Average turnaround time:

```text
7.33
```

I tested all three algorithms and checked that the calculated values matched the expected scheduling calculations.

---

# Example Program Flow

When the program starts, it shows:

```text
PIRATE KING'S SCHEDULER
```

Then it asks:

```text
Enter number of processes:
```

After that, it asks for each process:

```text
Process ID
Arrival Time
Burst Time
```

Once all processes are entered, it shows:

```text
1. FCFS
2. SJF (Non-Preemptive)
3. Round Robin
```

The user selects one.

If Round Robin is selected, the program additionally asks for:

```text
Enter Time Quantum:
```

After scheduling is complete, the program prints:

- Algorithm used
- Gantt chart
- Process details
- Completion Time
- Waiting Time
- Turnaround Time
- Average Waiting Time
- Average Turnaround Time

Everything is displayed directly in the terminal.

---

# Resources Used

I used the Go programming language documentation and the Go Tour/resource linked from the task page to understand the basic Go syntax and concepts needed for the program.

The main concepts I referred to were:

- Go variables and functions
- Structs
- Slices
- Loops and conditions
- Methods and basic Go programming structure
- Sorting
- Basic input/output
- Go modules

I also used the task description itself to understand the required scheduling algorithms and output.

---

# Things I Learned From This Task

The main thing I learned from this task is how CPU scheduling works by simulating it instead of only reading about the algorithms.

I understood that the three algorithms make different decisions.

FCFS mainly depends on arrival order.

SJF looks at the burst time of the processes that are currently available.

Round Robin gives every process a limited amount of CPU time and uses a queue to give other processes their turns.

I also understood better how waiting time and turnaround time are calculated from the execution timeline.

On the Go side, I got more practice with structs, slices, functions, loops, conditions, sorting, input handling, and using a queue-like structure.

Another useful part was separating the scheduling logic from the printing logic. It made it easier to test the algorithms individually.

---

# Assumptions

I made a few simple assumptions while implementing the simulator.

1. Process IDs are entered as non-empty values.
2. Arrival time cannot be negative.
3. Burst time must be greater than zero.
4. Time quantum must be greater than zero.
5. SJF is non-preemptive as required by the task.
6. If multiple processes have the same scheduling value, the program uses arrival time and then original input order as tie-breakers.
7. The program is intended to run in a terminal and does not require a GUI.
8. All calculations are based on integer time units.

---

# Files

The main project file is:

```text
main.go
```

The Go module information is stored in:

```text
go.mod
```

The README explains the approach, implementation, testing, Go concepts, and assumptions.

---

# How to Run

Inside the project directory:

```bash
go run main.go
```

To format the code:

```bash
gofmt -w main.go
```

To compile the program:

```bash
go build
```

Then the program can be run again using:

```bash
go run main.go
```

---

# Final Status

The actual scheduler implementation for Task 06 is complete.

I tested:

- FCFS
- SJF Non-Preemptive
- Round Robin

and checked their execution order and calculated values.

The program is terminal based, uses Go as required, and covers the main functionality mentioned in the task.
