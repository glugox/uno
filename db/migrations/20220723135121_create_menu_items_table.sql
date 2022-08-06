CREATE TABLE menu_items
	(
		`id`         INTEGER PRIMARY KEY AUTOINCREMENT,
		`menus_id`   INTEGER,
		`label`      TEXT NOT NULL,
		`path`       TEXT,
		`ordering`   INTEGER,
		`parent_id`  INTEGER,
		`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
		`updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP
	);