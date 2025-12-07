package day3

import (
	"fmt"
	"log/slog"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func RunChallange(input string, part int) {
	slog.Info("Hello from day2")

	switch part {
	case 1:
		part1(input)
	case 2:
		part2(input)
	}

}

func part1(input string) {
	var sum int
	for bank := range strings.SplitSeq(input, "\n") {
		slog.Info("Processing", "bank", bank)
		sum += joltageTwoDigit(bank)
	}
	slog.Info("Part one finished", "solution", sum)
}

func joltageTwoDigit(bank string) (jolts int) {
	for i := 0; i < len(bank)-1; i++ {
		for j := i + 1; j < len(bank); j++ {
			n := fmt.Sprintf("%c%c", bank[i], bank[j])
			d, _ := strconv.Atoi(n)
			if d > jolts {
				jolts = d
			}
		}
	}

	return
}

func part2(input string) {
	sum := part2Runner(input, runtime.NumCPU())
	slog.Info("Part two finished", "solution", sum)
}

func part2Runner(input string, runners int) int {

	var sum int

	// Create channels
	workCh := make(chan string, 200)
	sumCh := make(chan int, 200)
	var wg sync.WaitGroup

	for range runners {
		wg.Go(func() {
			worker(workCh, sumCh)
		})
	}

	// Send work to channel
	var count int
	for bank := range strings.SplitSeq(input, "\r\n") {
		workCh <- bank
		count++
	}
	close(workCh)

	wg.Wait()

	// Get result from all.
	for i := 0; i < count; i++ {
		result := <-sumCh
		slog.Info("Adding", "sum", result)
		sum += result
	}
	slog.Info("Done adding", "sum", sum)
	close(sumCh)

	return sum
}

func worker(workCh chan string, sumCh chan int) {
	slog.Info("Starting Working")
	for bank := range workCh {
		slog.Info("Processing", "bank", bank)
		bankInt, _ := strconv.Atoi(bank)
		sumCh <- joltageTwelveDigit(bankInt, len(bank))
	}
	slog.Info("Worker Finished, Returning")
}

// Dont Judge me!!
func joltageTwelveDigit(bank int, length int) (jolts int) {
	// speed := newSpeedemitor()
	for i := 0; i < length-11; i++ {
		for j := i + 1; j < length-10; j++ {
			for k := j + 1; k < length-9; k++ {
				for l := k + 1; l < length-8; l++ {
					for m := l + 1; m < length-7; m++ {
						for n := m + 1; n < length-6; n++ {
							for o := n + 1; o < length-5; o++ {
								for p := o + 1; p < length-4; p++ {
									for q := p + 1; q < length-3; q++ {
										for r := q + 1; r < length-2; r++ {
											for s := r + 1; s < length-1; s++ {
												for t := s + 1; t < length; t++ {
													var b int
													b += digit(bank, i, length, 12)
													b += digit(bank, j, length, 11)
													b += digit(bank, k, length, 10)
													b += digit(bank, l, length, 9)
													b += digit(bank, m, length, 8)
													b += digit(bank, n, length, 7)
													b += digit(bank, o, length, 6)
													b += digit(bank, p, length, 5)
													b += digit(bank, q, length, 4)
													b += digit(bank, r, length, 3)
													b += digit(bank, s, length, 2)
													b += digit(bank, t, length, 1)

													if b > jolts {
														jolts = b
													}
													// speed.messure()
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return
}

type speedemitor struct {
	start time.Time
	count float64
}

func newSpeedemitor() *speedemitor {
	return &speedemitor{
		start: time.Now(),
		count: 0,
	}
}

func (s *speedemitor) messure() {
	s.count++
	if math.Mod(s.count, 10000000) == 0 {
		rate := s.count / time.Since(s.start).Seconds()
		fmt.Printf("\rSpeed rate=%f/s count=%f", rate, s.count)
	}
}

func digit(bank int, index int, length int, position int) int {
	n := int(math.Pow10(length - index - 1))
	p := int(math.Pow10(position - 1))

	return ((bank / n) % 10) * p
}
