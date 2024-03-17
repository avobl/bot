CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT PRIMARY KEY NOT NULL, -- Telegram user id
    chat_id BIGINT NOT NULL,
    username VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tokens (
    user_id BIGINT PRIMARY KEY NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE schedules (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   user_id INTEGER,
   schedule_type VARCHAR(20), -- 'once', 'daily', 'weekly', 'custom', etc.
   cron_expression TEXT, -- The cron expression or a custom representation
   start_date TIMESTAMP,
   end_date TIMESTAMP,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
   is_active BOOLEAN NOT NULL DEFAULT TRUE,

   FOREIGN KEY (user_id) REFERENCES users(user_id)
);