* Connect to the db:

```bash
psql postgres://postgres:postgres@localhost:5432/gator
```

Useful:
* `\dt`: lists all tables in the db

* Execute Migrations

```bash
cd ~/Projects/bootdev-learning/go/blogaggr/sql/schema
goose postgres postgres://postgres:postgres@localhost:5432/gator up
goose postgres postgres://postgres:postgres@localhost:5432/gator down
```

* Write migrations inside `sql/schema/` and queries inside `sql/queries/`, and then generate the relevant Go code by running `sqlc generate` from the project root.
