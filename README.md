# Math Skills Project
This project calculates key statistical metrics—such as Average, Median, Variance, Standard Deviation, and Square Root—from a dataset stored in a text file. Each function is implemented separately, and the data is read from the file. The results are then displayed in the terminal.

## Features
#### Average: Calculates the mean of the dataset.
#### Median: Finds the middle value of the dataset.
#### Variance: Measures how far each value is from the mean.
#### Standard Deviation: Quantifies the spread of the dataset using variance.
#### Square Root: Used to calculate the standard deviation.
#### File Input: Reads data from a text file.
## File Structure

<pre>
math-skills/
├── functions/             # Contains statistical function implementations
│   ├── Average.go         # Calculates the average of the dataset
│   ├── Deviation.go       # Calculates the standard deviation
│   ├── Median.go          # Finds the median of the dataset
│   ├── Slice.go           # Reads and processes the dataset from a file
│   ├── sumOfSquares.go            # Calculates square roots
│   ├── Variance.go        # Calculates the variance of the dataset
├── go.mod                 # Go module file
├── main.go                # Main program that prints the calculated results
└── README.md              # Project description and usage instructions
</pre>
## Schema of Function Workflow

<pre>
+---------------+      +-----------------+      +----------------+
| Input (File)  +----->|   Slice.go       +----->| Output Results |
| (Data.txt)    |      | (Reads Data)     |      |   (Main.go)    |
+---------------+      +-----------------+      +----------------+
                               |
                               v
                        +--------------+
                        | Average.go   |
                        +--------------+
                               |
                               v
                        +--------------+
                        | Median.go    |
                        +--------------+
                               |
                               v
                        +--------------+
                        | Variance.go  |
                        +--------------+
                               |
                               v
                        +--------------+
                        | Deviation.go |
                        +--------------+
</pre>

