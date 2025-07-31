# go-chat-app-dh
https://youtu.be/W9SuX9c40s8

1. Postgres
    -> Create a go chat database
2. Database setup
    -> Make a connection to the db
    -> add a db migration file to create the `users` table
3. /signup end point to create a new user
    - repository <- Service <- handler (dependencies)

4. /login and /logout endpoints
    - jwt with http-only cookie
       should be prone to csrf attacks
    - vs token based authentication
        - prome to xss

        OR We can use
    : Best: short-lived access token + refresh token

5. After installing migrate

    - migrate create -ext sql -dir db/migrations add_users_table

        This will create two different files in migrations folder.
            1. up .sql  # up migration
            2. down .sql # down migration

    Inside add_users_table.up.sql migration file add
        create table "users"(
            "id" bigserial primary key,
            "username" varchar not null,
            "email" varchar not null,
            "password" varchar not null
        )

    Inside add_users_table.down.sql migration file add
        DROP TABLE IF EXISTS users;


6. You can create a Make file and add this command in it.
    
    // Below command for up migration

    migrate -path db/migrations -database "postgresql://postgres:user@localhost:5432/gochat?sslmode=disable" -verbose up

    Now you can check. If this table has been created once migrated

    // Below command for down migration

    migrate -path db/migrations -database "postgresql://postgres:user@localhost:5432/gochat?sslmode=disable" -verbose down