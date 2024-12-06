## Project Specs

- Framework : Golang Gin Framework
- Database : MySQL
- ORM: native

## Playing with migration

1. Generate migration file:

   ```
   make migrate-create name=create_yourtablename_table
   ```

   it will generate file migration up.sql & down.sql in folder scripts/migrations

2. Up table:

   ```
   make migrate-up
   ```

3. Down table:
   ```
   make migrate-down
   ```
