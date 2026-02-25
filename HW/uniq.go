package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func buildKey(line string, skipFields, skipChars int, ignoreCase bool) string {
	orig := line
	s := line

	// -f: пропустить первые N полей (слов)
	if skipFields > 0 {
		fields := strings.Fields(s)
		if skipFields >= len(fields) {
			s = ""
		} else {
			s = strings.Join(fields[skipFields:], " ")
		}
	}

	// -s: пропустить первые N символов (после -f), UTF-8 безопасно
	if skipChars > 0 && s != "" {
		r := []rune(s)
		if skipChars >= len(r) {
			s = ""
		} else {
			s = string(r[skipChars:])
		}
	}

	// -i: игнорировать регистр (только для сравнения)
	if ignoreCase {
		s = strings.ToLower(s)
	}

	// Чтобы "Thanks." (после -f 1 => ключ "") не слипался с реально пустыми строками
	if s == "" && strings.TrimSpace(orig) != "" {
		return "\x00" + orig
	}

	return s
}

type stat struct {
	count     int
	firstLine string
}

func main() {
	countFlag := flag.Bool("c", false, "Вывести количество повторений перед строкой")
	repeatFlag := flag.Bool("d", false, "Вывести только повторяющиеся строки (count>=2)")
	uniqFlag := flag.Bool("u", false, "Вывести только уникальные строки (count==1)")
	skipFields := flag.Int("f", 0, "Не учитывать первые num_fields полей")
	skipChars := flag.Int("s", 0, "Не учитывать первые num_chars символов (после -f)")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр букв")

	flag.Parse()

	// --------- ВХОД (stdin или файл) ----------
	var in io.Reader = os.Stdin
	var inFile *os.File

	if flag.NArg() >= 1 {
		inputName := flag.Arg(0)
		f, err := os.Open(inputName)
		if err != nil {
			log.Fatal(err)
		}
		inFile = f
		defer inFile.Close()
		in = inFile
	}

	// --------- ВЫХОД (stdout или файл) ----------
	var out io.Writer = os.Stdout
	var outFile *os.File

	if flag.NArg() >= 2 {
		outputName := flag.Arg(1)
		f, err := os.Create(outputName) // создаст/перезапишет файл
		if err != nil {
			log.Fatal(err)
		}
		outFile = f
		defer outFile.Close()
		out = outFile
	}

	// Буферизуем вывод (быстрее и удобно)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	stats := make(map[string]*stat)
	order := make([]string, 0, 128)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		orig := scanner.Text()
		key := buildKey(orig, *skipFields, *skipChars, *ignoreCase)

		if st, ok := stats[key]; ok {
			st.count++
		} else {
			stats[key] = &stat{count: 1, firstLine: orig}
			order = append(order, key)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Печатаем в исходном порядке
	for _, key := range order {
		st := stats[key]

		if *repeatFlag && st.count < 2 {
			continue
		}
		if *uniqFlag && st.count != 1 {
			continue
		}

		if *countFlag {
			fmt.Fprintf(writer, "%d %s\n", st.count, st.firstLine)
		} else {
			fmt.Fprintln(writer, st.firstLine)
		}
	}
}
