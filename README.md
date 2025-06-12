# Migrations
references: https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md

## 1. Create new migration file

```bash
migrate create -ext sql -dir db/migrations YOUR_NEW_MIGRATION_VERB
```

## 2. Run migration

```bash
migrate -database YOUR_DATABASE_URL -path db/migrations up
```

