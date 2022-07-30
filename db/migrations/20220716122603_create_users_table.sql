CREATE TABLE IF NOT EXISTS users
	(
		`id`         TEXT PRIMARY KEY,
		`email`      VARCHAR(255) NOT NULL,
		`password`   VARCHAR(255) NOT NULL,
		`is_active`  BOOLEAN DEFAULT true,
		`created_at` TIMESTAMP    null,
		`updated_at` TIMESTAMP    null
	)