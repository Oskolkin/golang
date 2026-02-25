package main

import "fmt"

type Options struct {
	Count      bool
	OnlyDup    bool
	OnlyUniq   bool
	SkipFields int
	SkipChars  int
	IgnoreCase bool
}

func (o Options) Validate() error {
	if o.SkipFields < 0 {
		return fmt.Errorf("SkipFields must be >= 0")
	}
	if o.SkipChars < 0 {
		return fmt.Errorf("SkipChars must be >= 0")
	}
	if o.OnlyDup && o.OnlyUniq {
		return fmt.Errorf("OnlyDup and OnlyUniq cannot both be true")
	}
	return nil
}

type Result struct {
	Lines  []string // что печатаем (в исходном порядке)
	Counts []int    // сколько раз встретилось
}

// UniqAllLines — чистая функция: принимает строки и опции, возвращает результат.
// Её и будем unit-тестировать.
func UniqAllLines(lines []string, opts Options) (Result, error) {
	if err := opts.Validate(); err != nil {
		return Result{}, err
	}

	stats := make(map[string]*stat)
	order := make([]string, 0, 128)

	for _, orig := range lines {
		key := buildKey(orig, opts.SkipFields, opts.SkipChars, opts.IgnoreCase)

		if st, ok := stats[key]; ok {
			st.count++
		} else {
			stats[key] = &stat{count: 1, firstLine: orig}
			order = append(order, key)
		}
	}

	out := Result{}
	for _, key := range order {
		st := stats[key]

		if opts.OnlyDup && st.count < 2 {
			continue
		}
		if opts.OnlyUniq && st.count != 1 {
			continue
		}

		out.Lines = append(out.Lines, st.firstLine)
		out.Counts = append(out.Counts, st.count)
	}

	return out, nil
}
