# Go Hexagonal Template

Template สำหรับเริ่มต้นพัฒนา Go Backend ด้วยแนวคิด **Hexagonal Architecture** โดยใช้ **Fiber** เป็น HTTP Framework

Template นี้ถูกออกแบบมาเพื่อให้เริ่ม Project ใหม่ได้เร็วขึ้น โดยไม่ต้องวางโครงสร้างซ้ำทุกครั้ง และช่วยให้แต่ละ Project มีมาตรฐานเดียวกันตั้งแต่วันแรก

---

## จุดเด่นของ Template นี้

* ใช้ Go + Fiber
* วางโครงแบบ Hexagonal Architecture
* แยกโค้ดตาม Module / Domain
* มีตัวอย่าง `user` module
* มี `BaseEntity` สำหรับ field กลาง เช่น `createdAt`, `updatedAt`, `deletedAt`
* มี Shared Error และ Shared Response กลาง
* มี Makefile สำหรับคำสั่งที่ใช้บ่อย
* มี Script สำหรับสร้าง Module ใหม่
* มี Script สำหรับเปลี่ยนชื่อ Project หลังสร้างจาก Template
* มี GitHub Actions CI
* มี Health Check endpoint
* พร้อมต่อยอดไปใช้ Database, Docker, Auth, Queue หรือ External Service ในอนาคต

---

## Requirements

เครื่องที่ใช้ควรมีสิ่งเหล่านี้:

```bash
go version
git --version
make --version
```

แนะนำให้ใช้:

```txt
Go 1.26+
Git
Make
```

ถ้าใช้ GitHub CLI สำหรับสร้าง repo จาก terminal ให้ติดตั้งเพิ่ม:

```bash
gh --version
```

---

## วิธี Run Project

Clone project:

```bash
git clone https://github.com/panuwat39/go-hexagonal-template.git
cd go-hexagonal-template
```

จัดการ dependency:

```bash
go mod tidy
```

ตรวจสอบ project:

```bash
make check
```

Run server:

```bash
make run
```

ทดสอบ Health Check:

```bash
curl http://localhost:8080/health
```

ตัวอย่าง response:

```json
{
  "env": "local",
  "status": "ok"
}
```

---

## วิธีใช้ Template นี้กับ Project ใหม่

หลังจากสร้าง repository ใหม่จาก template นี้แล้ว ให้รันคำสั่ง:

```bash
make init-project module=github.com/your-username/your-project name=your-project
```

ตัวอย่าง:

```bash
make init-project module=github.com/panuwat39/shop-api name=shop-api
```

จากนั้นรัน:

```bash
make tidy
make check
make run
```

ทดสอบ:

```bash
curl http://localhost:8080/health
```

---

## Project Structure

โครงสร้างหลักของ Project:

```txt
cmd/
  api/
    main.go

internal/
  bootstrap/
    config.go

  modules/
    user/

  platform/
    database/
    httpserver/
    logger/
    cache/
    queue/

  shared/
    domain/
      entity/
        base_entity.go
    errors/
      app_error.go
    response/
      http_status.go
      json_response.go

migrations/
deployments/
docs/
scripts/
.github/
  workflows/
    ci.yml

.env.example
Makefile
README.md
go.mod
go.sum
```

คำอธิบายแต่ละส่วน:

```txt
cmd/api
= entrypoint ของ application

internal/bootstrap
= โหลด config และประกอบ dependency หลักของระบบ

internal/modules
= เก็บ business modules เช่น user, order, payment, inventory

internal/platform
= technical infrastructure เช่น database, cache, queue, logger

internal/shared
= ของกลางที่ใช้ร่วมกัน เช่น error, response, base entity

migrations
= database migration files

deployments
= deployment config เช่น Docker, Kubernetes, CI/CD config เพิ่มเติม

docs
= เอกสารของ project

scripts
= shell scripts สำหรับช่วย automate งานซ้ำๆ

.github/workflows
= GitHub Actions workflow
```

