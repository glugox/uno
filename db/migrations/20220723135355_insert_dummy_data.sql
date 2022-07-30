/* Create in-memory temp table for variables */
BEGIN;

/* Variables **/
CREATE TEMP TABLE IF NOT EXISTS Variables (Name TEXT PRIMARY KEY, Value TEXT);
INSERT OR REPLACE INTO Variables VALUES ('AdminMenuId', hex(randomblob(16)));

/* Menu */
INSERT INTO `menus` (`id`, `label`) VALUES ((SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), 'Admin');

/* Menu Items */
INSERT INTO `menu_items` (`id`, `menu_id`, `label`, `path`, `ordering`) VALUES (hex(randomblob(16)), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), 'Home', 'home', 1);
INSERT INTO `menu_items` (`id`, `menu_id`, `label`, `path`, `ordering`) VALUES (hex(randomblob(16)), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), 'Users', 'users', 1);

/* - Contacts */
INSERT OR REPLACE INTO Variables VALUES ('ContactsMenuItemId', hex(randomblob(16)));
INSERT INTO `menu_items` (`id`, `menu_id`, `label`, `path`, `ordering`) VALUES ((SELECT Value FROM Variables WHERE Name = 'ContactsMenuItemId'), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), 'Contacts', 'contacts', 1);
INSERT INTO `menu_items` (`id`, `menu_id`, `parent_id`, `label`, `path`, `ordering`) VALUES (hex(randomblob(16)), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), (SELECT Value FROM Variables WHERE Name = 'ContactsMenuItemId'), 'View All', 'view-all', 1);
INSERT INTO `menu_items` (`id`, `menu_id`, `parent_id`, `label`, `path`, `ordering`) VALUES (hex(randomblob(16)), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), (SELECT Value FROM Variables WHERE Name = 'ContactsMenuItemId'), 'Add New', 'add-new', 1);

INSERT INTO `menu_items` (`id`, `menu_id`, `label`, `path`, `ordering`) VALUES (hex(randomblob(16)), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), 'Accounts', 'accounts', 1);
INSERT INTO `menu_items` (`id`, `menu_id`, `label`, `path`, `ordering`) VALUES (hex(randomblob(16)), (SELECT Value FROM Variables WHERE Name = 'AdminMenuId'), 'Settings', 'settings', 1);

END;

