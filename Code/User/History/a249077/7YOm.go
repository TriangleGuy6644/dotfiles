package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var hardcodedPaths = []string{
	"/home/you/Documents",
	"/home/you/.config",
	"C:\\Users\\You\\Desktop\\note.txt",
}

func main() {
	var zipNameFlag string
	var dropboxPathFlag string
	var chunkMB int
	flag.StringVar(&zipNameFlag, "zipname", "", "")
	flag.StringVar(&dropboxPathFlag, "dropboxpath", "", "")
	flag.IntVar(&chunkMB, "chunkmb", 8, "")
	flag.Parse()
	paths := hardcodedPaths
	tokenBytes, err := os.ReadFile("dropbox_token.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "missing dropbox_token.txt file with your token inside")
		os.Exit(1)
	}
	token := strings.TrimSpace(string(tokenBytes))
	now := time.Now().UTC()
	if zipNameFlag == "" {
		zipNameFlag = fmt.Sprintf("backup-%04d%02d%02d-%02d%02d%02d.zip", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	}
	if dropboxPathFlag == "" {
		dropboxPathFlag = "/" + filepath.Base(zipNameFlag)
	}
	tmpFile, err := os.CreateTemp("", "dbbak-*.zip")
	if err != nil {
		fmt.Fprintln(os.Stderr, "creating temp file:", err)
		os.Exit(1)
	}
	tmpPath := tmpFile.Name()
	tmpFile.Close()
	zipf, err := os.Create(tmpPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create zip file:", err)
		os.Exit(1)
	}
	zw := zip.NewWriter(zipf)
	for _, p := range paths {
		err = addPathToZip(zw, p)
		if err != nil {
			zw.Close()
			zipf.Close()
			os.Remove(tmpPath)
			fmt.Fprintln(os.Stderr, "zipping error:", err)
			os.Exit(1)
		}
	}
	err = zw.Close()
	if err != nil {
		zipf.Close()
		os.Remove(tmpPath)
		fmt.Fprintln(os.Stderr, "finalize zip:", err)
		os.Exit(1)
	}
	_, err = zipf.Seek(0, io.SeekStart)
	if err != nil {
		zipf.Close()
		os.Remove(tmpPath)
		fmt.Fprintln(os.Stderr, "seek zip:", err)
		os.Exit(1)
	}
	info, err := zipf.Stat()
	if err != n
