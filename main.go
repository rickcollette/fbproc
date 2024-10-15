package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Command-line arguments
	inputFilePath := flag.String("f", "", "Path to the BASIC file")
	definesFilePath := flag.String("d", "", "Path to the defines file")
	flag.Parse()

	// Read and process the input file
	processFile(*inputFilePath, *definesFilePath)
}

func processFile(inputFilePath, definesFilePath string) {
	// Extract base name (without extension) for log and output files
	baseName := strings.TrimSuffix(inputFilePath, filepath.Ext(inputFilePath))

	// Create the log file path
	logFilePath := baseName + ".fbp.log"
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()

	// Redirect log messages to the log file
	log := func(message string) {
		logFile.WriteString(message + "\n")
	}

	log("Processing file: " + inputFilePath)

	// Read the defines file and store key-value pairs
	defines, err := readDefines(definesFilePath)
	if err != nil {
		log("Error reading defines file: " + err.Error())
		return
	}

	// Open the BASIC file for reading
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		log("Error opening input file: " + err.Error())
		return
	}
	defer inputFile.Close()

	// Create a new file for the processed content
	outputFilePath := baseName + ".fbp"
	log("Creating output file: " + outputFilePath)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log("Error creating output file: " + err.Error())
		return
	}
	defer outputFile.Close()

	// Process each line and write to the output file
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processLine(line, defines, inputFilePath)
		if processedLine != "" {
			_, err := outputFile.WriteString(processedLine + "\n")
			if err != nil {
				log("Error writing to output file: " + err.Error())
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log("Error reading input file: " + err.Error())
		return
	}

	log("File processing completed successfully.")
}

func readDefines(definesFilePath string) (map[string]string, error) {
	// Create a map to store the defines
	defines := make(map[string]string)

	// Open the defines file
	file, err := os.Open(definesFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into key and value
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.Trim(parts[0], "%")
			value := parts[1]
			defines[key] = value
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return defines, nil
}

func processLine(line string, defines map[string]string, fileName string) string {
	// Handle #LOG# Directive
	if strings.Contains(line, "#LOG#") {
		logMessage := strings.TrimPrefix(line, "#LOG#")
		fmt.Println("Log:", logMessage)
		return "" // Return an empty string as the log line doesn't need to be in the output file
	}

	// Handle #FILE# Directive
	line = strings.ReplaceAll(line, "#FILE#", fileName)

	// Handle #INCLUDE# Directive
	if strings.Contains(line, "#INCLUDE#") {
		includeFile := strings.TrimSuffix(strings.TrimPrefix(line, "#INCLUDE#"), "#")
		includeContent, err := readIncludeFile(includeFile)
		if err != nil {
			fmt.Println("Error reading include file:", err)
			return ""
		}
		return includeContent
	}

	// Replace Defined Tags
	for key, value := range defines {
		tag := "%" + key + "%"
		line = strings.ReplaceAll(line, tag, value)
	}

	return line
}

func readIncludeFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content.String(), nil
}
