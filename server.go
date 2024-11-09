package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/faizallmaullana/al-quran-indonesia/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"golang.org/toolchain/src/math/rand"
)

// Start the server on the specified port.
func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	r := gin.New()

	rand.Seed(time.Now().UnixNano())
	port := rand.Intn(10000) + 1000
	address := fmt.Sprintf("http://localhost:%d", port)

	middleware.ActivateMiddleware(r)

	mu.Lock()
	servers = append(servers, address)
	mu.Unlock()

	r.Use(static.Serve("/", static.LocalFile("./alQuranFrontEnd/dist", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./alQuranFrontEnd/dist/index.html")
	})

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Printf("Server failed on port %d: %v", port, err)
	}
}
