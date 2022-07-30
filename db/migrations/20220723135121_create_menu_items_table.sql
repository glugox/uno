CREATE TABLE menu_items
	(
		`id`         TEXT PRIMARY KEY,
		`menu_id`    TEXT,
		`label`      TEXT NOT NULL,
		`path`       TEXT,
		`ordering`   INTEGER,
		`parent_id`  INTEGER,
		`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
		`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP
	);