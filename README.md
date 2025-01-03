# LeetCode Problem Organizer 🚀

This repository contains a script that helps you **organize** and **search** your LeetCode solutions. The script allows you to categorize problems by their **date**, **name**, and **problem number**, making it easier to track your progress and revisit past problems.

## Purpose 🎯

The goal of this script is to:

- **Organize** LeetCode problem solutions by date and problem number.
- **Easily search** for problems based on name, number, or the date you solved them.
- Help you keep your LeetCode journey **clean** and **organized**. 🧹✨

---

## Features 🛠️

- **Organize by Date and Problem Number**: Automatically organizes your solutions into folders based on the date and problem number.
- **Search Functionality**: Quickly find problems by name, number, or the date you solved them.
- **List All Problems**: View all problems you’ve solved, neatly organized.

---

## Requirements 📦

- **Node.js** or **Bun.js** installed on your system.

---

## Installation 📝

1. Clone the repository:

   ```bash
   git clone https://github.com/0xvish/leetcode.git
   cd leetcode
   ```

2. Install dependencies:

   ```bash
   bun install  # If using Bun.js
   # or
   npm install  # If using Node.js
   ```

---

## Usage 🔧

You can run the script via **Node.js** or **Bun.js**.

### Organizing Problems 🗂️

You can organize your problems by **date** and **name**. The script will create a folder for each problem based on the **date** and **problem number**.

```bash
# Usage:
# bun organize.js organize <date> <problem-statement>
# or
# node organize.js organize <date> <problem-statement>

bun organize.js organize "01-Jan-2024" "1422. Maximum Score After Splitting a String"
# or
node organize.js organize "01-Jan-2024" "1422. Maximum Score After Splitting a String"
```

- **date**: The date you solved the problem (e.g., `01-Jan-2024`).
- **problem-statement**: The full problem statement, including the problem number (e.g., `1422. Maximum Score After Splitting a String`).

---

### Searching Problems 🔍

You can search for problems using various criteria:

#### Search by Problem Name (Keyword):

Search for problems by a **keyword** in the problem statement.

```bash
# Usage:
# bun organize.js search-name <keyword>
# or
# node organize.js search-name <keyword>

bun organize.js search-name "String"
# or
node organize.js search-name "String"
```

#### Search by Date 📅:

Search for problems solved on a **specific date**.

```bash
# Usage:
# bun organize.js search-date <date>
# or
# node organize.js search-date <date>

bun organize.js search-date "01-Jan-2024"
# or
node organize.js search-date "01-Jan-2024"
```

#### Search by Problem Number #️⃣:

Search for problems by their **problem number** (e.g., `1422`).

```bash
# Usage:
# bun organize.js search-number <problem-number>
# or
# node organize.js search-number <problem-number>

bun organize.js search-number "1422"
# or
node organize.js search-number "1422"
```

#### Search by Keyword 🧐:

Search for problems that contain a **keyword** in the problem statement.

```bash
# Usage:
# bun organize.js search-keyword <keyword>
# or
# node organize.js search-keyword <keyword>

bun organize.js search-keyword "Split"
# or
node organize.js search-keyword "Split"
```

#### List All Problems 📜:

Display all **organized problems** in your directory.

```bash
# Usage:
# bun organize.js list
# or
# node organize.js list

bun organize.js list
# or
node organize.js list
```

---

## Example Commands 💡

Here are some example commands to showcase how you can use the script:

- **Organize a new problem**:

  ```bash
  bun organize.js organize "01-Jan-2024" "1422. Maximum Score After Splitting a String"
  ```

- **Search for a problem by name**:

  ```bash
  bun organize.js search-name "String"
  ```

- **Search for a problem by number**:

  ```bash
  bun organize.js search-number "1422"
  ```

- **List all organized problems**:
  ```bash
  bun organize.js list
  ```

---

## File Structure 🏗️

The LeetCode solutions will be organized in directories named as:

```
<date>_<problem-number>
  ├── main.go      (Your solution file)
  └── problems.json (Database of all organized problems)
```

Example:

```
01-Jan-2024_1422
  ├── main.go
  └── problems.json
```

The `problems.json` file contains metadata for all problems, making it easier to search for and organize your solutions. 🎉

---

## Contributing 🤝

Feel free to **fork** this repository and **create a pull request** if you'd like to improve it, add new features, or fix any bugs. All contributions are welcome!

---

## License 📄

This project is licensed under the **MIT License**.
