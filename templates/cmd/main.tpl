package {{.PackageName}}

import (
	"fmt"
	"go-code-generator/config"
	"html"
	"log"
	"net/http"
)

func main() {
	settings := config.New()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from {{.ServiceName}} with endpoint %q!", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", settings.HttpPort), nil))
}