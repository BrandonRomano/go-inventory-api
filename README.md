# Inventory API

This is just a dummy API that I'll be writing to learn go

## Setting Up

### Go Packages

Here's a list of dependencies.  I haven't quite looked into any package managers yet, so I'm just dumping them into this list.  You can run `go get dependency/url` in your project to install the dependencies.

```
github.com/go-sql-driver/mysql
github.com/gorilla/mux
```

### Originator

You'll also need to have [originator](https://github.com/DigitalCitadel/originator) installed.

After originator is intalled, update the config files to match your environment, run the migrations, and you should be good to go.

### database/database.go file

You'll also have to update the `const` values in database.go.

I'm sure there's a better way to do this, but I haven't looked to deep yet.
