package gitwalk

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

var blobQueue map[string]struct{} = make(map[string]struct{})

type ChangeType byte

const (
	Added       ChangeType = 'A'
	Copied      ChangeType = 'C'
	Deleted     ChangeType = 'D'
	Modified    ChangeType = 'M'
	Renamed     ChangeType = 'R'
	TypeChanged ChangeType = 'T'
	Unmerged    ChangeType = 'U'
	Unknown     ChangeType = 'X'
)

type DiffTreeEntry struct {
	OldMode string
	NewMode string
	OldBlob string
	NewBlob string
	Status  ChangeType
	OldPath string
	NewPath string
}

func newDiffTreeEntry(metadata []string, status ChangeType) DiffTreeEntry {
	entry := DiffTreeEntry{
		OldMode: metadata[0],
		NewMode: metadata[1],
		OldBlob: metadata[2],
		NewBlob: metadata[3],
		Status: status,
		OldPath: metadata[5],
	}

	if entry.Status == Renamed {
		entry.NewPath = metadata[6]
	}

	return entry
}

func diffTree(commitHash string) {
	cmd := exec.Command("git", "diff-tree", "-r", "--no-commit-id", "--root", "-M", commitHash)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		metadata := strings.Fields(scanner.Text())
		status := ChangeType(metadata[4][0])

		if ChangeType(status) == Deleted {
			continue
		}

		entry := newDiffTreeEntry(metadata, status)
		fmt.Println("Blob : " + entry.NewBlob)
		addBlobToQueue(entry.NewBlob)
		for blobHash := range blobQueue {
			readBlobContent(blobHash)
		}
		fmt.Println("-------------------------------")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

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