CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE qrcode (
    id VARCHAR(16) PRIMARY KEY,
    user_id INTEGER,
    link VARCHAR(255),
    created_at DATETIME,
    FOREIGN KEY(user_id) REFERENCES user(id)
);
