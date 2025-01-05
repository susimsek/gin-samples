CREATE TABLE greeting (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         message TEXT NOT NULL,
                         created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                         updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
