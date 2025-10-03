# Requirements

* `make`
* `docker`

# Migrations

## Migrate

```sh
make migrate
```

NOTE: the db container should have been started before the migrations are run

## Make a migration

```sh
make migration name=some_migration_name
```

NOTE: there is ownership issues with docker
