# Project: gRPC Student Service (Golang)

This project demonstrates a basic gRPC service in Go with a GetStudent RPC method.

---

## Step 1 — Install Requirements

Install required tools:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure you have:

- Go installed
- protoc (Protocol Buffers compiler)
- GOPATH/bin in PATH

Check:

```bash
go version
protoc --version
```

---

## Step 2 — Create Project

```bash
  mkdir grpc-student
  cd grpc-student
  go mod init grpc-student
```

---

## Step 3 — Create Project Structure

```bash
mkdir proto
mkdir server
mkdir client
````

### Structure of the project:

```text
grpc-student/
│
├── proto/
│   └── student.proto
├── server/
│   └── server.go
├── client/
│   └── client.go
├── go.mod
```

---
## Step 4 — Create Proto File
Create a proto file in the `proto` folder:
```text
proto/student.proto
`````

```proto
syntax = "proto3";

package student;

option go_package = "./studentpb";

service StudentService {
  rpc GetStudent (StudentRequest) returns (StudentResponse);
}

message StudentRequest {
  int32 id = 1;
}

message StudentResponse {
  int32 id = 1;
  string name = 2;
  string major = 3;
  string email = 4;
}
```

---
## Step 5 — Generate gRPC Code
Run the following command to generate the Go code from the proto file:
```bash
protoc --go_out=. --go-grpc_out=. proto/student.proto
```
### Explanation:
This command generates Go code from the `student.proto` file.
```text
| Part                | Meaning                        |
| ------------------- | ------------------------------ |
| protoc              | Protocol Buffer compiler       |
| --go_out=.          | Generate message code (.pb.go) |
| --go-grpc_out=.     | Generate gRPC service code     |
| proto/student.proto | Input proto file               |
```

This will generate:
```text
studentpb/
    student.pb.go
    student_grpc.pb.go
```

```text
| File               | Purpose                              |
|--------------------|------------------------------------- |
| student.pb.go      | Message structs (Request/Response)   |
| student_grpc.pb.go | gRPC client and server interfaces    |
```

---
## Step 6 — Create Server
File: `server/server.go`

```go
package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {

	log.Printf("Received request for student ID: %d", req.Id)

	// Mock data
	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
```

Run server terminal1:
```bash
go mod tidy // to update go.mod
go run server/server.go // to run server
```
Server will run on:
`localhost:50051`

---
Step 7 — Create Client
File: `client/client.go`
```go
package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 101,
	})

	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
}
```

Run client terminal2:
```bash
go run client/client.go
```

Expected output:
```text
Student Info:
ID: 101
Name: Alice Johnson
Major: Computer Science
Email: alice@university.com
```

---
## 💡Summary: How gRPC Works (Flow)
1. Write proto file
2. Generate Go code
3. Implement server
4. Start server
5. Implement client
6. Client calls GetStudent()
7. Server returns student data

### Architecture:
`Client → Stub → gRPC → Server Stub → Server Method`

Final Project Structure
```text
grpc-student/
│
├── proto/
│   └── student.proto
│
├── studentpb/
│   ├── student.pb.go
│   └── student_grpc.pb.go
│
├── server/
│   └── server.go
│
├── client/
│   └── client.go
│
├── go.mod
└── go.sum
```

---
🧪Assignments:
---
### 📌 Task 1: Add ListStudents RPC method to the `student.proto`
```proto
rpc ListStudents (Empty) returns (StudentListResponse);
```
Add messages:
```proto
message Empty {}

message StudentListResponse {
  repeated StudentResponse student = 1;
}
```
Students must:
- Update proto
- Regenerate code
- Implement a server method and return a list of students
- Update client to call ListStudents

---
### 📌 Task 2: Add New Field
Add `phone` in StudentResponse:
```proto
  string phone = 5;
```
Students must:
- Regenerate proto
- Update server response
- Update client output


---
## Submission - Due: Monday 30, 23:59
- Push project to your Git repository (Repository name: `grpc-student-รหัสนักศึกษา`)
- Copy GitHub link to a text file: `grpc-student-รหัสนักศึกษา.txt`
- Submit at: `https://courses.cs.tu.ac.th`

Note: `รหัสนักศึกษา` is your student ID

## Additional Feature
- Added faculty and year fields to StudentResponse
- Updated server and client to support new fields