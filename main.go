package main

import (
	"flag"
	"io"
	"log"
	"os"
)

var (
	startOffset = flag.Int64("start", 0, "start offset of chunk")
	endOffset   = flag.Int64("end", -1, "end offset of chunk, this is an exclusive offset (<0 indicated eof)")
	filePath    = flag.String("file", "", "path of input file")
	outputPath  = flag.String("output", "", "if specified, chunk will be written to the file at this path, stdout otherwise")
)

func main() {
	var output io.Writer
	flag.Parse()
	if *filePath == "" {
		log.Fatal("-file must be specified")
	}
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("failed to open file (%s): %v", *filePath, err)
	}
	defer file.Close()
	if *outputPath != "" {
		outFile, err := os.Create(*outputPath)
		if err != nil {
			log.Fatalf("failed to open file for output (%s): %v", *outputPath, err)
		}
		defer outFile.Close()
		output = outFile
	} else {
		output = os.Stdout
	}
	if *endOffset < 0 {
		n, err := file.Seek(0, 2)
		if err != nil {
			log.Fatalf("failed to see to end of file: %v", err)
		}
		*endOffset = n
	}
	if n, err := file.Seek(*startOffset, 0); err != nil {
		log.Fatalf("failed to seek to offset: %v", err)
	} else if n != *startOffset {
		log.Fatalf("failed to proper seek to start offset")
	}
	chunkSize := *endOffset - *startOffset
	lr := io.LimitReader(file, chunkSize)
	if n, err := io.Copy(output, lr); err != nil {
		log.Fatalf("encountered error while copy chunk: %v", err)
	} else if n != chunkSize {
		log.Fatalf("failed to copy all requested bytes: %d bytes copied, %d bytes requested", n, chunkSize)
	}
}
