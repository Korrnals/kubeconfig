# kubeconfig

**Лучший инструмент для объединения множественных kubeconfig-файлов от `kubeadm`**

> Решена главная боль всех, кто работает с несколькими кластерами:  
> **дублирование пользователя `kubernetes-admin` → потеря сертификатов**

Ты больше **никогда** не потеряешь доступ к старому кластеру после `kubeadm init`.

---

## Зачем это нужно?

Когда ты запускаешь `kubeadm init` на нескольких серверах — каждый раз создаётся:

- Уникальный кластер
- Уникальный контекст  
- → Но **одинаковый пользователь** `kubernetes-admin`

При обычном merge — остаётся **только последний** пользователь → **все старые сертификаты теряются**

**kubeconfig решает это раз и навсегда.**

---

### Что делает утилита

```bash
kubeconfig --dir ./my-clusters/
```

**→ Автоматически:**

- Читает любые файлы из директории, проверяя содержимое на наличие `apiVersion: v1` и `kind: Config` поля - тем самым фильтрует файлы.
- Обнаруживает `kubernetes-admin`
- Переименовывает в `kubernetes-admin@<cluster-name>`
- Обновляет ссылки в контекстах
- Объединяет кластеры, контексты, пользователей
- Сохраняет **все сертификаты**

---

### Возможности (v1.0.0)

| Функция                        | Команда                                  | Описание |
|--------------------------------|------------------------------------------|--------|
| Объединить конфиги             | `--dir ./configs` или `-d`               | Из директории |
| Сохранить результат            | `--output ~/.kube/config` или `-o`       | В любой файл |
| Показать список кластеров      | `--list` или `-l`                        | Как `kubectl config get-contexts` |
| Показать версию                | `--version` или `-v`                     | Информация о версии |
| Автоматическое создание `.kube`| Да                                       | Создаёт директорию и файл |

---

### Пример использования

```bash
# 0. Проверка версии утилиты
kubeconfig -v
##  и получения списка флагов
kubeconfig -h

# 1. Объединить все конфиги и сохранить в ~/.kube/config
kubeconfig -d ~/.kube/single-kubeconfig/ -o ~/.kube/config

# 2. Просто посмотреть, что получится
kubeconfig -d ./my-clusters/

# 3. Показать текущие кластеры (даже без --dir!)
kubeconfig --list
```

---

### Установка

#### Вариант 1 — Готовый бинарник (рекомендуется)

**Самый быстрый способ — одна команда:**

- **Linux** (x86_64)

  ```bash
  curl -L https://github.com/Korrnals/kubeconfig/releases/download/v1.0.0/kubeconfig-linux-amd64 \
    -o kubeconfig && chmod +x kubeconfig && sudo mv kubeconfig /usr/local/bin/
  ```

- **macOS** (Intel)

  ```bash
  curl -L https://github.com/Korrnals/kubeconfig/releases/download/v1.0.0/kubeconfig-darwin-amd64 \
    -o kubeconfig && chmod +x kubeconfig && sudo mv kubeconfig /usr/local/bin/
  ```

- **macOS** (Apple Silicon / M1–M3)

  ```bash
  curl -L https://github.com/Korrnals/kubeconfig/releases/download/v1.0.0/kubeconfig-darwin-arm64 \
    -o kubeconfig && chmod +x kubeconfig && sudo mv kubeconfig /usr/local/bin/
  ```

#### Вариант 2 — Сборка из исходников

Обычная сборка (текущая ОС):

```bash
# Склонировать и собрать
git clone https://github.com/Korrnals/kubeconfig.git
cd kubeconfig
go build -o kubeconfig cmd/kubeconfig/main.go

# Переместить в PATH
sudo mv kubeconfig /usr/local/bin/
```

Кросс-компиляция под все популярные платформы одной командой:

```bash
git clone https://github.com/Korrnals/kubeconfig.git
cd kubeconfig

# Создаём все бинарники
GOOS=linux   GOARCH=amd64 go build -o kubeconfig-linux-amd64   cmd/kubeconfig/main.go
GOOS=darwin  GOARCH=amd64 go build -o kubeconfig-darwin-amd64  cmd/kubeconfig/main.go
GOOS=darwin  GOARCH=arm64 go build -o kubeconfig-darwin-arm64  cmd/kubeconfig/main.go

# Перемещаем нужный в PATH (пример для Linux)
chmod +x kubeconfig-linux-amd64
sudo mv kubeconfig-linux-amd64 /usr/local/bin/kubeconfig
```

---

### Цели развития

| Задача                                 | Статус     | Приоритет |
|---------------------------------------|------------|---------|
| `--switch-context <name>`             | Запланировано | ★★★★★ |
| `--current-context`                   | Запланировано | ★★★★ |
| `--rename-context old new`            | Запланировано | ★★★ |
| Поддержка `$KUBECONFIG`               | Запланировано | ★★★★ |
| Интеграция с HashiCorp Vault          | В планах   | ★★★ |
| *Утилита как сервис                   | В планах   | ★★★ |
| GUI версия (TUI)                      | Идея       | ★★ |
| Плагин для Lens / OpenLens            | Идея       | ★★ |

>[!TIP] Примечание:
> `*` _**Утилита как сервис**_- предполагается реализовать функционал, который позволит через файл-конфигурации указать `data_source` (источники) где располагаются конфигурации Kubernetes:
> - динамическое обновление основного `~/.kube/config` при добавлении нового кластера;
> - актуализация конфигурации, если кластера обновили свои сертификаты;

---

### Почему это лучше, чем всё остальное

| Инструмент       | Сохраняет все сертификаты? | Умное переименование? | Простота |
|------------------|----------------------------|------------------------|----------|
| `kubectl`        | Нет                        | Нет                    | ★★★★ |
| `kubecm merge`   | Нет                        | Нет                    | ★★★ |
| `kubie`          | Да                         | Частично               | ★★ |
| **kubeconfig** | **Да**                 | **Да**                 | **★★★★★** |

---

### Maintainers

---

**Автор:** Korrnals  
**GitHub:** <https://github.com/Korrnals/kubeconfig>  
**Версия:** v1.0.0 (2025-12-09)

---
