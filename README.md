A Golang/webasm/SAT solver based sudoku solver


To run locally you can serve with:
``` shell
echo "install goexec"
go get -v -u github.com/shurcooL/goexec
echo "serve"
goexec http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
```


