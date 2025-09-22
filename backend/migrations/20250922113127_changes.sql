-- Modify "tasks" table
ALTER TABLE `tasks` MODIFY COLUMN `status` enum('in_progress','completed','waiting') NOT NULL DEFAULT 'in_progress', MODIFY COLUMN `importance` enum('high','medium','low') NOT NULL DEFAULT 'low';
