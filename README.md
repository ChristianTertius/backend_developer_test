# Backend Developer Test - Booking ToGO

Complete solution for the Backend Developer Test from Booking ToGO, by **Christian Tertius**.

This repo consists of three parts:

```
backend_developer_test/
├── database/            # 1. PostgreSQL schema (schema.sql)
├── golang-service/      # 2. CRUD web service (Golang, clean architecture, gorilla/mux)
└── laravel-frontend/    # 3. CRUD view (Laravel 9, Blade, dynamic family form)
```

## Task Summary

1. **Database** — create the schema (Customer, family_list, Nationality) in PostgreSQL.
2. **Process**
   - 2.1 CRUD for user + family via a **Golang** web service (clean architecture, **gorilla/mux** router).
   - 2.2 CRUD view for users via **Laravel 8/9**, with a family form that supports add/remove, submitted via a Laravel Form.

## Database Schema

- **nationality** (`nationality_id` PK, `nationality_name`, `nationality_code`)
- **customer** (`cst_id` PK, `nationality_id` FK, `cst_name`, `cst_dob`, `cst_phoneNum`, `cst_email`)
- **family_list** (`fl_id` PK, `cst_id` FK -> customer, `fl_relation`, `fl_name`, `fl_dob`)

Relationship: one `customer` has one `nationality` and many `family_list` entries
(`family_list` cascades on delete when the customer is removed).

## Golang Endpoint

| Method | Path | Keterangan |
|--------|------|------------|
| GET | `/api/customers` | List semua user + keluarga |
| POST | `/api/customers` | Tambah user + keluarga |
| GET | `/api/customers/{id}` | Detail user |
| PUT | `/api/customers/{id}` | Update user + keluarga |
| DELETE | `/api/customers/{id}` | Hapus user |
| GET | `/api/nationalities` | List kewarganegaraan |
| GET | `/api/health` | Health check |


## Tech Stack

- Go 1.26.3, gorilla/mux, lib/pq, godotenv
- Laravel 9, Blade, Bootstrap 5
- PostgreSQL
