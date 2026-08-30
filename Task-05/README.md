# Task 05 - Grand Line Guardian

## Program Description

Grand Line Guardian is a terminal-based process monitoring program made
using Python. It reads information from the Linux `/proc` filesystem
and shows the currently active processes.

The program displays the process ID, process name, CPU usage, memory
usage, and total number of active processes. The display keeps updating
so that the user can see the process information in real time.

I also added simple user input so that the program is not completely
fixed. When the program starts, the user can choose how the processes
should be sorted and can choose the refresh interval.

## Features

- Shows PID of active processes
- Shows process name
- Shows CPU usage
- Shows memory usage in MB
- Shows total active process count
- Refreshes continuously
- Allows sorting by CPU usage, memory usage, or PID
- Allows the user to choose a refresh interval between 0.5 and 1 second
- Handles processes that disappear while the program is reading them
- Runs directly in the Linux terminal

## Programming Language

Python 3

## Main Approach

The program uses the Linux `/proc` virtual filesystem instead of using
a separate process-monitoring library.

The `/proc` directory contains a folder for each active process, where
the folder name is the process ID. I scan these folders and use the
process information files to get the required values.

### Process Information

The `get_processes()` function checks the `/proc` directory and keeps
only numeric folder names because those represent process IDs.

For each process, I read `/proc/PID/stat`. From this information I get
the process name and CPU time used by that process.

### CPU Usage

The `get_total_cpu_time()` function reads `/proc/stat` to get the
total CPU time of the system.

The program takes the process CPU time and total CPU time at two
different points. It compares the difference between the two readings
to calculate an approximate CPU usage percentage.

This is why the program needs to wait for a short interval before
displaying the next reading.

### Memory Usage

The `get_memory()` function reads `/proc/PID/statm`.

The value from this file is based on memory pages, so the program also
gets the system page size and converts the result into MB before
displaying it.

### User Input

The `get_user_settings()` function is used at the beginning of the
program.

The user can select:

1. CPU Usage
2. Memory Usage
3. PID

The user can also enter a refresh interval between 0.5 and 1 second.

The program checks the input and asks again if an invalid value is
entered.

### Sorting

The process list is sorted based on the option selected by the user.

If CPU Usage is selected, processes using more CPU are shown first.

If Memory Usage is selected, processes using more memory are shown
first.

If PID is selected, the processes are shown in PID order.

### Continuous Monitoring

The main program uses a loop to keep checking the processes.

After every refresh interval, the process information is collected
again and the terminal display is updated.

The user can stop the program by pressing `Ctrl+C`.

## Files

- `main.py` - Main Python program
- `README.md` - Description and explanation of the program
- `requirements.txt` - Dependency information

## How to Run

Open the terminal inside the task folder and run:

```bash
python3 main.py
```

Then select the sorting option and enter the refresh interval.

To stop the program:

```text
Ctrl+C
```

## Requirements

This program is intended to run on Linux with Python 3. It uses only
Python standard library modules, so no external Python packages are
needed.

## What I Learned

Through this task, I learned how Linux provides process information
through the `/proc` virtual filesystem. I also learned how to read
system files using Python, calculate CPU usage using two readings,
convert memory values, handle changing processes, and continuously
update information in a terminal program.

The user input and sorting options also helped me understand how to
make a system-monitoring program a little more interactive instead of
keeping all settings fixed.
