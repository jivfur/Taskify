
CREATE TABLE IF NOT EXISTS tasks (
    taskId INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    deadline INTEGER,  -- You can store timestamps
    exitCriteria TEXT,
    complete BOOLEAN
);

