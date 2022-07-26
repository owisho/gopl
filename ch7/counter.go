package ch7

import "bufio"

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	start := 0
	for start < len(p) {
		step, _, err := bufio.ScanWords(p[start:], true)
		if err != nil {
			return 0, err
		}
		start += step
		*c += WordCounter(1)
	}
	return int(*c), nil
}
