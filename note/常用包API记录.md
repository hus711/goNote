### 常用包API记录

------

#### IO包

------

​	import "io/ioutil"

Variables

```
var Discard io.Writer = devNull(0)
```

Discard是一个io.Writer接口，对它的所有Write调用都会无实际操作的成功返回。

------

​	import "io"

func [Copy](https://github.com/golang/go/blob/master/src/io/io.go?name=release#341)

```
func Copy(dst Writer, src Reader) (written int64, err error)
```

将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。

------

#### time包

`import "time"`

time包提供了时间的显示和测量用的函数，Time代表一个纳秒精度的时间点。日历的计算采用的是公历。

------

func [Now](https://github.com/golang/go/blob/master/src/time/time.go?name=release#784)

```
func Now() Time
```

Now返回当前本地时间。

------

func [Since](https://github.com/golang/go/blob/master/src/time/time.go?name=release#646)

```
func Since(t Time) Duration
```

Since返回从t到现在经过的时间，等价于time.Now().Sub(t)。

------

#### net包

`import "net"`

net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。

------

##### `import "net/http"`

http包提供了HTTP客户端和服务端的实现。Get、Head、Post和PostForm函数发出HTTP/ HTTPS请求

type [Header](https://github.com/golang/go/blob/master/src/net/http/header.go?name=release#19)

```
type Header map[string][]string
```

Header代表HTTP头域的键值对。



type [ResponseWriter](https://github.com/golang/go/blob/master/src/net/http/server.go?name=release#51)

```
type ResponseWriter interface {
    // Header返回一个Header类型值，该值会被WriteHeader方法发送。
    // 在调用WriteHeader或Write方法后再改变该对象是没有意义的。
    Header() Header
    // WriteHeader该方法发送HTTP回复的头域和状态码。
    // 如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK)
    // WriterHeader的显式调用主要用于发送错误码。
    WriteHeader(int)
    // Write向连接中写入作为HTTP的一部分回复的数据。
    // 如果被调用时还未调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
    // 如果Header中没有"Content-Type"键，
    // 本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
    Write([]byte) (int, error)
}
```

ResponseWriter接口被HTTP处理器用于构造HTTP回复。



type [Request](https://github.com/golang/go/blob/master/src/net/http/request.go?name=release#76)

```
type Request struct {
    // Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
    Method string
    // URL在服务端表示被请求的URI，在客户端表示要访问的URL。
    //
    // 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
    // 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串。
    // （参见RFC 2616, Section 5.1.2）
    //
    // 在客户端，URL的Host字段指定了要连接的服务器，
    // 而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值。
    URL *url.URL
    // 接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0
    // Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
    //	accept-encoding: gzip, deflate
    //	Accept-Language: en-us
    //	Connection: keep-alive
    // 则：
    //	Header = map[string][]string{
    //		"Accept-Encoding": {"gzip, deflate"},
    //		"Accept-Language": {"en-us"},
    //		"Connection": {"keep-alive"},
    //	}
    // HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
    // 在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。
    Header Header
    // Body是请求的主体。
    //
    // 在客户端，如果Body是nil表示该请求没有主体买入GET请求。
    // Client的Transport字段会负责调用Body的Close方法。
    //
    // 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF。
    // Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段。
    Body io.ReadCloser
    // ContentLength记录相关内容的长度。
    // 如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据。
    // 在客户端，如果Body非nil而该字段为0，表示不知道Body的长度。
    ContentLength int64
    // TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
    // 本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码。
    TransferEncoding []string
    // Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。
    Close bool
    // 在服务端，Host指定URL会在其上寻找资源的主机。
    // 根据RFC 2616，该值可以是Host头的值，或者URL自身提供的主机名。
    // Host的格式可以是"host:port"。
    //
    // 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
    // 如过该字段为""，Request.Write方法会使用URL字段的Host。
    Host string
    // Form是解析好的表单数据，包括URL字段的query参数(get参数)和POST或PUT的表单数据。
    // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
    Form url.Values
    // PostForm是解析好的POST或PUT的表单数据。
    // 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
    PostForm url.Values
    // MultipartForm是解析好的多部件表单，包括上传的文件。
    // 本字段只有在调用ParseMultipartForm后才有效。
    // 在客户端，会忽略请求中的本字段而使用Body替代。
    MultipartForm *multipart.Form
    // Trailer指定了会在请求主体之后发送的额外的头域。
    //
    // 在服务端，Trailer字段必须初始化为只有trailer键，所有键都对应nil值。
    // （客户端会声明哪些trailer会发送）
    // 在处理器从Body读取时，不能使用本字段。
    // 在从Body的读取返回EOF后，Trailer字段会被更新完毕并包含非nil的值。
    // （如果客户端发送了这些键值对），此时才可以访问本字段。
    //
    // 在客户端，Trail必须初始化为一个包含将要发送的键值对的映射。（值可以是nil或其终值）
    // ContentLength字段必须是0或-1，以启用"chunked"传输编码发送请求。
    // 在开始发送请求后，Trailer可以在读取请求主体期间被修改，
    // 一旦请求主体返回EOF，调用者就不可再修改Trailer。
    //
    // 很少有HTTP客户端、服务端或代理支持HTTP trailer。
    Trailer Header
    // RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志。
    // 本字段不是ReadRequest函数填写的，也没有定义格式。
    // 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址。
    // 客户端会忽略请求中的RemoteAddr字段。
    RemoteAddr string
    // RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
    // （参见RFC 2616, Section 5.1）
    // 一般应使用URI字段，在客户端设置请求的本字段会导致错误。
    RequestURI string
    // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
    // 本字段不是ReadRequest函数填写的。
    // 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
    // 客户端会忽略请求中的TLS字段。
    TLS *tls.ConnectionState
}
```

Request类型代表一个服务端接受到的或者客户端发送出去的HTTP请求。

Request各字段的意义和用途在服务端和客户端是不同的。除了字段本身上方文档，还可参见Request.Write方法和RoundTripper接口的文档。

------

func (*Request) [ParseForm](https://github.com/golang/go/blob/master/src/net/http/request.go?name=release#736)

```
func (r *Request) ParseForm() error
```

ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。

对于POST或PUT请求，ParseForm还会将body当作表单解析，并将结果既更新到r.PostForm也更新到r.Form。解析结果中，POST或PUT请求主体要优先于URL查询字符串（同名变量，主体的值在查询字符串的值前面）。

如果请求的主体的大小没有被MaxBytesReader函数设定限制，其大小默认限制为开头10MB。

ParseMultipartForm会自动调用ParseForm。重复调用本方法是无意义的。



DefaultClient是用于包函数Get、Head和Post的默认Client。

```
var DefaultServeMux = NewServeMux()
```

DefaultServeMux是用于Serve的默认ServeMux。



ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

```
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```

要管理服务端的行为，可以创建一个自定义的Server：

```
s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())
```

------

func [Get](https://github.com/golang/go/blob/master/src/net/http/client.go?name=release#251)

```
func Get(url string) (resp *Response, err error)
```

Get向指定的URL发出一个GET请求，如果回应的状态码如下，Get会在调用c.CheckRedirect后执行重定向：

```
301 (Moved Permanently)
302 (Found)
303 (See Other)
307 (Temporary Redirect)
```

如果c.CheckRedirect执行失败或存在HTTP协议错误时，本方法将返回该错误；如果回应的状态码不是2xx，本方法并不会返回错误。如果返回值err为nil，resp.Body总是非nil的，调用者应该在读取完resp.Body后关闭它。

Get是对包变量DefaultClient的Get方法的包装。

```go
//示例程序
res, err := http.Get("http://www.google.com/robots.txt")
if err != nil {
    log.Fatal(err)
}
robots, err := ioutil.ReadAll(res.Body)
res.Body.Close()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%s", robots)
```

------

func (*ServeMux) [HandleFunc](https://github.com/golang/go/blob/master/src/net/http/server.go?name=release#1554)

```
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

HandleFunc注册一个处理器函数handler和对应的模式pattern。

TODO 到底是什么模式？怎么进行URL匹配的？

------

type [Handler](https://github.com/golang/go/blob/master/src/net/http/server.go?name=release#45) [¶](http://godoc.ml/pkg/net_http.htm#pkg-index)

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务。

ServeHTTP应该将回复的头域和数据写入ResponseWriter接口然后返回。返回标志着该请求已经结束，HTTP服务端可以转移向该连接上的下一个请求。

------

func (*Server) [ListenAndServe](https://github.com/golang/go/blob/master/src/net/http/server.go?name=release#1679)

```
func (srv *Server) ListenAndServe() error
```

ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。如果srv.Addr为空字符串，会使用":http"。

------



##### `import "net/url"`

url包解析URL并实现了查询的逸码，参见[RFC 3986](http://tools.ietf.org/html/rfc3986)。

------



type [Values](https://github.com/golang/go/blob/master/src/net/url/url.go?name=release#482)

```
type Values map[string][]string
```

Values将建映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，Values的键是大小写敏感的。

func (Values) [Get](https://github.com/golang/go/blob/master/src/net/url/url.go?name=release#488)

```
func (v Values) Get(key string) string
```

Get会获取key对应的值集的第一个值。如果没有对应key的值集会返回空字符串。获取值集请直接用map。

func (Values) [Set](https://github.com/golang/go/blob/master/src/net/url/url.go?name=release#501)

```
func (v Values) Set(key, value string)
```

Set方法将key对应的值集设为只有value，它会替换掉已有的值集。

func (Values) [Add](https://github.com/golang/go/blob/master/src/net/url/url.go?name=release#507)

```
func (v Values) Add(key, value string)
```

Add将value添加到key关联的值集里原有的值的后面。

func (Values) [Del](https://github.com/golang/go/blob/master/src/net/url/url.go?name=release#512)

```
func (v Values) Del(key string)
```

Del删除key关联的值集。

func [QueryEscape](https://github.com/golang/go/blob/master/src/net/url/url.go?name=release#173)

```
func QueryEscape(s string) string
```

QueryEscape函数对s进行转码使之可以安全的用在URL查询里。（对字符串转码，然后再拿去访问）





------

#### image包

type [GIF](https://github.com/golang/go/blob/master/src/image/gif/reader.go?name=release#423)

```
type GIF struct {
    Image     []*image.Paletted // 连续的图像（Paletted调色板）
    Delay     []int             // 每一帧延迟时间，单位是0.01s
    LoopCount int               // 总的循环时间
}
```

GIF类型代表可能保存在GIF文件里的多幅图像。

func [DecodeAll](https://github.com/golang/go/blob/master/src/image/gif/reader.go?name=release#431)

```
func DecodeAll(r io.Reader) (*GIF, error)
```

函数从r中读取一个GIF格式文件；返回值中包含了连续的图帧和时间信息。

func [EncodeAll](https://github.com/golang/go/blob/master/src/image/gif/writer.go?name=release#262)

```
func EncodeAll(w io.Writer, g *GIF) error
```

函数将g中所有的图像按指定的每帧延迟和累计循环时间写入w中。



------

#### log包

`import "log"`

log包实现了简单的日志服务。本包定义了Logger类型，该类型提供了一些格式化输出的方法。



------

func (*Logger) [Fatal](https://github.com/golang/go/blob/master/src/log/log.go?name=release#172)

```
func (l *Logger) Fatal(v ...interface{})
```

Fatal等价于{l.Print(v...); os.Exit(1)}

------

type [Logger](https://github.com/golang/go/blob/master/src/log/log.go?name=release#42) [¶](http://godoc.ml/pkg/log.htm#pkg-index)

```
type Logger struct {
    // contains filtered or unexported fields
}
```

Logger类型表示一个活动状态的记录日志的对象，它会生成一行行的输出写入一个io.Writer接口。每一条日志操作会调用一次io.Writer接口的Write方法。Logger类型的对象可以被多个线程安全的同时使用，它会保证对io.Writer接口的顺序访问。

------

func (*Logger) [Println](https://github.com/golang/go/blob/master/src/log/log.go?name=release#169)

```
func (l *Logger) Println(v ...interface{})
```

Println调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Println相同的方法处理。

------

func (*Logger) [Output](https://github.com/golang/go/blob/master/src/log/log.go?name=release#130)

```
func (l *Logger) Output(calldepth int, s string) error
```

Output写入输出一次日志事件。参数s包含在Logger根据选项生成的前缀之后要打印的文本。如果s末尾没有换行会添加换行符。calldepth用于恢复PC，出于一般性而提供，但目前在所有预定义的路径上它的值都为2。



#### sync包

`import "sync"`

sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。

本包的类型的值不应被拷贝。

------

type [Mutex](https://github.com/golang/go/blob/master/src/sync/mutex.go?name=release#21)

```
type Mutex struct {
    // 包含隐藏或非导出字段
}
```

Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。

func (*Mutex) [Lock](https://github.com/golang/go/blob/master/src/sync/mutex.go?name=release#41)

```
func (m *Mutex) Lock()
```

Lock方法锁住m，如果m已经加锁，则阻塞直到m解锁。

func (*Mutex) [Unlock](https://github.com/golang/go/blob/master/src/sync/mutex.go?name=release#82)

```
func (m *Mutex) Unlock()
```

Unlock方法解锁m，如果m未加锁会导致运行时错误。锁和线程无关，可以由不同的线程加锁和解锁。



#### OS包

`import "os"`

os包提供了操作系统函数的不依赖平台的接口，os包的接口规定为在所有操作系统中都是一致的。非公用的属性可以从操作系统特定的[syscall](http://godoc.org/syscall)包获取。



Variables

```
var (
    ErrInvalid    = errors.New("invalid argument")
    ErrPermission = errors.New("permission denied")
    ErrExist      = errors.New("file already exists")
    ErrNotExist   = errors.New("file does not exist")
)
```

一些可移植的、共有的系统调用错误。

```
var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。

```
var Args []string
```

Args保管了命令行参数，第一个是程序名。



#### strconv包

`import "strconv"`

strconv包实现了基本数据类型和其字符串表示的相互转换

------

func [Atoi](https://github.com/golang/go/blob/master/src/strconv/atoi.go?name=release#195)

```
func Atoi(s string) (i int, err error)
```

Atoi是ParseInt(s, 10, 0)的简写。

func [ParseInt](https://github.com/golang/go/blob/master/src/strconv/atoi.go?name=release#150)

```
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

返回字符串表示的整数值，接受正负号。

base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；

bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

------

#### encoding包

`import "encoding"`

encoding包定义了供其它包使用的可以将数据在字节水平和文本表示之间转换的接口。encoding/gob、encoding/json、encoding/xml三个包都会检查使用这些接口。因此，只要实现了这些接口一次，就可以在多个包里使用。标准包内建类型time.Time和net.IP都实现了这些接口。接口是成对的，分别产生和还原编码后的数据。

------

##### `import"encoding/json"`

json包实现了json对象的编解码，参见[RFC 4627](http://tools.ietf.org/html/rfc4627)。Json对象和go类型的映射关系请参见Marshal和Unmarshal函数的文档。



type [Marshaler](https://github.com/golang/go/blob/master/src/encoding/json/encode.go?name=release#191)

```
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
```

实现了Marshaler接口的类型可以将自身序列化为合法的json描述。

func [Marshal](https://github.com/golang/go/blob/master/src/encoding/json/encode.go?name=release#131)

```
func Marshal(v interface{}) ([]byte, error)
```

Marshal函数返回v的json编码。

func [MarshalIndent](https://github.com/golang/go/blob/master/src/encoding/json/encode.go?name=release#141)

```
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

MarshalIndent类似Marshal但会使用缩进将输出格式化。

func [Unmarshal](https://github.com/golang/go/blob/master/src/encoding/json/decode.go?name=release#67)

```
func Unmarshal(data []byte, v interface{}) error
```

Unmarshal函数解析json编码的数据并将结果存入v指向的值。

Unmarshal和Marshal做相反的操作，必要时映射、切片或指针按照某种规则也行。

------



type [Decoder](https://github.com/golang/go/blob/master/src/encoding/json/stream.go?name=release#14) [¶](http://godoc.ml/pkg/encoding_json.htm#pkg-index)

```
type Decoder struct {
    // 内含隐藏或非导出字段
}
```

Decoder从输入流解码json对象



func [NewDecoder](https://github.com/golang/go/blob/master/src/encoding/json/stream.go?name=release#26)

```
func NewDecoder(r io.Reader) *Decoder
```

NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。

func (*Decoder) [Decode](https://github.com/golang/go/blob/master/src/encoding/json/stream.go?name=release#39)

```
func (dec *Decoder) Decode(v interface{}) error
```

Decode从输入流读取下一个json编码值并保存在v指向的值里，参见Unmarshal函数的文档获取细节信息

------



#### text包

这个包没有，只有子包。

------

##### `import"text/template"`

​	规则较多，需要时再看文档。

template包实现了数据驱动的用于生成文本输出的模板。通过将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出。用作模板的输入文本必须是utf-8编码的文本。"Action"—数据运算和控制单位—由"{{"和"}}"界定；在Action之外的所有文本都不做修改的拷贝到输出中。

func (*Template) [Execute](https://github.com/golang/go/blob/master/src/text/template/exec.go?name=release#129)

```
func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
```

Execute方法将解析好的模板应用到data上，并将输出写入wr。如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行。

func [Must](https://github.com/golang/go/blob/master/src/text/template/helper.go?name=release#21)

```
func Must(t *Template, err error) *Template
```

Must函数用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化：

```
var t = template.Must(template.New("name").Parse("text"))

```

func [New](https://github.com/golang/go/blob/master/src/text/template/template.go?name=release#35)

```
func New(name string) *Template
```

创建一个名为name的模板。

func (*Template) [Funcs](https://github.com/golang/go/blob/master/src/text/template/template.go?name=release#146)

```
func (t *Template) Funcs(funcMap FuncMap) *Template
```

Funcs方法向模板t的函数字典里加入参数funcMap内的键值对。如果funcMap某个键值对的值不是函数类型或者返回值不符合要求会panic。但是，可以对t函数列表的成员进行重写。方法返回t以便进行链式调用。











