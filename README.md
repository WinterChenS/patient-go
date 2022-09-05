# patient-go

## Quick Start

### Init Database

```mysql

CREATE DATABASE `patient` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- patient.patients definition

CREATE TABLE `patients` (
  `id` bigint NOT NULL,
  `name` varchar(35) COLLATE utf8mb4_general_ci NOT NULL,
  `brithday` timestamp NOT NULL,
  `phone` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `photo` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `driver_license` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `appointment_time` timestamp NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='patient info';

```

### Update config-dev.yml

```yaml
name: patient-go
mysql: root:root@tcp(127.0.0.1:3306)/patient?charset=utf8&parseTime=true&loc=Asia%2FShanghai
port: 8098
logsPath: "/tmp/logs/patient/"

minio:
  endpoint: "127.0.0.1:9000"
  accessKeyID: "admin"
  secretAccessKey: "admin"
  bucketName: "picture"
  path: "https://winterchen.com/share/"
```

### `go build`

```bash
go build
```

### `go run`

```bash
go run main.go
```
