package gitwalk

import (
	"bufio"
	"fmt"
	"os/exec"
)

var blobQueue map[string]struct{} = make(map[string]struct{})

func addBlobToQueue(blobHash string) {
	blobQueue[blobHash] = struct{}{}
}

func readBlobContent(blobHash string) {
	cmd := exec.Command("git", "cat-file", "-p", blobHash)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		blobContent := scanner.Text()
		fmt.Println(blobContent)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
