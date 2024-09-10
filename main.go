package main

import (
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"path/filepath"

	"github.com/dihedron/entropy/version"
)

func main() {

	if len(os.Args) == 2 && os.Args[1] == "version" {
		//fmt.Printf("entropy v%s.%s.%s\n", version.VersionMajor, version.VersionMinor, version.VersionPatch)
		fmt.Printf("%s v%s.%s.%s (%s/%s built with %s on %s)\n", version.Name, version.VersionMajor, version.VersionMinor, version.VersionPatch,
			version.GoOS, version.GoArch, version.GoVersion, version.BuildTime)
		return
	}

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			file, err := os.Open(filepath.Clean(arg))
			if err != nil {
				slog.Error("error opening file", "path", arg, "error", err)
				continue
			}
			defer file.Close()
			entropy, err := compute(file)
			if err != nil {
				slog.Error("error computing entropy", "path", arg, "error", err)
				continue
			}
			fmt.Printf("%-32s%f\n", arg, entropy)
		}
	} else {
		entropy, err := compute(os.Stdin)
		if err != nil {
			slog.Error("error computing entropy", "error", err)
		}
		fmt.Printf("%f\n", entropy)
	}
}

func compute(stream io.Reader) (float64, error) {
	buffer := make([]byte, 64*1024*1024)
	index := [256]byte{}
	total := 0
	entropy := float64(0.0)

outer:
	for {
		read, err := stream.Read(buffer)
		switch err {
		case nil:
			total += read
			slog.Debug("read data", "count", read, "total", total)
			for i := 0; i < read; i++ {
				index[buffer[i]] += 1
			}
			continue
		case io.EOF:
			slog.Debug("end of stream reached", "total", total)
			for i, count := range index {
				if count == 0 {
					continue
				}
				probability := float64(count) / float64(total)
				entropy = entropy - probability*math.Log(probability)
				slog.Debug("computing entropy", "index", i, "count", count, "probability", probability, "entropy", entropy)
			}
			break outer
		default:
			slog.Error("error reading stream", "error", err)
			return 0.0, err
		}
	}
	slog.Debug("entropy computed", "entropy", entropy)
	return entropy / math.Log(256), nil
}
