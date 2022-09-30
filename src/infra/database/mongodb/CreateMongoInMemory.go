package mongodb

import (
	"fmt"
	"os"

	"github.com/tryvium-travels/memongo"
)

func CreateMongoTemp() *memongo.Server {
	opts := &memongo.Options{
		MongoVersion: "5.0.5",
		// Port:         45646,
	}
	// if runtime.GOARCH == "arm64" {
	// 	if runtime.GOOS == "darwin" {
	// 		// Only set the custom url as workaround for arm64 macs
	// 		opts.DownloadURL = "https://fastdl.mongodb.org/osx/mongodb-macos-x86_64-5.0.5.tgz"
	// 	}
	// }
	mongoServer, err := memongo.StartWithOptions(opts)

	if err != nil {
		panic(err)
	}

	os.Setenv("MONGO_URL", mongoServer.URI())
	// os.Setenv("MONGO_URL", "mongodb://localhost:27017/")
	os.Setenv("DATABASE", memongo.RandomDatabase())
	// os.Setenv("DATABASE", "test")

	fmt.Println("\n\n", "MONGO_URL:", os.Getenv("MONGO_URL"))
	fmt.Println("DATABASE: ", os.Getenv("DATABASE"))

	return mongoServer
}
