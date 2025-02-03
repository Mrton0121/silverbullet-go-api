# Go API wrapper for [Silverbullet](https://silverbullet.md/)

This is a small API wrapper that implements the basic CRUD methods.

You can see examples in the main.go

# Examples
```go
// Create client
url := os.Getenv("SB_URL")
token := os.Getenv("SB_TOKEN")
sbclient := sbapi.NewClient(url, token)
```

```go
// GET method gets the page text in markdown
main, err := sbclient.Get("main.md")
if err != nil {
  panic(err)
}
```

```go
// PUT method creates a new page with given data
new, err := sbclient.Put("ApiTest/tester2.md", "# Hello")
if err != nil {
  panic(err)
}
```

```go
// APPEND method is based on GET and PUT methods, adds data in a new line to the already existing page
append, err := sbclient.Append("ApiTest/tester.md", "# World", "\n")
if err != nil {
  panic(err)
}
```

```go
// DELETE method is removing the given file. If that's the only file in the folder, the folder will automatically get deleted
del, err := sbclient.Delete("ApiTest/tester2.md")
if err != nil {
  panic(err)
}
```

Author: Márton Gombócz (Mrton0121)