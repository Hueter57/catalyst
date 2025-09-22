-- Create "tasks" table
CREATE TABLE `tasks` (`id` uuid NOT NULL, `title` varchar(255) NOT NULL, `description` longtext NULL, `status` varchar(255) NOT NULL DEFAULT 'in_progress', `importance` varchar(255) NOT NULL DEFAULT 'low', `due_date` timestamp NULL, `message_id` varchar(255) NULL, `channel_id` varchar(255) NOT NULL, `created_at` timestamp NULL, `updated_at` timestamp NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "assignees" table
CREATE TABLE `assignees` (`id` uuid NOT NULL, `user_id` uuid NOT NULL, `task_assignee` uuid NULL, PRIMARY KEY (`id`), UNIQUE INDEX `task_assignee` (`task_assignee`), UNIQUE INDEX `user_id` (`user_id`), CONSTRAINT `assignees_tasks_assignee` FOREIGN KEY (`task_assignee`) REFERENCES `tasks` (`id`) ON UPDATE RESTRICT ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "reminders" table
CREATE TABLE `reminders` (`id` uuid NOT NULL, `remind_at` timestamp NULL, `task_reminder` uuid NULL, PRIMARY KEY (`id`), UNIQUE INDEX `task_reminder` (`task_reminder`), CONSTRAINT `reminders_tasks_reminder` FOREIGN KEY (`task_reminder`) REFERENCES `tasks` (`id`) ON UPDATE RESTRICT ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "categories" table
CREATE TABLE `categories` (`id` uuid NOT NULL, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `name` (`name`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "task_category" table
CREATE TABLE `task_category` (`task_id` uuid NOT NULL, `category_id` uuid NOT NULL, PRIMARY KEY (`task_id`, `category_id`), INDEX `task_category_category_id` (`category_id`), CONSTRAINT `task_category_category_id` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT `task_category_task_id` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_bin;
