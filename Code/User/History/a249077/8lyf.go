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

func main() {
	var zipNameFlag string
	var dropboxPathFlag string
	var chunkMB int
	flag.StringVar(&zipNameFlag, "zipname", "", "")
	flag.StringVar(&dropboxPathFlag, "dropboxpath", "", "")
	flag.IntVar(&chunkMB, "chunkmb", 8, "")
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		fmt.Fprintln(os.Stderr, "no paths provided")
		os.Exit(2)
	}
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
	if err != nil {
		zipf.Close()
		os.Remove(tmpPath)
		fmt.Fprintln(os.Stderr, "stat zip:", err)
		os.Exit(1)
	}
	size := info.Size()
	fmt.Println("created zip:", tmpPath, "size bytes:", size)
	chunkSize := int64(chunkMB) * 1024 * 1024
	if size <= 150*1024*1024 {
		err = uploadSimple(zipf, token, dropboxPathFlag)
		if err != nil {
			zipf.Close()
			os.Remove(tmpPath)
			fmt.Fprintln(os.Stderr, "upload error:", err)
			os.Exit(1)
		}
	} else {
		err = uploadSession(zipf, token, dropboxPathFlag, chunkSize)
		if err != nil {
			zipf.Close()
			os.Remove(tmpPath)
			fmt.Fprintln(os.Stderr, "upload session error:", err)
			os.Exit(1)
		}
	}
	zipf.Close()
	outLocal := zipNameFlag
	err = os.Rename(tmpPath, outLocal)
	if err != nil {
		fmt.Println("warning: could not move temp zip to", outLocal, "keeping temp file at", tmpPath)
	} else {
		fmt.Println("saved local zip as", outLocal)
	}
	fmt.Println("upload complete to Dropbox path", dropboxPathFlag)
}

func addPathToZip(zw *zip.Writer, root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return addFileToZip(zw, root, filepath.Base(root))
	}
	base := filepath.Clean(root)
	return filepath.Walk(base, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(filepath.Dir(base), path)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		return addFileToZip(zw, path, rel)
	})
}

func addFileToZip(zw *zip.Writer, realPath, relPath string) error {
	f, err := os.Open(realPath)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := zw.Create(filepath.ToSlash(relPath))
	if err != nil {
		return err
	}
	_, err = io.Copy(w, f)
	return err
}

func uploadSimple(r io.ReaderAt, token, dropboxPath string) error {
	size, err := sizeOfReaderAt(r)
	if err != nil {
		return err
	}
	url := "https://content.dropboxapi.com/2/files/upload"
	req, err := http.NewRequest("POST", url, io.NewSectionReader(r, 0, size))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/octet-stream")
	arg := map[string]interface{}{"path": dropboxPath, "mode": "overwrite", "autorename": false, "mute": false}
	b, _ := json.Marshal(arg)
	req.Header.Set("Dropbox-API-Arg", string(b))
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("dropbox upload failed %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	return nil
}

func uploadSession(r io.ReaderAt, token, dropboxPath string, chunkSize int64) error {
	size, err := sizeOfReaderAt(r)
	if err != nil {
		return err
	}
	firstSize := chunkSize
	if firstSize > size {
		firstSize = size
	}
	firstBuf := make([]byte, firstSize)
	_, err = r.ReadAt(firstBuf, 0)
	if err != nil && err != io.EOF {
		return err
	}
	startURL := "https://content.dropboxapi.com/2/files/upload_session/start"
	req, err := http.NewRequest("POST", startURL, bytes.NewReader(firstBuf))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Dropbox-API-Arg", `{"close": false}`)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("start session failed %d: %s", resp.StatusCode, string(body))
	}
	var startResp struct {
		SessionId string `json:"session_id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&startResp)
	if err != nil {
		return err
	}
	sessionId := startResp.SessionId
	offset := int64(len(firstBuf))
	for offset < size {
		nextSize := chunkSize
		remaining := size - offset
		if remaining < nextSize {
			nextSize = remaining
		}
		buf := make([]byte, nextSize)
		_, err := r.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			return err
		}
		if offset+nextSize < size {
			appendURL := "https://content.dropboxapi.com/2/files/upload_session/append_v2"
			cursor := map[string]interface{}{"cursor": map[string]interface{}{"session_id": sessionId, "offset": offset}, "close": false}
			cj, _ := json.Marshal(cursor)
			req, err := http.NewRequest("POST", appendURL, bytes.NewReader(buf))
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-Type", "application/octet-stream")
			req.Header.Set("Dropbox-API-Arg", string(cj))
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			resp.Body.Close()
			if resp.StatusCode >= 300 {
				body, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("append failed %d: %s", resp.StatusCode, string(body))
			}
		} else {
			finishURL := "https://content.dropboxapi.com/2/files/upload_session/finish"
			cursor := map[string]interface{}{"cursor": map[string]interface{}{"session_id": sessionId, "offset": offset}, "commit": map[string]interface{}{"path": dropboxPath, "mode": "overwrite", "autorename": false, "mute": false}}
			cj, _ := json.Marshal(cursor)
			req, err := http.NewRequest("POST", finishURL, bytes.NewReader(buf))
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-Type", "application/octet-stream")
			req.Header.Set("Dropbox-API-Arg", string(cj))
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode >= 300 {
				body, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("finish failed %d: %s", resp.StatusCode, string(body))
			}
		}
		offset += nextSize
	}
	return nil
}

func sizeOfReaderAt(r io.ReaderAt) (int64, error) {
	switch v := r.(type) {
	case *os.File:
		info, err := v.Stat()
		if err != nil {
			return 0, err
		}
		return info.Size(), nil
	default:
		return 0, fmt.Errorf("cannot determine size")
	}
}
