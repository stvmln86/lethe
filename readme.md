# Lethe

**Lethe** is a command-line note-taking engine, written in Go 1.23 by Stephen Malone. It's designed to let you quickly access and safely store your notes with a simple CLI API, revision tracking and a non-proprietary database format.

```
$ lethe list
movies
project-ideas
todo-2024

$ lethe join movies "[ ] The Godfather"
$ lethe hist movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
[2] 2025-03-06 08:52 - added 18 bytes

$ lethe show movies
[x] Apocalypse Now
[x] Dog Day Afternoon
[ ] The Godfather
```

## Installation

You can install Lethe using your Go tools...

```
go install github.com/stvmln86/lethe@latest
```

or download the [latest binary release][rel] for your platform.

## Configuration

Lethe stores all data in a single [SQLite][sql] database in one of three locations, depending on what environmental variables are set:

Variable          | Database Path
----------------- | -------------
`LETHE_DB`        | `$LETHE_DB`
`XDG_CONFIG_HOME` | `$XDG_CONFIG_HOME/lethe/lethe.db`
`HOME`            | `$HOME/.lethe`

You can change locations by changing your environment variables and moving your existing database file to the new location.

## Syntax

All note names in Lethe are lowercase and only allow alphanumeric characters, hyphens and underscores. Names are sanitised before use, so attempting to create a note called `My Note 123!` will actually create `my-note-123`.

## Commands

### `list [TEXT]`

Print the names of all existing notes, or notes with names containing `TEXT`. The search text is sanitised as per the above naming syntax.

```
$ lethe list
movies
project-ideas
todo-2024

$ lethe list TODO
todo-2024
```

### `find TEXT`

Print the names of all existing notes containing `TEXT`. The search text is case-insensitive.

```
$ lethe find "godfather"
movies
```

### `show NOTE`

Print the body of an existing note.

```
$ lethe show todo-2024
[ ] Take more notes.
```

### `edit NOTE`

Open a new or existing note in a temporary file in the default editor (according to `$EDITOR` or `$VISUAL`) and save the resulting changes.

```
$ lethe edit project-ideas
```

### `join NOTE LINES...`

Append one or more lines of text to a new or existing note.

```
$ lethe join project-ideas "iphone app" "movie podcast"
$ lethe show project-ideas
iphone app
movie podcast
```

### `make NOTE`

Create a new empty note.

```
$ lethe make new-note
```

### `wipe NOTE`

Delete an existing note and all its previous revisions.

```
$ lethe wipe old-note
```

### `hist NOTE`

Print a list of an existing note's revisions.

```
$ lethe hist movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
[2] 2025-03-06 08:52 - added 18 bytes
```

### `prev NOTE`

Revert an existing note to its previous revision.

```
$ lethe prev movies
$ lethe hist movies
[0] 2025-03-04 10:11 - initial
[1] 2025-03-05 22:37 - changed, removed 21 bytes
```

## Contributing

Please submit all bug reports and feature requests to the issue tracker, thank you.

[rel]: https://github.com/stvmln86/lethe/releases/latest
[sql]: https://www.sqlite.org/
