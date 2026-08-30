package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

type Segment struct {
	PID   string
	Start int
	End   int
}

func readInt(reader *bufio.Reader, prompt string, minimum int) int {
	for {
		fmt.Print(prompt)

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		value, err := strconv.Atoi(text)
		if err == nil && value >= minimum {
			return value
		}

		fmt.Printf("Enter an integer greater than or equal to %d.\n", minimum)
	}
}

func addSegment(segments *[]Segment, pid string, start, end int) {
	if start >= end {
		return
	}

	if len(*segments) > 0 {
		last := &(*segments)[len(*segments)-1]

		if last.PID == pid && last.End == start {
			last.End = end
			return
		}
	}

	*segments = append(*segments, Segment{
		PID:   pid,
		Start: start,
		End:   end,
	})
}

func calculateTimes(processes []Process) {
	for i := range processes {
		processes[i].TurnaroundTime =
			processes[i].CompletionTime - processes[i].ArrivalTime

		processes[i].WaitingTime =
			processes[i].TurnaroundTime - processes[i].BurstTime

		if processes[i].WaitingTime < 0 {
			processes[i].WaitingTime = 0
		}
	}
}

func printGanttChart(segments []Segment) {
	fmt.Println("\nGantt Chart:")

	if len(segments) == 0 {
		fmt.Println("No execution segments.")
		return
	}

	for _, segment := range segments {
		fmt.Printf("| %-5s ", segment.PID)
	}
	fmt.Println("|")

	fmt.Printf("%-2d", segments[0].Start)

	for _, segment := range segments {
		fmt.Printf("%-8d", segment.End)
	}

	fmt.Println()
}

func printResults(processes []Process, segments []Segment, algorithm string) {
	calculateTimes(processes)

	fmt.Println("\n============================================================")
	fmt.Printf("Algorithm: %s\n", algorithm)
	fmt.Println("============================================================")

	printGanttChart(segments)

	sort.Slice(processes, func(i, j int) bool {
		return processes[i].Order < processes[j].Order
	})

	fmt.Println("\nProcess Details:")
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("%-8s %-8s %-8s %-12s %-10s %-12s\n",
		"PID", "AT", "BT", "CT", "WT", "TAT")
	fmt.Println("------------------------------------------------------------")

	totalWaiting := 0
	totalTurnaround := 0

	for _, p := range processes {
		fmt.Printf("%-8s %-8d %-8d %-12d %-10d %-12d\n",
			p.ID,
			p.ArrivalTime,
			p.BurstTime,
			p.CompletionTime,
			p.WaitingTime,
			p.TurnaroundTime)

		totalWaiting += p.WaitingTime
		totalTurnaround += p.TurnaroundTime
	}

	count := float64(len(processes))

	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Average Waiting Time    : %.2f\n",
		float64(totalWaiting)/count)
	fmt.Printf("Average Turnaround Time : %.2f\n",
		float64(totalTurnaround)/count)
	fmt.Println("============================================================")
}

