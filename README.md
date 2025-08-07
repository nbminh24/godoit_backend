# Go To-Do List Backend

Đây là backend cho ứng dụng quản lý công việc (To-do List), được xây dựng bằng Golang và Gin framework.

## Yêu cầu

- [Golang](https://go.dev/dl/) (phiên bản 1.20 trở lên)

## Hướng dẫn cài đặt và sử dụng

1.  **Clone a repository (Nếu có):**
    ```bash
    git clone <your-repo-url>
    cd godoit_backend
    ```

2.  **Tải các thư viện cần thiết:**
    Lệnh này sẽ tự động tải Gin và các thư viện khác được định nghĩa trong `go.mod`.
    ```bash
    go mod tidy
    ```

3.  **Chạy server:**
    ```bash
    go run main.go
    ```
    Server sẽ chạy tại `http://localhost:8080`.

## Hướng dẫn sử dụng với Docker

Phần này hướng dẫn cách chạy ứng dụng Go-doit bằng Docker và Docker Compose.

### Yêu cầu

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/) (Thường được cài đặt sẵn cùng với Docker Desktop)

### Chạy ứng dụng

1.  **Build và chạy container:**

    Mở terminal hoặc command prompt, điều hướng đến thư mục gốc của dự án và chạy lệnh sau:

    ```bash
    docker-compose up --build
    ```

    Lệnh này sẽ:
    - Build Docker image cho ứng dụng dựa trên `Dockerfile`.
    - Tạo và khởi chạy một container từ image đó.
    - Ứng dụng sẽ chạy và lắng nghe ở cổng `8080`.

2.  **Chạy ở chế độ nền (detached mode):**

    Để chạy container ở chế độ nền, sử dụng cờ `-d`:

    ```bash
    docker-compose up --build -d
    ```

### Dừng ứng dụng

Để dừng và xóa container, chạy lệnh sau:

```bash
docker-compose down
```

## API Endpoints

Sau khi ứng dụng đã chạy, bạn có thể tương tác với các API endpoints sau tại `http://localhost:8080`:

- `GET /tasks`: Lấy danh sách tất cả công việc.
- `POST /tasks`: Thêm một công việc mới. Body: `{"title": "Nội dung công việc", "is_done": false}`
- `PUT /tasks/{id}`: Cập nhật trạng thái của một công việc. Body: `{"is_done": true}`
- `DELETE /tasks/{id}`: Xóa một công việc.
