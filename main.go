package main
import (
    "net/http"
    "golang.org/x/net/webdav"
)
func main() {
    http.ListenAndServe(":8089", &webdav.Handler{
        FileSystem: webdav.Dir("~/"),
        LockSystem: webdav.NewMemLS(),
    })
}
