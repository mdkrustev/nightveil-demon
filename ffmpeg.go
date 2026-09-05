package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func ffmpegPath() string {

	// DEV MODE
	if runtime.GOOS == "windows" {
		if _, err := os.Stat("./bin/ffmpeg.exe"); err == nil {
			return "./bin/ffmpeg.exe"
		}
	} else {
		if _, err := os.Stat("./bin/ffmpeg"); err == nil {
			return "./bin/ffmpeg"
		}
	}

	exe, err := os.Executable()
	if err != nil {
		return "ffmpeg"
	}

	appDir := filepath.Dir(exe)

	switch runtime.GOOS {

	case "windows":
		return filepath.Join(appDir, "ffmpeg.exe")

	case "darwin":
		return filepath.Join(appDir, "../Resources/ffmpeg")

	case "linux":
		return filepath.Join(appDir, "ffmpeg")

	default:
		return "ffmpeg"
	}
}

func renderHandler(w http.ResponseWriter, r *http.Request) {

	var req RenderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(
			w,
			"Invalid JSON: "+err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if len(req.Segments) == 0 {
		http.Error(
			w,
			"No segments provided",
			http.StatusBadRequest,
		)
		return
	}

	tmpDir, err := os.MkdirTemp(
		"",
		"nightveildemon-render-*",
	)

	if err != nil {
		http.Error(
			w,
			"Cannot create temp folder",
			500,
		)
		return
	}

	defer os.RemoveAll(tmpDir)

	var files []string

	for i, seg := range req.Segments {

		output := filepath.Join(
			tmpDir,
			fmt.Sprintf("segment_%d.mp4", i),
		)

		filter := fmt.Sprintf(
			"scale=%d:%d:force_original_aspect_ratio=increase,crop=%d:%d,setsar=1",
			req.Settings.Width,
			req.Settings.Height,
			req.Settings.Width,
			req.Settings.Height,
		)

		args := []string{
			"-y",

			"-loop",
			"1",

			"-i",
			seg.Image,

			"-i",
			seg.Audio,

			"-t",
			fmt.Sprintf("%f", seg.Duration),

			"-vf",
			filter,

			"-c:v",
			"libx264",

			"-preset",
			"fast",

			"-pix_fmt",
			"yuv420p",

			"-c:a",
			"aac",

			"-b:a",
			"192k",

			"-r",
			fmt.Sprintf("%d", req.Settings.Fps),

			output,
		}

		cmd := exec.Command(
			ffmpegPath(),
			args...,
		)

		out, err := cmd.CombinedOutput()

		if err != nil {

			json.NewEncoder(w).Encode(
				map[string]string{
					"status":  "error",
					"message": string(out),
				},
			)

			return
		}

		abs, _ := filepath.Abs(output)

		files = append(files, abs)
	}

	list := filepath.Join(
		tmpDir,
		"list.txt",
	)

	var builder strings.Builder

	for _, file := range files {

		safe := strings.ReplaceAll(
			file,
			"'",
			"'\\''",
		)

		builder.WriteString(
			fmt.Sprintf(
				"file '%s'\n",
				safe,
			),
		)
	}

	os.WriteFile(
		list,
		[]byte(builder.String()),
		0644,
	)

	cmd := exec.Command(
		ffmpegPath(),

		"-y",

		"-f",
		"concat",

		"-safe",
		"0",

		"-i",
		list,

		"-c",
		"copy",

		req.Output,
	)

	out, err := cmd.CombinedOutput()

	if err != nil {

		json.NewEncoder(w).Encode(
			map[string]string{
				"status":  "error",
				"message": string(out),
			},
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"status":  "success",
			"output":  req.Output,
			"message": "Video rendered successfully",
		},
	)
}
