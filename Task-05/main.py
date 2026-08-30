import os
import time


def get_processes():
    """Read active processes from the Linux /proc directory."""

    processes = {}

    for pid in os.listdir("/proc"):
        if not pid.isdigit():
            continue

        try:
            with open(f"/proc/{pid}/stat", "r") as file:
                data = file.read()

            start = data.find("(")
            end = data.rfind(")")

            name = data[start + 1:end]
            fields = data[end + 2:].split()

            # CPU time = user time + system time
            cpu_time = int(fields[11]) + int(fields[12])

            processes[int(pid)] = {
                "name": name,
                "cpu_time": cpu_time
            }

        except (FileNotFoundError, PermissionError, IndexError):
            # A process can disappear while it is being read.
            continue

    return processes


def get_total_cpu_time():
    """Get total CPU time of the system."""

    try:
        with open("/proc/stat", "r") as file:
            line = file.readline()

        values = line.split()[1:]
        return sum(int(value) for value in values)

    except (FileNotFoundError, ValueError):
        return 0


def get_memory(pid):
    """Get process memory usage in MB."""

    try:
        with open(f"/proc/{pid}/statm", "r") as file:
            resident_pages = int(file.read().split()[1])

        page_size = os.sysconf("SC_PAGE_SIZE")

        memory_mb = (
            resident_pages * page_size
        ) / (1024 * 1024)

        return memory_mb

    except (FileNotFoundError, PermissionError, IndexError, ValueError):
        return 0.0


def get_user_settings():
    """Get monitoring settings from the user."""

    print("=" * 60)
    print("             GRAND LINE GUARDIAN")
    print("=" * 60)

    print("\nChoose how to sort the processes:")
    print("1. CPU Usage")
    print("2. Memory Usage")
    print("3. PID")

    while True:
        choice = input("\nEnter your choice (1-3): ").strip()

        if choice in ["1", "2", "3"]:
            break

        print("Invalid choice. Please enter 1, 2 or 3.")

    while True:
        try:
            interval = float(
                input("\nEnter refresh interval (0.5 - 1 seconds): ")
            )

            if 0.5 <= interval <= 1:
                break

            print("Please enter a value between 0.5 and 1.")

        except ValueError:
            print("Please enter a valid number.")

    return choice, interval


def display_processes(
    current_processes,
    previous_processes,
    previous_total_cpu,
    sort_choice
):
    """Display current process information."""

    current_total_cpu = get_total_cpu_time()
    total_cpu_difference = current_total_cpu - previous_total_cpu

    process_list = []

    for pid, process in current_processes.items():
        cpu_usage = 0.0

        if pid in previous_processes:
            process_difference = (
                process["cpu_time"]
                - previous_processes[pid]["cpu_time"]
            )

            if total_cpu_difference > 0:
                cpu_usage = (
                    process_difference / total_cpu_difference
                ) * 100

        memory = get_memory(pid)

        process_list.append({
            "pid": pid,
            "name": process["name"],
            "cpu": cpu_usage,
            "memory": memory
        })

    # Sort according to the user's choice.
    if sort_choice == "1":
        process_list.sort(
            key=lambda process: process["cpu"],
            reverse=True
        )
    elif sort_choice == "2":
        process_list.sort(
            key=lambda process: process["memory"],
            reverse=True
        )
    else:
        process_list.sort(
            key=lambda process: process["pid"]
        )

    os.system("clear")

    print("=" * 85)
    print("                 GRAND LINE GUARDIAN")
    print("=" * 85)

    print(
        f"{'PID':<8}"
        f"{'PROCESS NAME':<30}"
        f"{'CPU %':<12}"
        f"{'MEMORY (MB)':<15}"
    )

    print("-" * 85)

    for process in process_list:
        print(
            f"{process['pid']:<8}"
            f"{process['name'][:28]:<30}"
            f"{process['cpu']:<12.2f}"
            f"{process['memory']:<15.2f}"
        )

    print("-" * 85)
    print(f"Total Active Processes: {len(process_list)}")
    print("Refresh Interval: 1 second or less")

    print("Sorted by: ", end="")

    if sort_choice == "1":
        print("CPU Usage")
    elif sort_choice == "2":
        print("Memory Usage")
    else:
        print("PID")

    print("Press Ctrl+C to exit.")

    return current_total_cpu


def main():
    sort_choice, refresh_interval = get_user_settings()

    print("\nStarting process monitor...")
    time.sleep(1)

    previous_processes = get_processes()
    previous_total_cpu = get_total_cpu_time()

    try:
        while True:
            time.sleep(refresh_interval)

            current_processes = get_processes()

            current_total_cpu = display_processes(
                current_processes,
                previous_processes,
                previous_total_cpu,
                sort_choice
            )

            previous_processes = current_processes
            previous_total_cpu = current_total_cpu

    except KeyboardInterrupt:
        print("\n\nGrand Line Guardian stopped.")
        print("Thank you for using the process monitor.")


if __name__ == "__main__":
    main()
