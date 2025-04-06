# gator

Gator is a command-line tool for managing RSS feeds. It allows users to register, log in, follow feeds, and browse the latest posts from their followed feeds. Gator is designed to simplify feed aggregation and provide a seamless experience for staying updated with your favorite content.

## Installation

1. Clone the repository:
  ```bash
  git clone https://github.com/yourusername/gator.git
  ```

2. Navigate to the project directory:
  ```bash
  cd gator
  ```

3. Build the project:
  ```bash
  go build
  ```

4. Run the executable:
  ```bash
  ./gator
  ```

Alternatively, you can install it globally:
  ```bash
  go install github.com/yourusername/gator@latest
  ```

Make sure you have [Go](https://golang.org/dl/) installed on your system before proceeding with the installation.

## Commands

### `login`
Logs in as an existing user.

**Usage:**
```bash
gator login <name>
```

**Example:**
```bash
gator login alice
```

---

### `register`
Registers a new user.

**Usage:**
```bash
gator register <name>
```

**Example:**
```bash
gator register alice
```

---

### `reset`
Resets the database by deleting all users.

**Usage:**
```bash
gator reset
```

**Example:**
```bash
gator reset
```

---

### `users`
Lists all users in the system. The current user is marked.

**Usage:**
```bash
gator users
```

**Example:**
```bash
gator users
```

---

### `agg`
Starts aggregating feeds at a specified interval.

**Usage:**
```bash
gator agg <time_between_reqs>
```

**Example:**
```bash
gator agg 10s
```

---

### `addfeed`
Adds a new feed and follows it.

**Usage:**
```bash
gator addfeed <feed_name> <url>
```

**Example:**
```bash
gator addfeed "Tech News" https://example.com/rss
```

---

### `feeds`
Lists all available feeds.

**Usage:**
```bash
gator feeds
```

**Example:**
```bash
gator feeds
```

---

### `follow`
Follows an existing feed.

**Usage:**
```bash
gator follow <url>
```

**Example:**
```bash
gator follow https://example.com/rss
```

---

### `following`
Lists all feeds the current user is following.

**Usage:**
```bash
gator following
```

**Example:**
```bash
gator following
```

---

### `unfollow`
Unfollows a feed.

**Usage:**
```bash
gator unfollow <url>
```

**Example:**
```bash
gator unfollow https://example.com/rss
```

---

### `browse`
Displays latest posts from feeds the user is following.

**Usage:**
```bash
gator browse [posts_number]
```

**Example:**
```bash
gator browse 5
```