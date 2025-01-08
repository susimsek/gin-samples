-- Create greeting table if it does not exist
CREATE TABLE IF NOT EXISTS greeting (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Unique identifier for the greeting
    message TEXT NOT NULL, -- Greeting message
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP -- Last update timestamp
);
