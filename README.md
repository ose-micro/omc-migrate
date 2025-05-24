# 🐘 ose-migrate

>A simple and powerful CLI tool for managing PostgreSQL schema migrations with built-in support for logging, and retry middleware via the `ose-micro/postgres` package.

---

## ✨ Features

- 📦 Apply migrations (`up`, `down`)
- 🛠 Create timestamped migration files
- 🔄 Middleware support (logging, retry)
- 🧪 Clean CLI experience powered by Cobra
- 🐚 Shell completion (bash, zsh, fish)

---

## 📥 Installation

### Clone and build

```bash
git clone https://github.com/ose-micro/migrate.git
cd migrate
make build
