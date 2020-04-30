//5382. HTML 实体解析器
//map匹配并替换
func entityParser(text string) string {
    var m = make(map[string]byte)
    m["&quot;"] = '"'
    m["&apos;"] = '\''
    m["&amp;"] = '&'
    m["&gt;"] = '>'
    m["&lt;"] = '<'
    m["&frasl;"] = '/'
    for k, v := range m {
        text = strings.Replace(text, k, string(v), -1)
    }
    return text
}