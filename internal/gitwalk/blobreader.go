package gitwalk

import (
	"bufio"
	"fmt"
	"os/exec"

	//"specter/internal/detector"
	"specter/internal/entropy"
)

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
		//detector.ApplyRules(scanner.Text())
		result := entropy.ComputeEntropyPerToken(scanner.Text())
		for _, te := range result {
			fmt.Println(te.Token, te.Entropy)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
