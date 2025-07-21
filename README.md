# Blog Aggregator CLI

A command-line RSS (Really Simple Syndication) aggregator for managing and reading blog posts from multiple feeds. This tool allows users to register, log in, follow RSS feeds, and view aggregated content directly from the terminal.

## Project Goal

The goal of this project is to provide a simple, backend-powered CLI for users to consume blog content using RSS feeds. It combines authentication, feed management, and persistent storage with a clean and scriptable interface.

---

## Features

- **User Authentication:** Register, log in, and manage your account.
- **Feed Management:** Add, list, and browse RSS feeds.
- **Follow Blogs:** Follow/unfollow blogs and view followed feeds.
- **Aggregated Content:** Browse posts from all your followed feeds in one place.
- **Persistent Storage:** Uses PostgreSQL to store user, feed, and post data.
- **Scriptable Interface:** All features accessible from the command line for automation and workflows.

---

## Available CLI Commands

| Command                          | Description                                                             |
| -------------------------------- | ----------------------------------------------------------------------- |
| `gator login <username>`         | Log in as an existing user                                              |
| `gator register <username>`      | Register a new user                                                     |
| `gator reset`                    | Reset your account/password                                             |
| `gator users`                    | View all registered users                                               |
| `gator agg`                      | View aggregated posts from your feeds                                   |
| `gator addfeed "<name>" "<url>"` | Add a new RSS feed                                                      |
| `gator feeds`                    | List all available feeds                                                |
| `gator follow "<url>"`           | Follow a feed                                                           |
| `gator following`                | List feeds you are following                                            |
| `gator unfollow "<url>"`         | Unfollow a feed                                                         |
| `gator browse [number]`          | Browse posts from feeds you follow (optionally specify number of posts) |

---

## Requirements

Before using this project, ensure you have the following installed:

- [Go](https://go.dev/dl/) (v1.20 or newer)
- [PostgreSQL](https://www.postgresql.org/download/)

---

## Installation

To install the CLI tool, run:

```bash
go install github.com/manonmission88/BlogAggregator/cmd/gator@latest
```

This will download and install the `gator` CLI binary in your `$GOPATH/bin`.

---

## Getting Started

1. **Set up PostgreSQL:**  
   Make sure you have a running PostgreSQL instance.  
   Create a database for the aggregator.

   ```bash
   createdb blog_aggregator
   ```

2. **Configure Database Connection:**  
   Set your PostgreSQL connection string in the `.gatorconfig.json` file in your home or project directory:

   ```json
   {
     "db_url": "postgres://username:password@localhost:5432/blog_aggregator?sslmode=disable",
     "current_user": ""
   }
   ```

3. **Run Database Migrations:**
   Use [Goose](https://github.com/pressly/goose) to apply and rollback migrations:

   ```bash
   goose postgres "postgres://username:password@localhost:5432/blog_aggregator?sslmode=disable" up   # Apply migrations
   goose postgres "postgres://username:password@localhost:5432/blog_aggregator?sslmode=disable" down # Rollback last migration
   ```

4. **Generate Go Code from SQL Queries:**
   Use [sqlc](https://sqlc.dev/) to generate type-safe Go code from your SQL queries:

   ```bash
   sqlc generate
   ```

4. **Start Using the CLI:**

   ```bash
   gator reset                                   # Reset your account/data
   gator register manish                          # Register user 'kahya'
   gator addfeed "Hacker News RSS" "https://hnrss.org/newest"      # Add a feed
   gator register niure                        # Register user 'holgith'
   gator addfeed "Blog1" "https://www.wagslane.dev/index.xml" # Add a feed
   gator follow "https://hnrss.org/newest"       # Follow a feed
   gator following                               # List feeds you follow
   gator login manish                            # Log in as 'kahya'
   gator following                               # List feeds you follow
   gator browse [number]                         # Browse posts (specify number for limited posts)
   ```

---

## Example Usage

```bash
gator reset

gator register kahya

gator addfeed "Hacker News RSS" "https://hnrss.org/newest"

gator register holgith

gator addfeed "Lanes Blog" "https://www.wagslane.dev/index.xml"

gator follow "https://hnrss.org/newest"

gator following
# Output should include:
# Hacker News RSS
# Lanes Blog

gator login kahya
gator following

gator browse 5
# Browse the latest 5 posts from your followed feeds
```

---

## Future Updates

- **Web User Interface:**  
  The next planned major update is to migrate Blog Aggregator from a CLI tool to a modern web user interface (Web UI). This will allow users to interact with and manage their feeds via a browser, providing an improved experience and easier access.

---

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature-xyz`)
3. Make your changes and commit them
4. Push to your fork and open a Pull Request

---

## RESOURCES
- [boot.dev](https://boot.dev) — Go backend and database learning platform
- [Go by Example](https://gobyexample.com/) — Practical Go code examples
- [Go Documentation](https://golang.org/doc/) — Official Go docs
- [PostgreSQL Documentation](https://www.postgresql.org/docs/) — Official PostgreSQL docs
- [Goose](https://github.com/pressly/goose) — Database migration tool for Go
- [sqlc](https://sqlc.dev/) — Generate type-safe Go from SQL
- [RSS Specification](https://validator.w3.org/feed/docs/rss2.html) — RSS feed format reference
- [CLI Best Practices](https://clig.dev/) — Guidelines for building great command-line tools