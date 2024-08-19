# Snippetbox

A simple & blazingly fast CRUD app to share snippets made with Go; features include but are not limited to Authentication, Security best practices, and ease of use.

## Technologies

- Front-end: Server-side rendered HTML with templates using the standard package `html/template`
- Back-end: Built completely with Go with minimal reliance on third-party packages. A list of all the packages used could be found inside `go.mod` file at the root directory.

## Usage

To run this app locally, you must have [Go](https://go.dev/doc/install) installed. \
Then, run these commands in your terminal:

```
git clone https://github.com/DrMorax/Snippetbox
```

Then install the required dependencies:

```
cd snippetbox && go mod tidy
```
This project uses [MySQL](https://dev.mysql.com/downloads/installer) database so you have to install it and create your own user for that database which will be used to make data operations like inserting and deleting snippets.
These are the tables you need to create:
```
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE snippetbox;

CREATE TABLE sessions (
 token CHAR(43) PRIMARY KEY,
 data BLOB NOT NULL,
 expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

Then run the app on your local machine:
```
go run ./cmd/web -dsn="your/own/dsn"
```

After that, go to http://localhost:4000 Where the app should be running.
