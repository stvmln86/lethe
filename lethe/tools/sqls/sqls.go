// Package sqls implements SQLite pragma and schema constants.
package sqls

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma foreign_keys = ON;
	pragma encoding = 'UTF-8';
`

// Schema is the default first-time-run database schema.
const Schema = `
	create table Notes (
		n_id integer primary key asc,
		init integer not null default (unixepoch()),
		name text    not null unique
	);

	create table Pages (
		p_id integer primary key asc,
		init integer not null default (unixepoch()),
		note integer not null references Notes(n_id),
		body text    not null
	);
`