---

## แนวคิด Hexagonal Architecture

Hexagonal Architecture คือแนวคิดที่แยก **Business Logic** ออกจาก Framework, Database และ External Service

แนว dependency ที่ต้องการคือ:

```txt
adapter -> application -> domain
```

ตัวอย่าง flow ที่ถูกต้อง:

```txt
HTTP Handler
  -> Use Case
  -> Repository Interface
  <- Repository Adapter
```

สิ่งที่ควรหลีกเลี่ยง:

```txt
Use Case -> Fiber โดยตรง
Use Case -> PostgreSQL client โดยตรง
Domain -> ORM Model
Module A -> Repository ของ Module B โดยตรง
```

เป้าหมายคือให้แกนกลางของระบบ เช่น `domain` และ `application` ไม่ผูกกับ framework หรือ database มากเกินไป

---

## Module Structure

แต่ละ Module ควรมีโครงสร้างแบบนี้:

```txt
internal/modules/<module>/
  domain/
    entity/
    valueobject/
    repository/
    service/

  application/
    command/
    query/
    usecase/
    port/

  adapter/
    inbound/
      http/
      consumer/
    outbound/
      persistence/
      external/

  module.go
```

คำอธิบาย:

```txt
domain/entity
= entity หลักของ domain นั้น

domain/valueobject
= value object เช่น Email, Money, Address

domain/repository
= repository interface ที่ domain/application ต้องการ

domain/service
= domain service สำหรับ business rule ที่ไม่เหมาะจะอยู่ใน entity เดียว

application/command
= input model สำหรับ operation ที่เปลี่ยนข้อมูล เช่น create, update, delete

application/query
= input model สำหรับ operation ที่อ่านข้อมูล

application/usecase
= use case ของระบบ เช่น CreateUserUseCase

application/port
= interface ที่ use case ต้องการจากโลกภายนอก

adapter/inbound/http
= HTTP route, handler, request, response

adapter/inbound/consumer
= queue consumer หรือ event consumer

adapter/outbound/persistence
= database implementation ของ repository

adapter/outbound/external
= adapter สำหรับเรียก external service

module.go
= จุดประกอบ dependency ของ module นั้น
```

---

## Sample User Module

Template นี้มีตัวอย่าง `user` module เพื่อให้ดู pattern การเขียนจริง

Flow ปัจจุบัน:

```txt
POST /users
  -> user route
  -> user handler
  -> create user use case
  -> user entity
  -> user repository interface
  -> in-memory user repository
```

โครงสร้างหลักของ user module:

```txt
internal/modules/user/
  domain/
    entity/
      user.go
    repository/
      user_repository.go

  application/
    command/
      create_user_command.go
    usecase/
      create_user_usecase.go

  adapter/
    inbound/
      http/
        user_route.go
        user_handler.go
        user_request.go
        user_response.go
        user_error_mapper.go
    outbound/
      persistence/
        in_memory_user_repository.go

  module.go
```

---

## Endpoint ที่มีตอนนี้

### Health Check

```txt
GET /health
```

ตัวอย่าง:

```bash
curl http://localhost:8080/health
```

ตัวอย่าง response:

```json
{
  "env": "local",
  "status": "ok"
}
```

---

### Create User

```txt
POST /users
```

