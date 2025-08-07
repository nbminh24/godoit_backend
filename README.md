# Go-doit Backend API

Đây là project backend cho ứng dụng To-do List, được xây dựng bằng Golang và Gin framework. Project này là một phần của bài test tuyển dụng Intern Developer, tập trung vào việc xây dựng một REST API ổn định, có cấu trúc tốt và dễ dàng triển khai.

**Live API Endpoint:** [https://godoit-backend.onrender.com](https://godoit-backend.onrender.com)

---

## ✨ Tính năng

-   **Quản lý công việc:** Thêm, xem, cập nhật và xóa công việc.
-   **Lưu trữ trong bộ nhớ:** Dữ liệu được lưu trữ tạm thời trong bộ nhớ (in-memory), phù hợp cho mục đích demo và kiểm thử.
-   **Đóng gói với Docker:** Ứng dụng được đóng gói bằng Docker, giúp việc triển khai và chạy local trở nên nhất quán và đơn giản.
-   **Triển khai tự động:** Tích hợp với Render để tự động deploy mỗi khi có thay đổi trên nhánh `main`.

---

## 🛠️ Công nghệ sử dụng

-   **Ngôn ngữ:** [Go](https://golang.org/)
-   **Framework:** [Gin](https://github.com/gin-gonic/gin)
-   **Containerization:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
-   **Deployment:** [Render](https://render.com/)

---

## 🚀 Hướng dẫn sử dụng

### Chạy dự án ở Local

Có hai cách để chạy dự án này trên máy của bạn.

#### Cách 1: Sử dụng Go (Yêu cầu đã cài đặt Go 1.21+)

1.  **Clone repository:**
    ```bash
    git clone https://github.com/nbminh24/godoit_backend.git
    cd godoit_backend
    ```

2.  **Cài đặt dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Chạy server:**
    ```bash
    go run main.go
    ```
    Server sẽ chạy tại `http://localhost:8080`.

#### Cách 2: Sử dụng Docker (Khuyến khích)

Đây là cách đơn giản và nhất quán nhất để chạy dự án mà không cần cài đặt Go.

1.  **Yêu cầu:**
    -   [Docker](https://docs.docker.com/get-docker/)
    -   [Docker Compose](https://docs.docker.com/compose/install/)

2.  **Build và chạy container:**
    Mở terminal tại thư mục gốc của dự án và chạy lệnh:
    ```bash
    docker-compose up --build
    ```
    Để chạy ở chế độ nền, thêm cờ `-d`:
    ```bash
    docker-compose up --build -d
    ```
    Ứng dụng sẽ chạy tại `http://localhost:8080`.

3.  **Dừng container:**
    ```bash
    docker-compose down
    ```

---

## 📚 Tài liệu API

**URL cơ sở:** `https://godoit-backend.onrender.com`

### 1. Lấy danh sách công việc

-   **Method:** `GET`
-   **Endpoint:** `/tasks`
-   **Mô tả:** Trả về một mảng chứa tất cả các công việc.
-   **Ví dụ (cURL):**
    ```bash
    curl -X GET https://godoit-backend.onrender.com/tasks
    ```
-   **Response thành công (200 OK):**
    ```json
    [
        {
            "id": 1,
            "title": "Học Golang",
            "description": "Tìm hiểu về Gin framework",
            "completed": false
        },
        {
            "id": 2,
            "title": "Làm bài test",
            "description": "Hoàn thành phần backend",
            "completed": true
        }
    ]
    ```

### 2. Thêm công việc mới

-   **Method:** `POST`
-   **Endpoint:** `/tasks`
-   **Mô tả:** Tạo một công việc mới.
-   **Request Body (JSON):**
    ```json
    {
      "title": "Làm giao diện React",
      "description": "Sử dụng Axios để gọi API"
    }
    ```
-   **Ví dụ (cURL):**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"title":"Làm giao diện React","description":"Sử dụng Axios để gọi API"}' https://godoit-backend.onrender.com/tasks
    ```
-   **Response thành công (201 Created):**
    ```json
    {
      "id": 3
    }
    ```

### 3. Cập nhật trạng thái công việc

-   **Method:** `PUT`
-   **Endpoint:** `/tasks/{id}`
-   **Mô tả:** Cập nhật trạng thái `completed` của một công việc.
-   **Request Body (JSON):**
    ```json
    {
      "completed": true
    }
    ```
-   **Ví dụ (cURL):**
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d '{"completed":true}' https://godoit-backend.onrender.com/tasks/1
    ```
-   **Response thành công (200 OK):**
    ```json
    {
      "id": 1,
      "title": "Học Golang",
      "description": "Tìm hiểu về Gin framework",
      "completed": true
    }
    ```

### 4. Xóa một công việc

-   **Method:** `DELETE`
-   **Endpoint:** `/tasks/{id}`
-   **Mô tả:** Xóa một công việc dựa trên ID.
-   **Ví dụ (cURL):**
    ```bash
    curl -X DELETE https://godoit-backend.onrender.com/tasks/1
    ```
-   **Response thành công:** `204 No Content`
