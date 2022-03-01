package main

import(
  "net/http"
  "io/ioutil"
  "log"
)

type User struct {
  Name     string
  Birthday string
}

func main() {
  http.HandleFunc("/api/users", HandleGetUsersRequest)
  http.ListenAndServe(":8080", nil)
}

func HandleGetUsersRequest(w http.ResponseWriter, r *http.Request) {
  // users: []bytes
  users, err := ioutil.ReadFile("./sampleJson/users.json")
    if err != nil {
        log.Fatal(err)
    }
  w.Header().Set("Content-Type", "application/json")
  w.Write(users)
}