ตัวอย่าง request:

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","name":"John"}'
```

ตัวอย่าง success response:

```json
{
  "id": "generated-id",
  "email": "john@example.com",
  "name": "John",
  "createdAt": "2026-07-14T10:00:00Z",
  "updatedAt": "2026-07-14T10:00:00Z"
}
```

ตัวอย่าง error response:

```json
{
  "code": "CONFLICT",
  "message": "email already exists"
}
```

---

## BaseEntity

`BaseEntity` อยู่ที่:

```txt
internal/shared/domain/entity/base_entity.go
```

ใช้สำหรับ field กลางที่ entity หลายตัวต้องใช้ร่วมกัน เช่น:

```txt
createdAt
updatedAt
deletedAt
createdBy
updatedBy
deletedBy
```

แนวคิดการใช้งาน:

```go
type User struct {
    sharedentity.BaseEntity

    id    string
    email string
    name  string
}
```

การตรวจว่า record ถูกลบแล้วหรือยัง ควรใช้ method:

```go
user.IsDeleted()
```

ไม่จำเป็นต้องมี field `isDeleted` ซ้ำใน domain เพราะสามารถคำนวณจาก `deletedAt` ได้

```txt
deletedAt != nil หมายถึงถูก soft delete แล้ว
deletedAt == nil หมายถึงยังไม่ถูกลบ
```

ข้อควรระวัง:

```txt
ไม่ควรใส่ ORM tag ลงใน domain entity
เช่น gorm tag, db tag หรือ json tag ที่เกี่ยวกับ database mapping
```

เพราะ domain ไม่ควรผูกกับ database หรือ ORM โดยตรง

---

## Shared Error

Shared error อยู่ที่:

```txt
internal/shared/errors/app_error.go
```

ใช้กำหนด error code กลาง เช่น:

```txt
BAD_REQUEST
UNAUTHORIZED
FORBIDDEN
NOT_FOUND
CONFLICT
INTERNAL
```

ตัวอย่างการใช้งาน:

```go
return apperror.New(
    apperror.CodeBadRequest,
    "invalid request body",
)
```

ถ้าต้องการ wrap error ภายใน:

```go
return apperror.Wrap(
    apperror.CodeInternal,
    "internal server error",
    err,
)
```

---

## Shared Response

Shared response อยู่ที่:

```txt
internal/shared/response/
  http_status.go
  json_response.go
```

ใช้สำหรับกำหนด response format กลางของระบบ

ตัวอย่าง success response:

```go
return response.Created(c, body)
```

ตัวอย่าง error response:

```go
return response.Error(c, err)
```

รูปแบบ error response มาตรฐาน:

```json
{
  "code": "BAD_REQUEST",
  "message": "invalid request body"
}
```

HTTP status code กลางอยู่ใน:

```txt
internal/shared/response/http_status.go
```

เช่น:

```txt
StatusOK                  = 200
StatusCreated             = 201
StatusNoContent           = 204
StatusBadRequest          = 400
StatusUnauthorized        = 401
StatusForbidden           = 403
StatusNotFound            = 404
StatusConflict            = 409
StatusUnprocessableEntity = 422
StatusInternalServerError = 500
```

---

## Environment Variables

ตัวอย่าง env อยู่ที่:

```txt
.env.example
```

ค่าปัจจุบัน:

```env
APP_NAME=go-hexagonal-template
APP_ENV=local
HTTP_PORT=8080

