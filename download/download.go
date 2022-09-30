package download

import (
	"net/http"
	"os"
)

func DownloadHandler(mux *http.ServeMux) {
	mux.HandleFunc("/download/", downloadHandler)
	mux.HandleFunc("/download/stream/", streamDownloadHandler)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	f, ff, _ := setStreamFile(w, "./resources/file.txt")
	defer f.Close()
	http.ServeContent(w, r, ff.Name(), ff.ModTime(), f)
}

func streamDownloadHandler(w http.ResponseWriter, r *http.Request) {
	f, ff, _ := setStreamFile(w, "./resources/bigfile.txt")
	defer f.Close()
	http.ServeContent(w, r, ff.Name(), ff.ModTime(), f)
}

func setStreamFile(w http.ResponseWriter, fileLocation string) (*os.File, os.FileInfo, error) {
	f, _ := os.Open(fileLocation)
	ff, _ := f.Stat()
	
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=" + ff.Name())

	return f, ff, nil
}