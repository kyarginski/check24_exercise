CREATE TABLE blog_entries (
    entry_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    creation_date DATETIME NOT NULL,
    author VARCHAR(100) NOT NULL,
    text TEXT NOT NULL,
    image_link VARCHAR(255)
);
