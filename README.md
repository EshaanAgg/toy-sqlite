[![progress-banner](https://backend.codecrafters.io/progress/sqlite/33452344-e83b-43b7-b9f4-50e42d459472)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is my Go solution to the
["Build Your Own SQLite" Challenge](https://codecrafters.io/challenges/sqlite) by the team at Codecrafters.

This is a toy SQLite implementation that supports basic SQL queries like `SELECT`. It's no where near a full implementation of SQLite, but I make a honest attempt and try to go beyond what the challenge asks for, in hopes of learning more about SQLite, databases and gaining more experience with Go.

The implementation relies on a fully self written SQL parser and query executor. It is a decent parser which can't handle subqueries or joins, but can parse and understand many queries which are beyond the scope of the challenge.

The current implementation supports:
- Parsing and reading the metadata section of a SQLite database file
- Dot commands like `.tables` and `.dbinfo`
- Some sample queries like the following:

```sql
-- COUNT queries
SELECT COUNT(*) FROM table_name;
SELECT COUNT(a, b, c) FROM table_name;
```

# Sample Database

To make it easy to test queries locally, the Codecrafters team has added a sample database in the root of this repository: `sample.db`.

This contains two tables: `apples` & `oranges`. 

You can explore this database by running queries against it like this:

```sh
$ sqlite3 sample.db "select id, name from apples"
1 | Granny Smith
2 | Fuji
3 | Honeycrisp
4 | Golden Delicious
```

There are two other databases that you can use:

1. `superheroes.db`:
   - This is a small version of the test database used in the table-scan stage.
   - It contains one table: `superheroes`.
   - It is ~1MB in size.
2. `companies.db`:
   - This is a small version of the test database used in the index-scan stage.
   - It contains one table: `companies`, and one index: `idx_companies_country`
   - It is ~7MB in size.

These aren't included in the repository because they're large in size. They can be download them by running this script:

```sh
./download_sample_databases.sh
```

If the script doesn't work for some reason, you can download the databases directly from [codecrafters-io/sample-sqlite-databases](https://github.com/codecrafters-io/sample-sqlite-databases).
