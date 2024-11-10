# Сервис сокращения URL

Данный репозиторий содержит сервис сокращения URL, реализованный на языке Go. Он позволяет пользователям создавать сокращенные ссылки, отслеживать статистику их использования и управлять существующими сокращенными URL через удобный API. Сервис обладает продуманной архитектурой и включает поддержку базы данных, настройку срока действия ссылок и документацию Swagger для удобства интеграции.

## Функциональные возможности

1. **Сокращение URL**: Генерирует уникальные короткие ссылки для длинных URL, предоставленных пользователем. Обрабатывает дублирующиеся запросы, чтобы избежать создания избыточных записей.
2. **Срок действия ссылок**: Сокращенные ссылки имеют настраиваемый срок действия (например, 30 дней), после которого они автоматически становятся недействительными.
3. **Перенаправление по URL**: Перенаправляет пользователя с короткой ссылки на оригинальный длинный URL.
4. **Управление URL**:
   - **Просмотр списка ссылок**: Позволяет просматривать все созданные пользователем ссылки.
   - **Удаление ссылки**: Удаляет указанную короткую ссылку.
5. **Статистика**: Отслеживает количество переходов и дату последнего использования для каждой короткой ссылки.
6. **Обработка ошибок**: Включает проверку и обработку ошибок для случаев некорректного ввода URL или несуществующих ссылок.
7. **Документация Swagger**: Предоставляет интерактивную документацию API для удобства тестирования и интеграции.

## API Эндпоинты

### 1. Создание короткой ссылки

- **Эндпоинт**: `POST /shortener`
- **Тело запроса**:
  ```json
  {
    "url": "https://example.com"
  }
  ```
- **Описание**: Принимает JSON с ключом `url`, содержащим длинный URL. Возвращает идентификатор для сокращенной ссылки.
- **Валидация**: Проверяет URL на валидность и непустое значение.
- **Ответ**:
  ```json
  {
    "short-url"
  }
  ```

### 2. Просмотр списка коротких ссылок

- **Эндпоинт**: `GET /shortener`
- **Описание**: Возвращает список всех созданных пользователем коротких ссылок.

### 3. Перенаправление на длинный URL

- **Эндпоинт**: `GET /{link}`
- **Описание**: Перенаправляет на оригинальный длинный URL при доступе по короткой ссылке.
- **Обработка ошибок**: Возвращает 404, если ссылка не найдена или срок её действия истек.

### 4. Удаление короткой ссылки

- **Эндпоинт**: `DELETE /{link}`
- **Описание**: Удаляет указанную короткую ссылку.

### 5. Статистика по ссылке

- **Эндпоинт**: `GET /stats/{link}`
- **Описание**: Предоставляет статистику для указанной короткой ссылки, включая общее количество переходов и дату последнего использования.
- **Ответ**:
  ```json
  {
    "click_counts": 5,
    "last_entered_at": "2023-11-09T14:48:00Z"
  }
  ```

## Документация Swagger

Swagger-документация предоставляет интерактивное тестирование и визуализацию API. Она включает описание всех эндпоинтов, схем запросов/ответов и сообщений об ошибках.

### Доступ к Swagger UI

Для доступа к документации Swagger откройте `http://<service-url>/swagger/index.html`.

## [Установка и настройка](https://github.com/Tr8ch/url-shortener/DEVELOPMENT.md)