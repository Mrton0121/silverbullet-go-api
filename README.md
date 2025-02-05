# Go API wrapper for [Silverbullet](https://silverbullet.md/)

This is a small API wrapper that implements the basic CRUD methods.

You can see examples in the main.go

This package is used in [silverbullet-api-gateway](https://github.com/Mrton0121/silverbullet-api-gateway)

# Examples
```go
// Create client
url := os.Getenv("SB_URL") //use any way to get he protected data
token := os.Getenv("SB_TOKEN")
sbclient := sbapi.NewClient(url, token)
```

```go
// GET method gets the page text in markdown
main, err := sbclient.Get("main.md")
if err != nil {
  panic(err)
}

fmt.PrintLn(main)
```

```go
// PUT method creates a new page with given data
new, err := sbclient.Put("ApiTest/tester2.md", "# Hello")
if err != nil {
  panic(err)
}

fmt.PrintLn(new)
```

```go
// APPEND method is based on GET and PUT methods, adds data
// in a new line to the already existing page
append, err := sbclient.Append("ApiTest/tester.md", "# World", "\n")
if err != nil {
  panic(err)
}

fmt.PrintLn(append)
```

```go
// DELETE method is removing the given file. If that's the
// only file in the folder, the folder will automatically get deleted
del, err := sbclient.Delete("ApiTest/tester2.md")
if err != nil {
  panic(err)
}

fmt.PrintLn(del)
```

Author: Márton Gombócz (Mrton0121)