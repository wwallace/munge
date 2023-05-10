package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func swapCase(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsUpper(r) {
			result = append(result, unicode.ToLower(r))
		} else if unicode.IsLower(r) {
			result = append(result, unicode.ToUpper(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func munge(outputWriter io.Writer, wrd string, level int) {
	wordList := []string{wrd, strings.ToUpper(wrd), strings.Title(wrd)}

	if level > 2 {
		wordList = append(wordList, swapCase(strings.Title(wrd)))
	}

	leetReplacements := []struct {
		old, new string
	}{
		{"e", "3"},
		{"a", "4"},
		{"o", "0"},
		{"i", "!"},
		{"i", "1"},
		{"l", "1"},
		{"a", "@"},
		{"s", "$"},
	}

	if level > 4 {
		for _, r := range leetReplacements {
			wordList = append(wordList, strings.ReplaceAll(wrd, r.old, r.new))
		}
		temp := wrd
		for _, r := range leetReplacements {
			temp = strings.ReplaceAll(temp, r.old, r.new)
		}
		wordList = append(wordList, temp)
	}

	if level > 4 {
		for _, r := range leetReplacements {
			temp := strings.Title(wrd)
			temp = strings.ReplaceAll(temp, r.old, r.new)
			wordList = append(wordList, temp)
		}
	}

	for _, word := range wordList {
		fmt.Fprintln(outputWriter, word)
	}
}

func mungeword(outputWriter io.Writer, wrd string, level int) {
	suffixes := []string{"1", "123456", "12", "2", "123", "!", "."}

	if level > 4 {
		for _, suffix := range suffixes {
			munge(outputWriter, wrd+suffix, level)
		}
	}

	if level > 6 {
		for _, suffix := range []string{"?", "_", "0", "01", "69", "21", "22", "23", "1234", "8", "9", "10", "11", "13", "3", "4", "5", "6", "7"} {
			munge(outputWriter, wrd+suffix, level)
		}
	}

	if level > 7 {
		for _, suffix := range []string{"07", "08", "09", "14", "15", "16", "17", "18", "19", "24", "77", "88", "99", "12345", "123456789"} {
			munge(outputWriter, wrd+suffix, level)
		}
	}

	if level > 8 {
		for _, suffix := range []string{"00", "02", "03", "04", "05", "06", "19", "20", "25", "26", "27", "28", "007", "1234567", "12345678", "111111", "111", "777", "666", "101", "33", "44", "55", "66", "2020", "2021", "2022", "2023", "86", "87", "89", "90", "91", "92", "93", "94", "95", "98"} {
			munge(outputWriter, wrd+suffix, level)
		}
	}
}

func main() {
	wordPtr := flag.String("word", "", "word to munge")
	levelPtr := flag.Int("level", 5, "munge level [0-9] (default 5)")
	inputPtr := flag.String("input", "", "input file")
	outputPtr := flag.String("output", "", "output file")
	flag.Parse()

	var outputWriter io.Writer = os.Stdout

	if *outputPtr != "" {
		file, err := os.Create(*outputPtr)
		if err != nil {
			fmt.Println("Exiting\nCould not write file:", *outputPtr)
			os.Exit(1)
		}
		defer file.Close()
		outputWriter = bufio.NewWriter(file)
	}

	if *wordPtr != "" {
		mungeword(outputWriter, strings.ToLower(*wordPtr), *levelPtr)
	} else if *inputPtr != "" {
		file, err := os.Open(*inputPtr)
		if err != nil {
			fmt.Println("Exiting\nCould not read file:", *inputPtr)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			wrd := strings.ToLower(scanner.Text())
			mungeword(outputWriter, wrd, *levelPtr)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Exiting\nError scanning input file:", *inputPtr)
			os.Exit(1)
		}
	} else {
		fmt.Println("Nothing to do!!\nTry -h for help.\n")
		os.Exit(0)
	}

	if bufferedWriter, ok := outputWriter.(*bufio.Writer); ok {
		bufferedWriter.Flush()
	}
}
