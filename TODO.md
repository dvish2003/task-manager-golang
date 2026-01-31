task-manager/
 ├── backend/
 │    ├── main.go
 │    ├── models/
 │    │     ├── user.go
 │    │     └── task.go
 │    ├── routes/
 │    │     └── routes.go
 │    └── config/
 │          └── db.go
 └── frontend/
      └── react app




mkdir backend && cd backend
go mod init task-manager



go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo



if go project clone install dependencies type 'go mod tidy'
go project install air its like nodemon in nodejs   and cd project => type 'air'