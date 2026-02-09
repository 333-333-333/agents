---
name: db-naming-conventions
description: >
  Naming conventions for relational databases: databases, tables, columns, indexes, constraints, and enums.
  Applies to SQL migrations, sqlc queries, and ER diagrams in Mermaid.
  Trigger: When creating database tables, writing migrations, designing ER diagrams, or naming columns and constraints.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api, docs/diagrams]
  auto_invoke:
    - "Creating database tables or schemas"
    - "Writing SQL migrations"
    - "Designing ER diagrams with entity attributes"
    - "Naming columns, indexes, or constraints"
    - "Creating join tables for many-to-many relationships"
    - "Adding foreign keys to tables"
---

## When to Use

- Creating new database tables or modifying existing schemas
- Writing SQL migrations (`golang-migrate`, Flyway, etc.)
- Designing ER diagrams in Mermaid (`.mmd` files)
- Reviewing naming consistency between diagrams and migrations
- Adding indexes, constraints, or foreign keys

---

## Dependencies

| Skill | Purpose | Path |
|-------|---------|------|
| **`go-repository-pattern`** | Migration files, sqlc queries follow these naming rules | [`../go-repository-pattern/SKILL.md`](../go-repository-pattern/SKILL.md) |
| **`mermaid-diagrams`** | ER diagrams must use the same naming as SQL | [`../mermaid-diagrams/SKILL.md`](../mermaid-diagrams/SKILL.md) |

---

## Core Principle

**One naming convention governs everything**: ER diagrams, SQL migrations, sqlc queries, and application code mappings. Zero translation, zero ambiguity. What the diagram says is what the migration creates.

---

## Critical Rules

### Databases

| Rule | Convention | Example |
|------|-----------|---------|
| Format | `snake_case` | `bastet`, `bastet_production` |
| Environment suffix | Only when multiple DBs on same server | `bastet_dev`, `bastet_staging` |

### Tables

| Rule | Convention | Example |
|------|-----------|---------|
| Format | `snake_case` | `users`, `pet_sitters` |
| Plurality | **Always plural** | `bookings`, NOT `booking` |
| No prefixes | Never `tbl_`, `t_` | `users`, NOT `tbl_users` |
| Join tables | Alphabetical order or semantic name | `pets_services` or `pet_sitter_certifications` |

### Columns

| Rule | Convention | Example |
|------|-----------|---------|
| Format | `snake_case` | `first_name`, `phone_number` |
| Primary key | Always `id` | `id UUID PRIMARY KEY` |
| Foreign key | `{singular_table}_id` | `user_id`, `pet_sitter_id` |
| Booleans | Prefix `is_` or `has_` | `is_active`, `has_premium` |
| Timestamps | Suffix `_at` | `created_at`, `updated_at`, `deleted_at` |
| Counts | Suffix `_count` | `booking_count`, `review_count` |
| Monetary | Suffix `_cents` (store as integer) | `price_cents`, `fee_cents` |

### Indexes

| Type | Pattern | Example |
|------|---------|---------|
| Regular index | `idx_{table}_{columns}` | `idx_users_email` |
| Unique index | `uq_{table}_{columns}` | `uq_users_email` |
| Composite | `idx_{table}_{col1}_{col2}` | `idx_bookings_user_id_status` |

### Constraints

| Type | Pattern | Example |
|------|---------|---------|
| Primary key | `pk_{table}` | `pk_users` |
| Foreign key | `fk_{source}_{target}` | `fk_bookings_users` |
| Check | `ck_{table}_{column}` | `ck_bookings_status` |
| Unique | `uq_{table}_{columns}` | `uq_users_email` |

### Enum / Status Values

| Rule | Convention | Example |
|------|-----------|---------|
| Type name | `snake_case` | `booking_status`, `payment_status` |
| Values | `UPPER_SNAKE_CASE` | `PENDING`, `IN_PROGRESS`, `COMPLETED` |

---

## ER Diagrams in Mermaid

ER diagrams MUST mirror SQL naming exactly. This means:

| Element | In SQL | In Mermaid ER | Match? |
|---------|--------|---------------|--------|
| Table name | `users` | `users` | Yes |
| Column name | `first_name` | `first_name` | Yes |
| Foreign key | `user_id` | `user_id FK` | Yes |
| Primary key | `id` | `id PK` | Yes |
| Data type | `UUID` | `uuid` | Lowercase in Mermaid |

> See [assets/example-er.mmd](assets/example-er.mmd) for a complete ER diagram following these conventions.

> See [assets/example-migration.sql](assets/example-migration.sql) for the matching SQL migration.

---

## Anti-Patterns

| Don't | Do | Why |
|-------|------|-----|
| `UPPER_CASE` entity names in ER | `snake_case` plural matching SQL | Diagram must match implementation |
| `camelCase` columns in ER diagrams | `snake_case` matching SQL | Zero translation between diagram and migration |
| `userId` as foreign key | `user_id` | Consistent snake_case |
| `tbl_users`, `t_users` | `users` | Prefixes add noise, zero value |
| `booking` (singular table) | `bookings` (plural) | Table holds a collection |
| `active` (boolean without prefix) | `is_active` | Prefix makes intent obvious |
| `price DECIMAL` | `price_cents INTEGER` | Integer cents avoid float rounding |
| `data`, `info`, `details` columns | Specific names | Vague names hide meaning |
| `id INTEGER AUTO_INCREMENT` | `id UUID` | UUIDs are safer for distributed systems |

---

## Commands

```bash
# Create a migration with proper naming
migrate create -ext sql -dir migrations -seq create_bookings

# Validate table naming in existing migrations
rg "CREATE TABLE" migrations/ --no-heading
```

---

## Checklist

- [ ] Table names are `snake_case` and **plural**
- [ ] Columns are `snake_case` and **singular**
- [ ] Primary key is `id` (UUID)
- [ ] Foreign keys follow `{singular_table}_id` pattern
- [ ] Booleans have `is_` or `has_` prefix
- [ ] Timestamps end in `_at`
- [ ] Indexes follow `idx_{table}_{columns}` pattern
- [ ] Constraints follow `{type}_{table}_{detail}` pattern
- [ ] ER diagram entity names match SQL table names exactly
- [ ] ER diagram column names match SQL column names exactly
- [ ] Monetary values stored as integer cents with `_cents` suffix
