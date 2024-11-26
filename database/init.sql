
CREATE TABLE IF NOT EXISTS tasks (
    taskId INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255),
    description TEXT,
    deadline INTEGER,  -- You can store timestamps
    exitCriteria TEXT,
    complete INTEGER,   -- Use INTEGER to represent BOOLEAN (0 for false, 1 for true)
    UNIQUE (title, deadline, description, exitCriteria)  -- Composite unique index
);
