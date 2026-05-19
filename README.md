# Go REST API --- MVC Pattern with GORM & JWT Authentication

Project ini adalah implementasi REST API sederhana menggunakan bahasa Go
dengan menggunakan **Design Pattern MVC**, **GORM sebagai ORM**, dan
**JWT (JSON Web Token)** untuk sistem autentikasi.

API ini memiliki dua tabel utama:

1.  **Users** --- digunakan untuk autentikasi (register & login)
2.  **Tasks** --- digunakan untuk operasi CRUD (Create, Read, Update,
    Delete)

## 🚀 Tech Stack

-   **Go**
-   **Gin Web Framework**
-   **GORM (MySQL)**
-   **JWT Authentication**
-   **MVC Project Structure**

## 🧩 Project Structure (MVC)

    /api
     ├── config
     │    └── database.go
     ├── controllers
     │    ├── auth_controller.go
     │    └── task_controller.go
     ├── middleware
     │    └── jwt_middleware.go
     ├── models
     │    ├── user.go
     │    └── task.go
     ├── routes
     │    └── routes.go
     └── main.go

## 🗄️ Database Schema

### users table

 | Field       | Type      | Description             |
|-------------|-----------|--------------------------|
| id          | uint (PK) | primary key              |
| name        | string    | nama pengguna            |
| email       | string    | unique email             |
| password    | string    | hashed password (bcrypt) |

### tasks table

| Field       | Type      | Description                    |
|-------------|-----------|--------------------------------|
| id          | uint (PK) | primary key                    |
| title       | string    | judul task                     |
| description | string    | detail tugas                   |
| status      | enum      | todo / doing / done            |
| due_date    | date      | batas waktu                    |

### Authentication

-   POST /auth/register
-   POST /auth/login

### Tasks

-   GET /tasks
-   GET /tasks/:id
-   POST /tasks
-   PUT /tasks/:id
-   DELETE /tasks/:id

## ⚙️ Cara Menjalankan

    go mod tidy
    go run main.go
