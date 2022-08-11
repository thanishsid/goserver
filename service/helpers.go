package service

import (
	"encoding/base64"
	"encoding/json"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/thanishsid/goserver/config"
)

// Get file path from image directory in config and the image filename.
func getImagePath(filename string) string {
	return filepath.Join(config.C.ImageDirectory, filename)
}

// Get file path from video directory in config and the video filename.
func getVideoPath(filename string) string {
	return filepath.Join(config.C.VideoDirectory, filename)
}

// Get Video Duration.
func getVideoDuration(path string) (float64, error) {
	ffmpegProber := exec.Command(
		"ffprobe",
		"-v", "error",
		"-show_entries",
		"format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		path,
	)

	data, err := ffmpegProber.Output()
	if err != nil {
		return 0, err
	}

	durationString := strings.TrimSuffix(string(data), "\n")

	duration, err := strconv.ParseFloat(durationString, 64)
	if err != nil {
		return 0, err
	}

	return duration, nil
}

// Encode cursor to json and url safe base64.
func encodeCursor(obj any) (string, error) {
	jsn, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(jsn), nil
}

// Decode cursor from url safe bas64 and json to target type.
func decodeCursor[T any](cursor string) (T, error) {
	var obj T

	jsn, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		return obj, err
	}

	if err := json.Unmarshal(jsn, &obj); err != nil {
		return obj, err
	}

	return obj, nil
}
