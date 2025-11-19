# Сервис проверки доступности сайтов (Health Checker)

Простой REST API сервис на Go для мониторинга состояния веб-сайтов.
Главное преимущество — использование **горутин** для одновременной проверки множества ссылок.

## Особенности
- **Конкурентность:** Сайты проверяются параллельно, а не по очереди. Это сильно ускоряет работу.
- **Pure Go:** Написано на стандартной библиотеке (`net/http`).
- **JSON API:** Принимает и отдает данные в удобном формате JSON.

## Как запустить

1. **Склонировать репозиторий:**
   ```bash
   git clone https://github.com/ТВОЙ_НИК/go-health-checker.git
   cd go-website-health-checker
   ```

2. **Запустить сервер:**
   ```bash
   go run main.go
   ```
   *Сервер запустится на порту 8080.*

## Пример использования

Отправьте POST-запрос на `http://localhost:8080/check` со списком сайтов.

**Пример запроса (curl):**
```bash
curl -X POST http://localhost:8080/check \
-H "Content-Type: application/json" \
-d '{"websites": ["https://ya.ru", "https://google.com", "https://github.com"]}'
```

**Ответ сервиса:**
```json
[
  {
    "website": "https://ya.ru",
    "status": "UP",
    "latency": 120000000
  },
  {
    "website": "https://google.com",
    "status": "UP",
    "latency": 145000000
  }
]
```

## Стек
- **Язык:** Go
- **Библиотеки:** Standard Library (net/http, encoding/json, sync)