# Godoit - Backend API

🚀 **Live API:** [https://godoit-backend.onrender.com/tasks](https://godoit-backend.onrender.com/tasks) 🚀

Đây là backend API cho ứng dụng quản lý công việc **Godoit**, được xây dựng hoàn toàn bằng **Golang**. Hệ thống cung cấp các endpoint RESTful để quản lý công việc một cách hiệu quả, an toàn và có hiệu suất cao. Dự án này được phát triển như một phần của bài kiểm tra kỹ năng cho vị trí Intern Developer.

## Các tính năng chính

- **API RESTful Toàn diện:** Cung cấp đầy đủ các hoạt động **CRUD** (Create, Read, Update, Delete) cho việc quản lý công việc.
- **Hiệu suất cao:** Xây dựng trên nền tảng **Gin**, một trong những web framework nhanh nhất của Golang, đảm bảo thời gian phản hồi API tối thiểu.
- **Cấu trúc rõ ràng:** Mã nguồn được tổ chức theo các module `handlers`, `store`, và `models`, giúp dễ dàng bảo trì và mở rộng.
- **Hỗ trợ CORS:** Cấu hình sẵn sàng để chấp nhận các yêu cầu từ frontend một cách an toàn.
- **Sẵn sàng cho Containerization:** Đi kèm với `Dockerfile` và `docker-compose.yml` để dễ dàng đóng gói và triển khai.

## Công nghệ và Công cụ

- **Ngôn ngữ:** [Golang](https://go.dev/)
- **Web Framework:** [Gin](https://gin-gonic.com/)
- **Containerization:** [Docker](https://www.docker.com/)
- **Deployment:** [Render](https://render.com/)

## Hướng dẫn sử dụng

### Kiểm thử với Postman

Để dễ dàng kiểm thử các API endpoint, bạn có thể sử dụng Postman Collection đã được cung cấp sẵn.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/a8a5f3c5-7c1f-4d2e-9d3d-7e7e5e7e5e7e?action=collection%2Ffork&source=rip_markdown&collection-url=https%3A%2F%2Fraw.githubusercontent.com%2Fnbminh24%2Fgodoit_backend%2Fmain%2Fgodoit_backend.postman_collection.json)

**Các bước thực hiện:**

1.  Nhấn vào nút **Run in Postman** ở trên.
2.  Chọn workspace của bạn trong Postman để sao chép (fork) collection này về.
3.  Collection đã được cấu hình sẵn để sử dụng API đã deploy. Bạn có thể bắt đầu gửi request ngay lập tức.

### Cài đặt và Chạy dự án local

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/nbminh24/godoit_backend.git
    cd godoit_backend
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Run the development server:**
    Ứng dụng sẽ chạy tại `http://localhost:8080`.
    ```sh
    go run main.go
    ```

### Sử dụng với Docker

1.  **Build và chạy container:**
    ```sh
    docker-compose up --build
    ```
    API sẽ có sẵn tại `http://localhost:8080`.  **Yêu cầu:**
    -   [Docker](https://docs.docker.com/get-docker/)
    -   [Docker Compose](https://docs.docker.com/compose/install/)

2.  **Build và chạy container:**
    Mở terminal tại thư mục gốc của dự án và chạy lệnh:
    ```sh
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
