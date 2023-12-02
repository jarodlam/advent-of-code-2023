# Advent of Code 2023
My solutions for Advent of Code 2023, written in Go.

## Usage

### Make commands
```sh
make build  # Build the aoc2023 binary
make clean  # Clean build files
make test   # Run all tests
make fmt    # Format code using gofumpt
```

### Running solutions
```sh
./aoc2023 -d <day> [-i input file]

# Examples:
./aoc2023 -d 1
./aoc2023 -d 2 -i ./data/my_custom_input.txt
```

### Development
Run from source:
```sh
go run ./cmd -d <day> [-i input file]

# Examples:
go run ./cmd -d 1
go run ./cmd -d 2 -i ./data/my_custom_input.txt
```
Run tests for a particular day:
```sh
go test ./pkg/day<day>

# Examples:
go test ./pkg/day01
```
