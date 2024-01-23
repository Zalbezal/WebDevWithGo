Postgres Constraints

Example SQL:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
```
 
Constraints are rules that we can apply to columns in our table. For example, we might want to ensure that every user in our database has a unique id, so we could use the UNIQUE constraint. We could also verify that an integer column is above or below a specific value.

Like data types, there are many constraints available in Postgres. For now we will only be using a few:

Constraint Description
UNIQUE This ensures that every record in your database has a unique value for the field that is set to unique. For example, you might want every user to have a unique email address. It is important to note that Postgres is case sensitive, so jon@CALHOUN.io is not the same as jon@calhoun.io. You will need to account for this on your own when writing data to your database.
NOT NULL This ensure that every record in your database has a value for this field. When you don’t provide a value for a field the database will traditionally store null, but this prevents that from being valid.
PRIMARY KEY This constraint is similar to combining both UNIQUE and NOT NULL but it can only be used once on each table, and it will automatically result in a the creation of an index for this field. The index is used to make it faster to look up records by this field. For example, we typically set the id to be the primary key, and then we look up users by their id when we need to fetch them from the database.
While it is common to validate data inside our own Go code, many constraints need to also be set at the database layer to ensure they are always enforced. For instance, if two users submit a form at the same time it is possible for your code to process both forms at roughly the same time and, depending on the timing, each might check and verify that a field is unique before either has inserted data into the database. As a result, it is possible to end up with a duplicate on a field that should be unique if validations only occur in your application code. On the other hand, if we set up constraints on your database one of the two submissions would fail when our Go code attempted to create a record with duplicate data.

In short, it is important to use database level constraints, and we will discuss each of these that we need as the need arises.
