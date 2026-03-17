# wavecrush

> 🌊 **Beautiful FFmpeg TUI in Go** - Built with Charm (Bubble Tea, Lip Gloss). Modern, fast, type-safe audio processing.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go 1.22](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org)
[![Built with Charm](https://img.shields.io/badge/Built%20with-Charm-pink.svg)](https://charm.sh)

## What is this?

`wavecrush` is a **terminal user interface (TUI)** for FFmpeg audio processing, built entirely in Go using the gorgeous [Charm libraries](https://charm.sh). Think of it as a modern, interactive FFmpeg wrapper that makes audio manipulation delightful.

### Why wavecrush?

- 🎨 **Beautiful TUI** - Powered by Bubble Tea & Lip Gloss
- ⚡ **Fast & Type-Safe** - Pure Go, compiled binary
- 🔧 **FFmpeg Under the Hood** - All the power, none of the complexity
- 🎵 **Audio-First** - Optimized for music producers & audio engineers
- 📦 **Zero Config** - Just run it

---

## Features

### Core Capabilities

- **Batch Processing** - Process entire directories with progress tracking
- **Format Conversion** - WAV, MP3, FLAC, AAC, OGG, M4A
- **Smart Compression** - Lossless & lossy with visual quality preview
- **Metadata Editing** - Edit tags, album art, track info
- **Audio Effects** - Normalize, fade in/out, EQ, compression
- **Waveform Visualization** - ASCII waveforms in your terminal

### UI/UX

- **Interactive File Browser** - Navigate your audio library
- **Real-time Progress** - See encoding stats live
- **Preset System** - Save your favorite settings
- **Keyboard Shortcuts** - Vim-like navigation
- **Theme Support** - Customizable color schemes

---

## Installation

### From Source

```bash
git clone https://github.com/Phoenixai36/wavecrush.git
cd wavecrush
go build -o wavecrush ./cmd/wavecrush
```

### Prerequisites

- **Go 1.22+**
- **FFmpeg** (must be in PATH)

---

## Usage

### Quick Start

```bash
# Launch TUI
./wavecrush

# Or process files directly
./wavecrush convert input.wav output.mp3 --bitrate 320k
```

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| `j/k` | Navigate up/down |
| `enter` | Select file/option |
| `space` | Toggle selection (batch mode) |
| `p` | Open preset menu |
| `q` | Quit |
| `?` | Help |

---

## Architecture

```
wavecrush/
├── cmd/
│   └── wavecrush/      # Main entry point
├── internal/
│   ├── tui/            # Bubble Tea UI components
│   ├── ffmpeg/         # FFmpeg wrapper & executor
│   ├── audio/          # Audio processing logic
│   └── config/         # Config & preset management
├── pkg/
│   └── models/         # Shared data structures
└── go.mod
```

### Tech Stack

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - TUI framework (Elm architecture)
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - Styling & layout
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - Pre-built TUI components
- **FFmpeg** - Audio encoding/decoding engine

---

## Development

### Project Status

🚧 **In Active Development** - Sprint-ready, modular design.

### Roadmap

#### Phase 1: Core MVP ✅
- [x] Repository setup
- [ ] FFmpeg wrapper
- [ ] Basic TUI shell (Bubble Tea)
- [ ] File browser
- [ ] Format conversion

#### Phase 2: Effects & Batch
- [ ] Batch processing queue
- [ ] Audio normalization
- [ ] Fade in/out effects
- [ ] Progress indicators

#### Phase 3: Advanced Features
- [ ] Waveform visualization
- [ ] Preset system
- [ ] Metadata editor
- [ ] Theme customization

#### Phase 4: Polish
- [ ] Error handling & recovery
- [ ] Unit tests
- [ ] CI/CD pipeline
- [ ] Binary releases

### Contributing

Contributions welcome! This is a learning project built with modern Go practices.

```bash
# Run tests
go test ./...

# Format code
go fmt ./...

# Lint
golangci-lint run
```

---

## License

MIT © [Phoenixai36](https://github.com/Phoenixai36)

---

## Acknowledgments

- **[Charm](https://charm.sh)** - For making terminal UIs gorgeous
- **[FFmpeg](https://ffmpeg.org)** - The Swiss Army knife of multimedia
- Inspired by modern TUI tools like [lazygit](https://github.com/jesseduffield/lazygit) & [k9s](https://github.com/derailed/k9s)

---

**Built with ❤️ in Barcelona by a music producer who got tired of typing FFmpeg commands.**
