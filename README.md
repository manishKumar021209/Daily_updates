# Daily_updates
# Pizza Ordering Simulation in Go

This is a simple Go program that simulates a pizza ordering system. It demonstrates the use of goroutines and channels to manage pizza orders, track their status, and handle potential failures during the pizza-making process.

## Features

- Randomly generates pizza orders with success or failure outcomes.
- Simulates pizza making with random delays.
- Tracks the number of pizzas made and failed.
- Prints colorful messages to the console to indicate the status of each pizza order.
- Provides a summary of the day's pizza production at the end.

## Getting Started

To run the pizza ordering simulation, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/pizza-ordering-simulation.git
2.Change into the project directory: cd pizza-ordering-simulation
3.Build and run the Go program:go run main.go
4.Observe the output in the terminal to see the pizza orders being processed.

Code Structure
->main.go: The main program file that contains the pizza ordering simulation logic.
->README.md: This file, providing documentation for the code.
->go.mod and go.sum: Go module files for dependency management.

Author
Manish

Acknowledgments
->The code uses the github.com/fatih/color package for colorful console output.
->It demonstrates Go concurrency features such as goroutines and channels.
->Inspiration for the project may come from real-world pizza ordering systems.
