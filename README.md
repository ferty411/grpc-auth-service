# 🛡️ gRPC SSO Auth Service

Микросервис единой точки входа (Single Sign-On) для безопасной аутентификации, авторизации и управления доступом пользователей. Построен на **Go** с использованием **gRPC** протокола для высокопроизводительного межсервисного взаимодействия.

## 🌟 Особенности архитектуры

Проект строго следует принципам слоистой архитектуры (Layered Architecture) с использованием интерфейсов (Dependency Injection), что позволяет легко тестировать и масштабировать приложение:

*   **Transport Layer (`internal/grpc`)**: Реализация gRPC-сервера, валидация входящих запросов.
*   **Service Layer (`internal/services`)**: Изолированная бизнес-логика (генерация токенов, хэширование).
*   **Storage Layer (`internal/storage/sqlite`)**: Паттерн Repository для работы с базой данных, изолирующий SQL-запросы от сервисного слоя.

## 🛠 Технологический стек

*   **Язык:** Go (Golang)
*   **API / Транспорт:** gRPC, Protocol Buffers
*   **База данных:** SQLite (через `mattn/go-sqlite3`)
*   **Безопасность:** 
    *   Генерация токенов: `golang-jwt/jwt/v5`
    *   Криптография паролей: `golang.org/x/crypto/bcrypt`
*   **Конфигурация:** `ilyakaznacheev/cleanenv` (чтение из `.yaml` и ENV-переменных)
*   **Логирование:** `log/slog` (структурированные логи)

## 📡 gRPC Контракты (API)

Сервис реализует следующие RPC методы:

*   `Register(RegisterRequest) returns (RegisterResponse)` — Хэширует пароль и создает новую учетную запись. Возвращает уникальный `user_id`. Поддерживает обработку конфликтов (уникальность email).
*   `Login(LoginRequest) returns (LoginResponse)` — Верифицирует учетные данные и генерирует подписанный **JWT-токен** на основе секрета приложения.
*   `IsAdmin(IsAdminRequest) returns (IsAdminResponse)` — Проверяет наличие флага администратора у пользователя (используется для проверки прав доступа в других микросервисах).

## 🚀 Быстрый старт

### 1. Клонирование репозитория
```bash
git clone [https://github.com/ferty411/grpc-auth-service.git](https://github.com/ferty411/grpc-auth-service.git)
cd grpc-auth-service
