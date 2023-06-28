package mainx

import (
	"github.com/Natchalit/account/src/services/logx"
	"github.com/joho/godotenv"
)

func Init() {
	logx.Init()
	_ = godotenv.Load(".env.local")
	_ = godotenv.Load(".env")

}
