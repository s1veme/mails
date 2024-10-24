### REST API Mails

#### 1. **Регистрация пользователя**
- **POST users/register**
- **Описание**: Регистрация нового пользователя.

**Request:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["User registered successfully"]
}
```

#### 2. **Логин пользователя**
- **POST users/login**
- **Описание**: Авторизация и получение JWT токенов.

**Request:**
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1...",
    "refresh_token": "eyJhbGciOiJIUzI2..."
  },
  "errors": null,
  "messages": ["Login successful"]
}
```

#### 3. **Получение текущего пользователя (по JWT токену)**
- **GET users/me**
- **Описание**: Получение информации о текущем пользователе.

**Response:**
```json
{
  "data": {
    "id": "cbbbd1af-36d8-4a2d-b829-9bf08b39a3ae",
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-10-23T12:34:56Z",
    "updated_at": "2024-10-23T12:34:56Z"
  },
  "errors": null,
  "messages": []
}
```

#### 4. **Создание папки**
- **POST /folders**
- **Описание**: Создание новой папки для пользователя.

**Request:**
```json
{
  "name": "Work"
}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Folder created successfully"]
}
```

#### 5. **Отправка письма**
- **POST /mails**
- **Описание**: Отправка нового письма.

**Request:**
```json
{
  "subject": "Meeting Reminder",
  "body": "This is a reminder for our meeting tomorrow at 10 AM.",
  "recipients_to": ["recipient1@example.com"],
  "recipients_cc": ["recipient2@example.com"],
  "recipients_bcc": ["recipient3@example.com"]
}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Mail sent successfully"]
}
```

#### 6. **Получение списка писем с пагинацией**
- **GET /mails**
- **Описание**: Получение списка писем с поддержкой пагинации.

**Response:**
```json
{
  "data": {
    "items": [
      {
        "id": "d1e7f1ff-9134-4fbc-8f0b-0b793cf61c0e",
        "subject": "Meeting Reminder",
        "body": "This is a reminder for our meeting tomorrow at 10 AM.",
        "sender_id": "cbbbd1af-36d8-4a2d-b829-9bf08b39a3ae",
        "created_at": "2024-10-23T12:36:00Z",
        "updated_at": "2024-10-23T12:36:00Z",
        "recipients": [
          {"email": "recipient1@example.com", "type": "TO"},
          {"email": "recipient2@example.com", "type": "CC"},
          {"email": "recipient3@example.com", "type": "BCC"}
        ]
      }
    ],
    "pagination": {
      "all_count": 100,
      "limit": 10,
      "offset": 0,
      "page": 1
    }
  },
  "errors": null,
  "messages": []
}
```

#### 7. **Обновление письма**
- **PATCH /mails/{id}**
- **Описание**: Обновление письма.

**Request:**
```json
{
  "is_important": true,
}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Mail marked as important"]
}
```

#### 8. **Обновление токена**
- **PATCH /users/refresh**
- **Описание**: Обновление токена.

**Request:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI2..."
}
```

**Response:**
```json
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1...",
    "refresh_token": "eyJhbGciOiJIUzI2..."
  },
  "errors": null,
  "messages": ["Tokens refreshed successfully"]
}
```

#### 9. **Добавление письма в папку**
- **POST /folders/{folder_id}/mails**
- **Описание**:  Добавляет письмо в указанную папку.

**Request:**
```json
{
  "mail_id": "3fa85f64-5717-4562-b3fc-2c963f66afa6"
}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Mail added to folder successfully"]
}
```

#### 10. **Удаление письма из папки**
- **DELETE /folders/{folder_id}/mails/{mail_id}**
- **Описание**:  Удаляет письмо из указанной папки.

**Request:**
```json
{}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Mail removed from folder successfully"]
}
```

#### 11. **Удаление папки**
- **DELETE /folders/{folder_id}**
- **Описание**:  Удаляет указанную папку.

**Request:**
```json
{}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Folder deleted successfully"]
}
```

### 12. **Удаление письма**
- **DELETE /mails/{mail_id}**
- **Описание**:  Удаляет указанное письмо.

**Request:**
```json
{}
```

**Response:**
```json
{
  "data": {},
  "errors": null,
  "messages": ["Mail deleted successfully"]
}
```
