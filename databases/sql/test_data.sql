INSERT INTO blog_entries (title, creation_date, author, text, image_link)
VALUES
    ('Sample Blog Post 1', NOW(), 'John Doe', 'This is the content of the first blog post.', 'image1.jpg'),
    ('Sample Blog Post 2', NOW(), 'Jane Smith', 'Here is the content of the second blog post.', NULL),
    ('Sample Blog Post 3', NOW(), 'Alice Johnson', 'This is the third blog post without an image.', NULL),
    ('Sample Blog Post 3', NOW(), 'Victor Kyarginskiy', 'This is my dummy post.', NULL);

INSERT INTO comments (entry_id, name, email, url, comment_text, comment_date)
VALUES
    (1, 'Commenter 1', 'commenter1@example.com', 'http://commenter1.com', 'This is the first comment on the first blog post.', NOW()),
    (1, 'Commenter 2', 'commenter2@example.com', NULL, 'This is the second comment on the first blog post.', NOW()),
    (2, 'Commenter 3', 'commenter3@example.com', 'http://commenter3.com', 'A comment on the second blog post.', NOW());
