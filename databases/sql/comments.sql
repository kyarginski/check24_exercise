CREATE TABLE comments (
    comment_id INT AUTO_INCREMENT PRIMARY KEY,
    entry_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255),
    url VARCHAR(255),
    comment_text TEXT NOT NULL,
    comment_date DATETIME NOT NULL,
    FOREIGN KEY (entry_id) REFERENCES blog_entries(entry_id) ON DELETE CASCADE
);
