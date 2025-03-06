# Lethe

**Lethe** is a command-line note-taking engine, written in Go 1.23 by Stephen Malone. It's designed to let you quickly access and safely store your notes with a simple CLI API, revision tracking and a non-proprietary database format.

```bash
$ lethe list
movies
project-ideas
todo-2024

$ lethe join movies "[ ] The Godfather"
$ lethe show movies
[x] Apocalypse Now
[x] Dog Day Afternoon
[ ] The Godfather

$ lethe hist movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
[2] 2025-03-06 08:52 - added 18 bytes
```

## Installation

You can install Lethe using your Go tools...

```
go install github.com/stvmln86/lethe@latest
```

...or download the [latest binary release][rel] for your platform.

## Configuration

Lethe stores all data in a single [SQLite][sql] database in one of three locations, depending on what environmental variables are set:

Variable          | Database Path
----------------- | -------------
`LETHE_DB`        | `$LETHE_DB`
`XDG_CONFIG_HOME` | `$XDG_CONFIG_HOME/lethe/lethe.db`
`HOME`            | `$HOME/.lethe`

To change the database location, set the appropriate variables and move your existing database file to the new path.

## Syntax

All note names in Lethe are lowercase and only allow alphanumeric characters, hyphens and underscores. Names are sanitised before use, so `"My Note 123!"` becomes `"my-note-123"`.

## Commands

**Search notes by name:**

```bash
$ lethe list
movies
project-ideas
todo-2024

$ lethe list 2024
todo-2024
```

**Search notes by text:**

```bash
$ lethe find "godfather"
movies
```

**Print an existing note:**

```bash
$ lethe show todo-2024
[ ] Take more notes.
```

**Edit a new or existing note:**

```bash
$ lethe edit project-ideas
```

**Append lines to a new or existing note:**

```bash
$ lethe join project-ideas "iphone app" "movie podcast"
$ lethe show project-ideas
iphone app
movie podcast
```

**Create a new empty note:**

```bash
$ lethe make new-note
```

**Delete an existing note:**

```
$ lethe wipe old-note
```

**View an existing note's history:**

```bash
$ lethe hist movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
[2] 2025-03-06 08:52 - added 18 bytes
```

**Revert an existing note to the previous version:**

```bash
$ lethe prev movies
$ lethe hist movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
```

## Contributing

Please submit all bug reports and feature requests to the issue tracker.

[rel]: https://github.com/stvmln86/lethe/releases/latest
[sql]: https://www.sqlite.org/
