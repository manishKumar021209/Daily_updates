# Daily_updates

# XKCD Comic Downloader

This Go program fetches data from the XKCD webcomic API and saves it to a JSON file.

## Prerequisites

To run this program, you need to have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

## Usage

1. Clone or download this repository to your local machine.

2. Open a terminal and navigate to the directory where the code is located.

3. Build the Go program using the following command:

   ```bash
   go build xkcd.go
--> Run the program using the following command:  
./xkcd
This will fetch XKCD comic data and save it to a file named "xkcd.json."


Configuration
You can configure the number of jobs and workers in the main function to control the number of comics to fetch and the concurrency level.

// Number of jobs (comics) to fetch
noOfJobs := 3000

// Number of workers (concurrent fetches)
noOfWorkers := 100
