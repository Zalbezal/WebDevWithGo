# Create users table

```sql
DROP TABLE IF EXISTS users;
```

```sql
CREATE TABLE USERS (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
```