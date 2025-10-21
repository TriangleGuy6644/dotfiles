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
	token := "sl.u.AGCX8mlk8j2ospZa0v6_dHRYG5kXGcK7coNhum0S6gwxxXLf6w5y2AhUhj5r_d6QQaGS-BdMY6BqjCuOBXDuv0xmhTXncBvJyF6Ds84yXXlc6_wK4sjJ2_e7Tt3fCF0XAZgyU69JU3If-OGOeDURTCJLNtbOwzcsNX9qUFcLsXwVOKFXbxrawWubk_v1w7d654nOP2zL77MaznMsoNjDusl8QcG1wqpyPb0fiG73OQoe3Beif-e6PQ5S9cJbilNM2Y6TIhTeEWft-doYPmEtzNMu-wDVTS2sBEY4N-2DH0dEdlQBeH2LvjAK0zghfJV5uDEDZM2ZemE5n3HkS4hdUti_wRTSDfFwZmL1_1MXR3rSQMCbxmZ6nn_ikFXYIrPTDb-kvrKMiqf1TbftGiHjQDOXi3g3EJ9R4Pa1UbyiPJDE3rm6oxfNzutxVIDyshptrzLySslS6PP2QhrYa-KEk8ArpAexP-Cu6famLjTnvrpgKd-MfeD9i51sQHlz9kvLKjufWEZWipUNj29RPkxG5UCwhTo5_ldyNEhyvz2R8kjQr1MpxSucIHAYUu7gbF8oFqM9KnLeKkJXQXZ6jED5tREOOTT5pmC4OnbPS3ktVBsdnLGgh2ZgS5g7RK_Ps_lUybQSxv1ZM_Gts6C3IOFjTx7zfUeRs1nZw782Fy7uRYhykthAysSVYpUX50K3z1EmMgxbznLMl7H-iMNkmf-rCg86AvaN7x0EIlPYoEO8yba1umAk58hJOyyH5DMqJcJ2a6m3G8BeD-ttXRlJXkTz_jgxyE1GbVIUiZQokPiYwPkmliCU817qebmIvRIra4otbAgXawOpoMcsluWcjvByfN0r9zijSH5pz8HBP5ZyAhjRq1aVlzglY5eEkxpZfBmnVPsRxxnc5pwAyxUttxNKDADaPkofdLxSEbMqveORLwYlCaMeWmAX7gzJqqfDCJhNRcTdst3QzpCc0ZgB64-uvlYIA0x5eTgvr96eaw3f6PsNR8pjhxPTdxvchXkaXkkehG8Tum7Myzznr3Xo3dhPzpCaNP588QlnO18ZvZgxyYJAVdJlRVTBO6XEYJ5FN47O3AyUE7nVmyyJBtXb2P54c4AD43tSHQYS0WCDNzPL8j1Kue8dz6IEoNn1Jif9l3sE2Vd6OBhtEoz9PDe8wZ_KS1OjntCdY33kEEQckU_rubq-ECP1-EiKD_erzUY5Z5OUWwimNintz7RWyGBRpJIFG7byp5kjVe08n-lMVdFM5w_W5WfiJ2gVY23ezn8NMRgx8VaOIsE7xJEOkJO7h-0bKB4TguNfmnwqKSi-2ubrSghT0jbtSa9lr8_ITujSaQfRelEWnZGykKdp8JoVWpUZJ81pP-RxsATT1oK4RMHuPpRVX4MF0DITCHsnTR6o_R1blU5DduN9-sbZhBjwOBiHw6cvmuWr1zWzFcnJpij-qf35PQ"
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
