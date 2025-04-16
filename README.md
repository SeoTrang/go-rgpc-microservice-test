
# Hướng Dẫn Dự Án gRPC Microservices với Go + Gin và NestJS

Dự án này xây dựng một hệ thống microservices sử dụng gRPC cho giao tiếp giữa các service, với các service chính bao gồm `api-gateway`, `order-service`, `product-service`, và `user-service`. Dưới đây là hướng dẫn chi tiết về cách cấu hình và chạy hệ thống.

## Môi Trường Phát Triển

Dự án này được phát triển với các công nghệ chính sau:

- **Go + Gin** cho các service chính (`order-service`, `product-service`, `user-service`).
- **NestJS** cho `order-service`, giao tiếp qua gRPC.
- **gRPC** cho giao tiếp giữa các service.
- **Yarn** cho quản lý dependencies trong dự án NestJS.

## Cấu Trúc Thư Mục Dự Án

Dự án có cấu trúc thư mục như sau:

```
/api-gateway/
  - Chịu trách nhiệm giao tiếp HTTP với client và giao tiếp với các service qua gRPC.

order-service/
  - Xây dựng bằng NestJS, sử dụng gRPC để giao tiếp với các service khác.

product-service/
  - Xây dựng bằng Go + Gin, cung cấp API gRPC cho `api-gateway` và các service khác.

user-service/
  - Xây dựng bằng Go + Gin, cung cấp API gRPC cho `api-gateway` và các service khác.
```

## Cài Đặt

1. **Cài Đặt dependencies cho API Gateway (Go + Gin):**

   - Đảm bảo bạn đã cài Go.
   - Vào thư mục `api-gateway` và chạy lệnh sau để cài đặt dependencies:

   ```bash
   go mod tidy
   ```

2. **Cài Đặt dependencies cho `order-service` (NestJS):**

   - Vào thư mục `order-service` và cài đặt dependencies bằng Yarn:

   ```bash
   yarn install
   ```

3. **Cài Đặt dependencies cho `product-service` và `user-service` (Go + Gin):**

   - Vào thư mục `product-service` và `user-service`, sau đó chạy lệnh:

   ```bash
   go mod tidy
   ```

## Cấu Hình gRPC

Dự án sử dụng gRPC cho giao tiếp giữa các service. Mỗi service đều có một file `.proto` để định nghĩa các API.

### Cấu Hình gRPC trong `order-service`:

- `order-service` sử dụng gRPC để giao tiếp với các service khác như `user-service` và `product-service`. 
- File `.proto` nằm trong thư mục `proto` của `order-service` và được tham chiếu trong cấu hình NestJS.

### Cấu Hình gRPC trong `user-service` và `product-service`:

- Các service này cung cấp các API gRPC cho `api-gateway` và các service khác.

## Chạy Dự Án

1. **Chạy API Gateway (Go + Gin):**

   - Vào thư mục `api-gateway` và chạy lệnh sau:

   ```bash
   go run main.go
   ```

2. **Chạy `order-service` (NestJS):**

   - Vào thư mục `order-service` và chạy lệnh sau:

   ```bash
   yarn start:dev
   ```

3. **Chạy `product-service` và `user-service` (Go + Gin):**

   - Vào thư mục `product-service` và `user-service`, sau đó chạy lệnh sau:

   ```bash
   go run main.go
   ```


## Các Service trong Dự Án

### 1. `api-gateway`:

- Chịu trách nhiệm nhận yêu cầu HTTP từ client và gửi yêu cầu đến các service thông qua gRPC.
- Cung cấp API cho frontend để giao tiếp với các service backend.

### 2. `order-service` (NestJS):

- Xây dựng bằng NestJS, sử dụng gRPC để giao tiếp với service `user-service`.

### 3. `product-service` (Go + Gin):

- Xây dựng bằng Go + Gin, cung cấp API gRPC để quản lý thông tin sản phẩm.
- Có thể giao tiếp với `user-service`.

### 4. `user-service` (Go + Gin):

- Xây dựng bằng Go + Gin, cung cấp API gRPC để quản lý thông tin người dùng.