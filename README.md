# Marvel API

Обертка для Marvel API

```sh
go get -u github.com/AndrewMislyuk/marvel-api/marvel
```

Import

```go
import "github.com/AndrewMislyuk/marvel-api/marvel"
```

Usage

```go
s := marvel.New(publicAPIKey, privateAPIKey, time.Second*10)
result, err := s.GetComicsList()
if err != nil {
	log.Fatal(err)
}

fmt.Println(result)

result, err = s.GetComic(82967)
if err != nil {
	log.Fatal(err)
}

fmt.Println(result)
```
