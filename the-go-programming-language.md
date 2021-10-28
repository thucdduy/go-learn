# 1. Tutorial

### 1.1. Hello, Worlds

### 1.2. Command-line Arguments

### 1.3. Finding Duplicate Lines 


- Hàm bufio.NewScanner(f) nhận vào interface `io.Reader`

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

Vậy bất kỳ kiểu dữ liệu nào có func `Read(p []byte) (n int, err error)` đều có thể  truyên fvào bufio.NewScanner(f)

- Cách sử dụng hàm `fmt.Fprintf`:

```
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)

```

ví dụ:
```
fmt.Fprintf(os.Stderr, "dup2: %v\n",err)
```

### 1.4 Animated Gifs

### 1.5 Fetching a URL

- `for _, url := os.Args[1:]`   duyệt qua các arguments
- `http.Get(url)` để gửi GET request lên web server
- `ioutil.ReadAll(resp.Body)` để đọc dữ liệu response body.
- `resp.Body.Close()` đóng stream sau khi đọc.
- `fmt.Fprintf(os.Stderr, "fetch: %v\n", err)` Xuất lỗi ra stderr.
- `io.Copy(dst, src)` có thể copy trực tiếp ra stdout mà không cần buffer trung gian ReadAll.

### 1.6 Fetching URLs Concurrently

- `start:=time.Now()` & `time.Since(start).Seconds())` tính thời gian kể từ `start`.
- `nbytes, err := io.Copy(ioutil.Discard, resp.Body)` discard data và trả về số bytes.

### 1.7 A Web Server

- `http.HandleFunc(urlPattern, handler` và `func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}` để xử lý http requests 
- `log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))` vai trò web server.
- `var mu sync.Muxex`, `mu.Lock()` và `mu.Unlock` dùng khi muốn giới hạn 1 go routine duy nhất được phép truy cập vùng nhớ tại một thời điểm.
-  `r.Method`, `r.URL`, `r.Proto`,... 
- `k, v := range r.Header`, `k, v = `, `err := r.ParseForm()`, `k,v := range r.Form`...
 

