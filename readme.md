# Lethe

**Lethe** is a command-line note-taking engine, written in Go 1.23 by Stephen Malone. It's designed to let you quickly access and safely store your notes with a simple CLI API, revision tracking, import/export features and a non-proprietary database format.

<details> <summary>Demo:</summary>

```
$ lethe
movies
project-ideas
todo-2024

$ lethe add movies "[ ] The Godfather"
$ lethe log movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
[2] 2025-03-06 08:52 - added 18 bytes 

$ lethe cat movies
[x] Apocalypse Now
[x] Dog Day Afternoon
[ ] The Godfather
```

</details>

## Installation

You can install Lethe using your Go tools...

```
go get github.com/stvmln86/lethe@latest
```

or download the latest binary for your platform.

## Configuration

`TODO: $HOME vs $XDG vs custom database path.`

## Commands

`TODO: Syntax explainer, abbreviations, basic commands.`

## Contributing

Please submit all bug reports and feature requests to the issue tracker, thank you.