DATABASE_URL=
REDIS_URL=
```

ตัวอย่าง run ด้วย env:

```bash
APP_NAME=shop-api APP_ENV=local HTTP_PORT=9090 make run
```

แล้วทดสอบ:

```bash
curl http://localhost:9090/health
```

ข้อควรระวัง:

```txt
ห้าม commit secret จริงลง repository
เช่น password, token, private key, API key
```

ให้ใช้ environment variable หรือ secret manager แทน

---

## Make Commands

Template นี้มี `Makefile` เพื่อให้ใช้คำสั่งสั้นและจำง่าย

---

### Run Server

```bash
make run
```

ใช้สำหรับ run API server:

```bash
go run ./cmd/api
```

---

### Development Run

```bash
make dev
```

ใช้ run ผ่าน script:

```bash
./scripts/dev.sh
```

เหมาะสำหรับต่อยอดในอนาคต เช่น load `.env`, hot reload หรือ local setup เพิ่มเติม

---

### Run Test

```bash
make test
```

รัน test ทั้ง project:

```bash
go test ./...
```

---

### Run Test พร้อม Race Detector และ Coverage

```bash
make test-race
```

รัน:

```bash
go test -race -cover ./...
```

เหมาะสำหรับตรวจปัญหา concurrency เช่น goroutine หลายตัวอ่าน/เขียนข้อมูลพร้อมกัน

---

### Build Binary

```bash
make build
```

build binary ไปที่:

```txt
bin/app
```

ถ้าต้องการกำหนดชื่อ binary:

```bash
make build APP_NAME=shop-api
```

จะได้:

```txt
bin/shop-api
```

---

### Format Code

```bash
make fmt
```

รัน:

```bash
gofmt -w .
```

ใช้จัด format Go code ทั้ง project

---

### Check Format

```bash
make fmt-check
```

ตรวจว่า code format แล้วหรือยัง แต่ไม่แก้ไฟล์ให้

เหมาะกับการใช้ใน CI

---

### Go Vet

```bash
make vet
```

รัน:

```bash
go vet ./...
```

ใช้ตรวจ bug pattern ที่ compiler อาจไม่จับ

---

### Go Mod Tidy

```bash
make tidy
```

รัน:

```bash
go mod tidy
```

ใช้จัด dependency ใน `go.mod` และ `go.sum`

---

### Run All Checks

```bash
make check
```

ใช้ตรวจทุกอย่างก่อน commit หรือ push

โดยรวมจะตรวจ:

```txt
go mod tidy
gofmt
go vet
go test -race -cover
go build
```

แนะนำให้รันก่อน push ทุกครั้ง

---

### Vulnerability Check

```bash
make vulncheck
```

ใช้ตรวจ dependency vulnerability ด้วย `govulncheck`

ถ้าเครื่องยังไม่มี `govulncheck` script จะพยายามติดตั้งให้

---

### Create New Module

```bash
make new-module name=user
```

ตัวอย่าง:

```bash
make new-module name=order
make new-module name=payment
make new-module name=inventory
```

ชื่อ module ต้องเป็นตัวพิมพ์เล็กและเป็น alphanumeric เท่านั้น

ตัวอย่างชื่อที่ถูก:

```txt
user
order
payment
inventory
product
```

---

### Init Project จาก Template

```bash
make init-project module=github.com/your-username/your-project name=your-project
```

ตัวอย่าง:

```bash
make init-project module=github.com/panuwat39/shop-api name=shop-api
```

ใช้หลังจากสร้าง repository ใหม่จาก template นี้

---

### Clean Build Output

```bash
make clean
```

ลบ folder:

```txt
bin/
```

---

## Scripts

ไฟล์ script อยู่ใน:

```txt
scripts/
```

---

### check.sh

```txt
scripts/check.sh
```

ใช้โดย:

```bash
make check
```

หน้าที่หลัก:

```txt
ตรวจ go.mod / go.sum
ตรวจ gofmt
รัน go vet
รัน go test -race -cover
รัน go build
```

---

### dev.sh

```txt
scripts/dev.sh
```

ใช้โดย:

```bash
make dev
```

ตอนนี้ใช้ run API local

---

### init-project.sh

```txt
scripts/init-project.sh
```

ใช้เปลี่ยน module path หลังสร้าง project ใหม่จาก template

---

### new-module.sh

```txt
scripts/new-module.sh
```

ใช้สร้าง module ใหม่ตามโครง Hexagonal

---

## GitHub Actions CI

Workflow อยู่ที่:

```txt
.github/workflows/ci.yml
```

CI จะรันเมื่อ:

```txt
push เข้า main
เปิด pull request
กด run แบบ manual ผ่าน workflow_dispatch
```

Jobs หลัก:

```txt
Verify
Vulnerability Check
```

Verify จะรัน:

```bash
make check
```

Vulnerability Check จะรัน:

```bash
govulncheck ./...
```

เป้าหมายคือให้แน่ใจว่า code format ถูก, test ผ่าน, build ผ่าน และไม่มี dependency vulnerability สำคัญก่อนนำไปใช้ต่อ

---

## Recommended Development Flow

ก่อนเริ่มงาน:

```bash
git pull
```

ระหว่างพัฒนา:

```bash
make fmt
make test
```

ก่อน commit:

```bash
make check
```

commit:

```bash
git add .
git commit -m "Your commit message"
git push
```

หลัง push ให้ตรวจ GitHub Actions ว่า CI เป็นสีเขียว

---

## Rule สำคัญของ Module

Module หนึ่งไม่ควรเรียก repository ของอีก module โดยตรง

ไม่แนะนำ:

```txt
order usecase -> user repository
```

ควรทำ:

```txt
order usecase -> order port -> adapter -> user module/service
```

เหตุผลคือช่วยให้ boundary ระหว่าง module ชัดเจน และลด coupling ระหว่าง domain

---

## วิธีเพิ่ม Module ใหม่

ตัวอย่างเพิ่ม `order` module:

```bash
make new-module name=order
```

จากนั้นเติม code ตาม pattern:

```txt
domain/entity
domain/repository
application/command
application/usecase
adapter/inbound/http
adapter/outbound/persistence
module.go
```

ถ้า module ต้องใช้ข้อมูลจาก module อื่น ให้ประกาศ port ใน module ตัวเองก่อน แล้วค่อยให้ adapter ไปเชื่อมกับ module ภายนอก

---

## หมายเหตุเรื่อง Database

ตอนนี้ `user` module ใช้ `InMemoryUserRepository` เพื่อให้ template เบาและเข้าใจง่าย

ถ้าจะเพิ่ม PostgreSQL, MySQL, MongoDB หรือ Redis ให้เพิ่ม implementation ที่:

```txt
internal/modules/<module>/adapter/outbound/persistence/
```

ตัวอย่าง:

```txt
postgres_user_repository.go
user_mapper.go
```

ไม่ควรเขียน database logic ใน use case หรือ domain entity โดยตรง

---

## หมายเหตุเรื่อง Framework

Fiber ควรอยู่ใน layer:

```txt
adapter/inbound/http
```

ไม่ควรให้ `domain` หรือ `application` รู้จัก Fiber

ตัวอย่างที่ควรหลีกเลี่ยง:

```go
func (uc *CreateUserUseCase) Execute(c fiber.Ctx) {}
```

ที่ถูกควรเป็น:

```go
func (uc *CreateUserUseCase) Execute(ctx context.Context, cmd CreateUserCommand) {}
```

เพราะ use case ควรไม่ผูกกับ HTTP framework

---

## สิ่งที่แนะนำให้ทำต่อ

หลังจาก template นี้พร้อมใช้งานแล้ว สิ่งที่สามารถเพิ่มต่อได้:

```txt
Dockerfile
docker-compose สำหรับ local development
PostgreSQL adapter
Database migration tool
Request validation layer
Authentication module
Authorization middleware
Logger middleware
Request ID middleware
Graceful shutdown ที่ละเอียดขึ้น
Integration test
```

---

## สรุป

Template นี้เหมาะสำหรับเริ่มต้น Go Backend ที่ต้องการโครงสร้างชัดเจนและต่อยอดได้ในระยะยาว

แนวคิดหลักคือ:

```txt
แยก business logic ออกจาก framework และ database
แยก module ตาม domain
ใช้ shared error/response กลาง
ใช้ Makefile และ scripts ลดงานซ้ำ
ใช้ CI ตรวจคุณภาพก่อน merge/push
```

เมื่อเริ่ม project ใหม่จาก template นี้ ควรทำตามลำดับ:

```txt
1. สร้าง repository จาก template
2. รัน make init-project
3. รัน make check
4. รัน make run
5. เริ่มเพิ่ม business module จริง
```
