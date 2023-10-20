package main


import(
	"os"
	"github.com/f-godoy/produtos-go/infrastructure/db"
	"github.com/f-godoy/produtos-go/application/grpc"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main(){
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}