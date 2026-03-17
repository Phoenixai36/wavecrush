package ffmpeg

import (
	"fmt"
	"os/exec"
)

// Wrapper handles FFmpeg command execution
type Wrapper struct {
	BinaryPath string
}

// NewWrapper creates a new FFmpeg wrapper
func NewWrapper() (*Wrapper, error) {
	// Check if ffmpeg is available
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		return nil, fmt.Errorf("ffmpeg not found in PATH: %w", err)
	}

	return &Wrapper{
		BinaryPath: path,
	}, nil
}

// ConvertFormat converts audio from one format to another
func (w *Wrapper) ConvertFormat(input, output, format string, bitrate string) error {
	args := []string{
		"-i", input,
		"-b:a", bitrate,
		output,
	}

	cmd := exec.Command(w.BinaryPath, args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg conversion failed: %w", err)
	}

	return nil
}

// GetVersion returns the FFmpeg version
func (w *Wrapper) GetVersion() (string, error) {
	cmd := exec.Command(w.BinaryPath, "-version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get ffmpeg version: %w", err)
	}

	return string(output), nil
}

// TODO: Add more FFmpeg operations:
// - Batch processing
// - Audio normalization
// - Fade in/out effects
// - Metadata editing
// - Waveform generation
