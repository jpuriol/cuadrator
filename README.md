# Cuadrator

Cuadrator is a tool written in Go to easily manage and generate shift quadrants. It allows defining participants, shifts, and occupations using YAML files and generating a PDF document with the result.

## Available Commands

The tool offers several commands to interact with the quadrant data:

- `check`: Verifies if the information contained in the YAML files is valid and consistent.
- `pdf`: Generates a PDF document with the quadrant information.
- `stats`: Shows statistics about participant distribution across shifts and occupations.
- `tui`: Interactive Terminal User Interface to browse participants and their assigned occupations.

To see the detailed help for any command, you can run:
```bash
./cuadrator [command] --help
```

## Configuration

The project uses three main configuration files in YAML format:

### `participants.yaml`
Contains the list of names of the participants who can be assigned to the quadrant.

Example:
```yaml
- Alice
- Bob
- Charlie
```

### `schema.yaml`
Defines the structure of the quadrant:
- **name**: Internal name of the quadrant (used for the PDF filename).
- **title**: Title shown in the exported PDF.
- **subtitle**: Subtitle shown in the exported PDF.
- **shifts**: Definition of the shift schedules.
- **occupations**: List of the different tasks or occupations available.

Example:
```yaml
name: "Shift_Schedule_2023"
title: "Shift Schedule 2023"
subtitle: "First Quarter"
shifts:
  1: "08:00-12:00"
  2: "12:00-16:00"
occupations:
  1: "Reception"
  2: "Security"
```

### `schedule.yaml`
Contains the actual assignment of participants to the shifts and occupations defined in the schema.

Example:
```yaml
1: # Shift ID
  1: # Occupation ID
    - [Alice]
    - [Bob]
  2:
    - [Charlie, Dave] # A team of two
2:
  1:
    - [Eve]
```

## Installation and Usage

If you have Go installed, you can compile the project by running:

```bash
go build -o cuadrator ./cmd/cuadrator/*.go
```

Then, you can run the commands mentioned above:

```bash
./cuadrator tui
./cuadrator check
./cuadrator pdf
```

## Author
Developed by jpuriol.
