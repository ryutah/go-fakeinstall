package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	minFileSize          = 200
	maxFileSize          = 3000
	maxBps               = 200
	maxDownloadDuration  = 500
	maxProgressbarLength = 30
)

var filenames = []string{
	"afjlkjers",
	"jitajlkaer",
	"jitewoalj",
	"iewitue",
}

const (
	maxMajorVersion = 20
	maxMinorVersion = 20
	maxPatchVersion = 20
)

func main() {
	for {
		totalsize := Download(RandomFilename())
		Install(totalsize)
	}
}

func Download(filename string) int {
	totalsize, progress := RandomeFileSize(), 0
	fmt.Printf("Downloading... %s(%dKB)\n", RandomFilename(), totalsize)
	for {
		progress = downloadProgress(totalsize, progress)
		if progress >= totalsize {
			break
		}
		slpTime := time.Duration(rand.Int63n(maxDownloadDuration) * int64(time.Millisecond))
		time.Sleep(slpTime)
	}
	fmt.Println()
	return totalsize
}

func downloadProgress(totalsize, progress int) int {
	bps := rand.Intn(maxBps)
	progress += bps
	if progress >= totalsize {
		progress = totalsize
	}
	printProgress(totalsize, progress, bps)
	return progress
}

func printProgress(totalsize, progress, bps int) {
	rate := (float64(progress) / float64(totalsize))
	progBarLen := int(maxProgressbarLength * rate)
	progBar := progressBar("#", progBarLen, maxProgressbarLength)
	dlRates := fmt.Sprintf("%dKB (%dKB/s)", progress, bps)
	fmt.Printf("\r %3d%% _progress... %-30s| %-20s", int(rate*100), progBar, dlRates)
}

func progressBar(char string, repeatCount, maxCount int) string {
	left := maxCount - repeatCount
	return strings.Repeat(char, repeatCount) + strings.Repeat(" ", left)
}

func Install(totalsize int) {
	wg := new(sync.WaitGroup)
	installEnd := false

	printProgress := func(w *sync.WaitGroup) {
		w.Add(1)
		for i := 0; ; i++ {
			fmt.Printf("\r Installing... %s", circle(i))
			slpTime := time.Duration(rand.Int63n(200) * int64(time.Millisecond))
			time.Sleep(slpTime)
			if installEnd {
				break
			}
		}
		w.Done()
	}
	go printProgress(wg)

	installTime := time.Duration(float64(totalsize) * 1.5 * float64(time.Millisecond))
	time.Sleep(installTime)
	installEnd = true
	wg.Wait()
	fmt.Println()
}

func RandomeFileSize() int {
	return rand.Intn(maxFileSize) + minFileSize
}

func RandomFilename() string {
	index := rand.Intn(len(filenames) - 1)
	return fmt.Sprintf("%s_%s.tar.gz", filenames[index], randomVersion())
}

func randomVersion() string {
	mejar := rand.Intn(maxMajorVersion)
	miner := rand.Intn(maxMinorVersion)
	patch := rand.Intn(maxPatchVersion)
	return fmt.Sprintf("%d.%d.%d", mejar, miner, patch)
}

var circles = []string{"-", "\\", "|", "/"}

func circle(count int) string {
	index := count % 3
	return circles[index]
}
