
# fbproc: FastBasic Pre-Processor

## Compilation

To compile the `fbproc` binary:

- On Linux: `go build -o fbproc`
- On Windows: `go build -o fbproc.exe`

## Usage

Use `fbproc` to preprocess FastBasic files.

### Command Line Syntax

```bash
fbproc -f <path_to_basic_file> -d <path_to_defines_file>
```

- `-f`: Path to the FastBasic file.
- `-d`: (Optional) Path to the defines file.

### Example Files

- **helloworld.bas** (FastBasic File):

  ```basic
  PRINT "Hello World! %GREETING%"
  #LOG#Processing HelloWorld.bas
  #INCLUDE#additional.bas#
  ```

- **defines.def** (Defines File):

  ```
  %GREETING%=from FastBasic
  ```

- **additional.bas** (Additional FastBasic File):

  ```basic
  PRINT "This is an included line."
  ```

### Running the Pre-Processor

Execute the command:

```bash
fbproc -f helloworld.bas -d defines.def
```

### Expected Output

- Console logs the message "Processing HelloWorld.bas".
- Output will be:

  ```basic
  Hello World! from FastBasic
  This is an included line.
  ```

## Notes

- Make sure the paths to the files are correct.
- Log messages are printed to the console, not in the output file.
- `#FILE#` directive will be replaced with the name of the current file.
- `#INCLUDE#` inserts the content of the included file at its position.
