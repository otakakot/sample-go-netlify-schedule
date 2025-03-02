package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/fujiwara/ridge"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("X-Api-Key: " + r.Header.Get("X-Api-Key") + ", API_KEY: " + os.Getenv("API_KEY"))

		if r.Header.Get("X-Api-Key") != os.Getenv("API_KEY") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}

		ctx, _ := lambdacontext.FromContext(r.Context())

		slog.Info(fmt.Sprintf("Headers: %v", r.Header))

		w.Write([]byte(ctx.AwsRequestID))
	})

	ridge.Run(":8080", "/", mux)
}
