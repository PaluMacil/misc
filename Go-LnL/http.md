# Go HTTP!

[Official Docs for net/http package](https://golang.org/pkg/net/http/)



Highlights from the official docs are below. The official docs cover everything in all the depth you need since Go encourages such great documentation.

```
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```

Handle and HandleFunc add handlers to DefaultServeMux:

```
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

[&#x2190; Back to Go Lunch and Learn](../README.md)

[&#x2190; Back to Project List](../../README.md)