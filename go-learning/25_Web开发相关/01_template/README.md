# template

### 模板引擎：
text template,
html用于生成html页面

1. 板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
2. 模板文件中使用{{和}}包裹和标识需要传入的数据。
3. 传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
4. 除{{和}}包裹的内容外，其他内容均不做修改原样输出
### 解析模板文件

```go
func (t *Template) Parse(src string) (*Template, error)
func ParseFiles(filenames ...string) (*Template, error)
func ParseGlob(pattern string) (*Template, error)
```

### 模板渲染
```go
