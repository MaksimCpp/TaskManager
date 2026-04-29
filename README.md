<div align="center">

# 🧩 Task Manager API

**Современный REST API для управления задачами с JWT-аутентификацией**

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go\&logoColor=white)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16+-336791?logo=postgresql\&logoColor=white)](https://www.postgresql.org/)
[![JWT](https://img.shields.io/badge/Auth-JWT-black?logo=jsonwebtokens\&logoColor=white)](https://jwt.io/)

</div>

## 👨‍💻 Автор: @MaksimCpp

---

## 🚀 Возможности

* Регистрация и авторизация пользователей
* Аутентификация через **JWT (Bearer Token)**
* CRUD операции с задачами
* Поддержка описания задач
* Изоляция данных пользователей (каждый видит только свои задачи)
* Чистая архитектура (Clean Architecture)
* Разделение слоёв: domain / usecase / repository / delivery
* Работа с тегами (в разработке)

---

## 🏗️ Технологии

| Технология    | Назначение               |
| ------------- | ------------------------ |
| Go            | Основной язык разработки |
| net/http      | HTTP сервер              |
| PostgreSQL    | База данных              |
| pgx (pgxpool) | Драйвер и пул соединений |
| JWT           | Аутентификация           |
| bcrypt        | Хеширование паролей      |

---

## 📋 Основные эндпоинты

### 🔓 Аутентификация

| Метод  | Путь        | Описание                 |
| ------ | ----------- | ------------------------ |
| `POST` | `/users` | Регистрация пользователя |
| `POST` | `/users/login`    | Получение JWT токена     |

---

### 🔒 Задачи (требуется JWT)

| Метод    | Путь     | Описание                            |
| -------- | -------- | ----------------------------------- |
| `POST`   | `/tasks` | Создать задачу (title, description) |
| `GET`    | `/tasks` | Получить список задач               |
| `PATCH`  | `/tasks` | Обновить задачу                     |
| `DELETE` | `/tasks` | Удалить задачу (`?id=`)             |

---

### 🏷️ Теги (в разработке)

| Метод  | Путь          | Описание               |
| ------ | ------------- | ---------------------- |
| `POST` | `/tags`       | Создать тег            |
| `GET`  | `/tags`       | Получить список тегов  |
| `POST` | `/tasks/tags` | Привязать тег к задаче |

---

## 🔐 Аутентификация

Используется **JWT (JSON Web Token)**.

Каждый защищённый запрос должен содержать заголовок:

```http
Authorization: Bearer <token>
```

### 🔄 Flow:

1. Пользователь логинится
2. Получает JWT токен
3. Передаёт его в каждом запросе
4. Middleware:

   * валидирует токен
   * извлекает `user_id`
   * добавляет его в `context`


---

## 🗄️ Структура базы данных

### schema

```sql
CREATE SCHEMA taskschema;
```

---

### users

```sql
id UUID PRIMARY KEY
email VARCHAR(255) UNIQUE NOT NULL
password VARCHAR(100) NOT NULL
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
```

---

### tasks

```sql
id UUID PRIMARY KEY
title VARCHAR(255) NOT NULL
description VARCHAR(1000)
completed BOOLEAN NOT NULL DEFAULT FALSE
user_id UUID REFERENCES taskschema.users(id) NOT NULL
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
```

---

### tags

```sql
id UUID PRIMARY KEY
name VARCHAR(100) UNIQUE NOT NULL
```

---

### tasks_tags (many-to-many)

```sql
task_id UUID REFERENCES taskschema.tasks(id) NOT NULL
tag_id UUID REFERENCES taskschema.tags(id) NOT NULL
PRIMARY KEY (task_id, tag_id)
```

---

## 🔒 Безопасность

* Пароли хранятся в виде bcrypt-хеша
* JWT подписывается секретным ключом
* Все операции с задачами ограничены `user_id`

---