func fcfs(input []Process) {
	processes := append([]Process(nil), input...)

	sort.SliceStable(processes, func(i, j int) bool {
		if processes[i].ArrivalTime == processes[j].ArrivalTime {
			return processes[i].Order < processes[j].Order
		}
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	currentTime := 0
	var segments []Segment

	for i := range processes {
		if currentTime < processes[i].ArrivalTime {
			currentTime = processes[i].ArrivalTime
		}

		start := currentTime
		currentTime += processes[i].BurstTime

		addSegment(
			&segments,
			processes[i].ID,
			start,
			currentTime,
		)

		processes[i].CompletionTime = currentTime
	}

	printResults(
		processes,
		segments,
		"FCFS (First Come First Serve)",
	)
}

func sjf(input []Process) {
	processes := append([]Process(nil), input...)
	n := len(processes)

	completed := make([]bool, n)
	completedCount := 0
	currentTime := 0

	var segments []Segment

	for completedCount < n {
		selected := -1

		for i := 0; i < n; i++ {
			if completed[i] {
				continue
			}

			if processes[i].ArrivalTime > currentTime {
				continue
			}

			if selected == -1 ||
				processes[i].BurstTime < processes[selected].BurstTime ||
				(processes[i].BurstTime == processes[selected].BurstTime &&
					processes[i].ArrivalTime < processes[selected].ArrivalTime) ||
				(processes[i].BurstTime == processes[selected].BurstTime &&
					processes[i].ArrivalTime == processes[selected].ArrivalTime &&
					processes[i].Order < processes[selected].Order) {

				selected = i
			}
		}

		if selected == -1 {
			nextArrival := -1

			for i := 0; i < n; i++ {
				if !completed[i] {
					if nextArrival == -1 ||
						processes[i].ArrivalTime < nextArrival {
						nextArrival = processes[i].ArrivalTime
					}
				}
			}

			currentTime = nextArrival
			continue
		}

		start := currentTime
		currentTime += processes[selected].BurstTime

		addSegment(
			&segments,
			processes[selected].ID,
			start,
			currentTime,
		)

		processes[selected].CompletionTime = currentTime
		completed[selected] = true
		completedCount++
	}

	printResults(
		processes,
		segments,
		"SJF (Shortest Job First - Non-Preemptive)",
	)
}

func roundRobin(input []Process, quantum int) {
	processes := append([]Process(nil), input...)

	sort.SliceStable(processes, func(i, j int) bool {
		if processes[i].ArrivalTime == processes[j].ArrivalTime {
			return processes[i].Order < processes[j].Order
		}
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	n := len(processes)

	for i := range processes {
		processes[i].RemainingTime = processes[i].BurstTime
	}

	queue := make([]int, 0)
	nextArrival := 0
	completed := 0
	currentTime := 0

	var segments []Segment

	for completed < n {

		for nextArrival < n &&
			processes[nextArrival].ArrivalTime <= currentTime {

			queue = append(queue, nextArrival)
			nextArrival++
		}

		if len(queue) == 0 {
			if nextArrival < n {
				currentTime = processes[nextArrival].ArrivalTime
				continue
			}
		}

		index := queue[0]
		queue = queue[1:]

		start := currentTime

		runTime := quantum

		if processes[index].RemainingTime < runTime {
			runTime = processes[index].RemainingTime
		}

		currentTime += runTime
		processes[index].RemainingTime -= runTime

		addSegment(
			&segments,
			processes[index].ID,
			start,
			currentTime,
		)

		for nextArrival < n &&
			processes[nextArrival].ArrivalTime <= currentTime {

			queue = append(queue, nextArrival)
			nextArrival++
		}

		if processes[index].RemainingTime > 0 {
			queue = append(queue, index)
		} else {
			processes[index].CompletionTime = currentTime
			completed++
		}
	}

	printResults(
		processes,
		segments,
		"Round Robin",
	)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("============================================================")
	fmt.Println("                 PIRATE KING'S SCHEDULER")
	fmt.Println("============================================================")

	numberOfProcesses := readInt(
		reader,
		"Enter number of processes: ",
		1,
	)

	processes := make([]Process, numberOfProcesses)

	fmt.Println("\nEnter process information:")

	for i := 0; i < numberOfProcesses; i++ {
		fmt.Printf("\nProcess %d\n", i+1)

		var id string

		for {
			fmt.Print("Process ID: ")

			idInput, _ := reader.ReadString('\n')
			id = strings.TrimSpace(idInput)

			if id != "" {
				break
			}

			fmt.Println("Process ID cannot be empty.")
		}

		arrivalTime := readInt(
			reader,
			"Arrival Time: ",
			0,
		)

		burstTime := readInt(
			reader,
			"Burst Time: ",
			1,
		)

		processes[i] = Process{
			ID:             id,
			ArrivalTime:    arrivalTime,
			BurstTime:      burstTime,
			RemainingTime:  burstTime,
			Order:          i,
		}
	}

	fmt.Println("\n============================================================")
	fmt.Println("Select Scheduling Algorithm")
	fmt.Println("============================================================")
	fmt.Println("1. FCFS")
	fmt.Println("2. SJF (Non-Preemptive)")
	fmt.Println("3. Round Robin")

	choice := readInt(
		reader,
		"Enter choice: ",
		1,
	)

	switch choice {
	case 1:
		fcfs(processes)

	case 2:
		sjf(processes)

	case 3:
		quantum := readInt(
			reader,
			"Enter Time Quantum: ",
			1,
		)

		roundRobin(processes, quantum)

	default:
		fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
	}
}
