# ent-sqlite
this is an ent sqlite driver adapter, because ent must use foreignkey.

# install
```bash
go get github.com/dstgo/ent-sqlite@latest
```

# example
```go
import (
    _ "github.com/dstgo/ent-sqlite"
)


func main(){
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatal(err)
    }
}
```