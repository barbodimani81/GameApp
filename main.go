package main

import (
	"Game/repository/mysql"
	"Game/service/userservice"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/register", userRegisterHandler)

	log.Println("server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(writer, `{"error", "invalid method"}`)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)

	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))
		return
	}
	writer.Write([]byte(`{"message": "user created"}`))
}

func healthCheckHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, `{"message": "all good!"}`)
}
