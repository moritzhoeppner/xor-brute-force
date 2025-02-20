package main

import (
	"crypto/subtle"
	"encoding/json"
	"flag"
	"fmt"
	"path"
	"os"
	"sort"
	"github.com/moritzhoeppner/xor-brute-force/internal/mbxor"
)

type SrcFile struct {
	Name    string
	Content []byte
}

func processCmdArgs() (map[byte]float64, []SrcFile, error) {
	distFilename := flag.String("dist", "",
		"path of a JSON file that contains the expected byte distribution")
	msgDir := flag.String("messages", ".",
		"directory that contatins the encrypted messages")

	flag.Parse()

	var expectedDist map[byte]float64
	var srcFiles     []SrcFile
	
	// Read and parse dist argument.
	distFile, err := os.ReadFile(*distFilename)
	if err != nil {
		return expectedDist, srcFiles, err
	}
	json.Unmarshal(distFile, &expectedDist)

	// Read all files in messages directory.
	msgDirEntries, err := os.ReadDir(*msgDir)
	if err != nil {
		return expectedDist, srcFiles, err
	}
	srcFiles = make([]SrcFile, len(msgDirEntries))
	for i, e := range msgDirEntries {
		content, err := os.ReadFile(path.Join(*msgDir, e.Name()))
		if err != nil {
			return expectedDist, srcFiles, err
		}
		srcFiles[i] = SrcFile{ Name: e.Name(), Content: content }
	}

	return expectedDist, srcFiles, nil
}

func xorSrcFiles(srcFiles []SrcFile) []SrcFile {
	// Find length of the source file.
	sort.Slice(srcFiles, func (i, j int) bool {
		return len(srcFiles[i].Content) < len(srcFiles[j].Content)
	})
	minLen := len(srcFiles[0].Content)
	
	// Truncate source files to minLen and XOR the first with every other one.
	for i := 1; i < len(srcFiles); i++ {
		srcFiles[i].Content = srcFiles[i].Content[:minLen]
		subtle.XORBytes(srcFiles[i].Content, srcFiles[0].Content, srcFiles[i].Content[:minLen])
	}

	return srcFiles
}

func main() {
	expectedDist, srcFiles, err := processCmdArgs()
	if err != nil {
		panic(err)
	}

	// XOR the given source files.
	xoredSrcFiles := xorSrcFiles(srcFiles)
	messages := make([][]byte, len(xoredSrcFiles) - 1)
	for i, _ := range messages {
		messages[i] = xoredSrcFiles[i + 1].Content
	}

	// I'll assume that each plaintext consists only of bytes that appear in the given distribution.
	expectedBytes := make([]byte, len(expectedDist))
	i := 0
	for k, _ := range expectedDist {
		expectedBytes[i] = k
		i++
	}
	
	// Get all possible keys, then determine which is the most likely one.

	xor := mbxor.Mbxor{
		Ciphertexts: messages,
		KeyBytes: expectedBytes,
		ResultBytes: expectedBytes,
	}

	candidates, err := xor.Candidates()
	if err != nil {
		panic(err)
	}

	key := candidates.MostLikely(expectedDist)

	// And output.
	for i, xoredSrcFile := range xoredSrcFiles {
		var decrypted []byte
		if i == 0 {
			decrypted = key
		} else {
			decrypted = make([]byte, len(key))
			subtle.XORBytes(decrypted, key, xoredSrcFile.Content)
		}	
		fmt.Printf("%s: %s\n", xoredSrcFile.Name, decrypted)
	}
}
