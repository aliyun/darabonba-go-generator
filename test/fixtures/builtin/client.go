// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  dara "github.com/alibabacloud-go/tea/tea"
  daraarray "github.com/alibabacloud-go/dara/array"
  "strings"
  "encoding/hex"
  "encoding/base64"
  darabytes "github.com/alibabacloud-go/dara/bytes"
  darafile "github.com/alibabacloud-go/dara/file"
  darafrom "github.com/alibabacloud-go/dara/form"
  darajson "github.com/alibabacloud-go/dara/json"
  "fmt"
  daramath "github.com/alibabacloud-go/dara/math"
  darastream "github.com/alibabacloud-go/dara/stream"
  "regexp"
  "strconv"
  daraurl "github.com/alibabacloud-go/dara/url"
  daraxml "github.com/alibabacloud-go/dara/xml"
)

type M struct {
  A *string `json:"a,omitempty" xml:"a,omitempty" require:"true"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) SetA(v string) *M {
  s.A = &v
  return s
}


func ArrayTest (args []*string) (_err error) {
  if (len(args) > 0) && daraarray.Contains(args, "cn-hanghzou") {
    index := daraarray.Index(args, "cn-hanghzou")
    regionId := args[index]
    all := daraarray.Join(args, ",")
    first := daraarray.Shift(&args)
    last := daraarray.Pop(&args)
    length1 := daraarray.Unshift(&args, first)
    length2 := daraarray.Push(&args, last)
    length3 := length1 + length2
    longStr := "long" + first + last
    fullStr := daraarray.Join(args, ",")
    newArr := []*string{dara.String("asc"), dara.String("test1"), dara.String("test2")}
    cArr := dara.ToStringSlice(daraarray.Concat(newArr, args))
    acsArr := dara.ToStringSlice(daraarray.Sort(newArr, "asc"))
    acsArr1 := dara.ToStringSlice(daraarray.Sort(newArr, dara.StringValue(newArr[0])))
    descArr := dara.ToStringSlice(daraarray.Sort(newArr, "desc"))
    tmpStr := dara.StringValue(newArr[1])
    llArr := dara.ToStringSlice(daraarray.Concat(acsArr, descArr))
    daraarray.Append(&llArr, "test")
    daraarray.Remove(&llArr, "test")
    if CheckStr(newArr[3]) {
      // TODO
    } else if CheckStr(dara.String(all)) {
      // TODO
    }

  }

  return _err
}

func CheckStr (str *string) (_result *bool) {
  _result = dara.Bool(strings.Length(str) > 2)
  return _result
}

func BytesTest (args []*string) (_err error) {
  fullStr := daraarray.Join(args, ",")
  data := dara.ToBytes(fullStr, "utf8")
  newFullStr := string(data)
  if fullStr != newFullStr {
    return
  }

  hexStr := hex.EncodeToString(data)
  base64Str := base64.StdEncoding.EncodeToString(data)
  length := len(data)
  obj := string(data)
  data2 := darabytes.From(fullStr, "base64")
  return _err
}

func DateTest (args []*string) (_err error) {
  date, _err := *dara.Date.NewDate("2023-09-12 17:47:31.916000 +0800 UTC")
  if _err != nil {
    return _err
  }

  dateStr := date.Format("YYYY-MM-DD HH:mm:ss")
  if CheckStr(dara.String(dateStr)) {
    m1 := &M{
      A: date.Format("YYYY-MM-DD HH:mm:ss"),
    }
  }

  timestamp := date.Unix()
  yesterday := date.Sub("day", 1)
  oneDay := date.Diff("day", yesterday)
  tomorrow := date.Add("day", 1)
  twoDay := tomorrow.Diff("day", date) + oneDay
  hour := date.Hour()
  minute := date.Minute()
  second := date.Second()
  dayOfMonth := date.DayOfMonth()
  dayOfWeek := date.DayOfWeek()
  weekOfYear := date.WeekOfYear()
  month := date.Month()
  year := date.Year()
  return _err
}

func EnvTest (args []*string) (_err error) {
  es := os.Getenv("TEST")
  ma, _err := os.Setenv("TEST", es + "test")
  if _err != nil {
    return _err
  }

  ma1, _err := os.Setenv("TEST1", "test1")
  if _err != nil {
    return _err
  }

  ma2, _err := os.Setenv("TEST2", es)
  if _err != nil {
    return _err
  }

  return _err
}

func FileTest (args []*string) (_err error) {
  if darafile.Exists("/tmp/test") {
    file, _err := *dara.File.NewFile("/tmp/test")
    if _err != nil {
      return _err
    }

    path := file.Path()
    length := file.Length() + 10
    createTime, _err := file.CreateTime()
    if _err != nil {
      return _err
    }

    modifyTime, _err := file.ModifyTime()
    if _err != nil {
      return _err
    }

    timeLong := modifyTime.Diff("minute", createTime)
    data, _err := file.Read(300)
    if _err != nil {
      return _err
    }

    _err = file.Write(darabytes.From("test", "utf8"))
    if _err != nil {
      return _err
    }
    rs, _err := darafile.CreateReadStream("/tmp/test")
    if _err != nil {
      return _err
    }

    ws, _err := darafile.CreateWriteStream("/tmp/test")
    if _err != nil {
      return _err
    }

  }

  return _err
}

func FormTest (args []*string) (_err error) {
  m := map[string]interface{}{
    "key1": "test1",
    "key2": "test2",
    "key3": 3,
    "key4": map[string]interface{}{
      "key5": 123,
      "key6": "321",
    },
  }
  form := darafrom.ToFormString(m)
  form = form + "&key7=23233&key8=" + darafrom.GetBoundary()
  r := darafrom.ToFileForm(m, darafrom.GetBoundary())
  return _err
}

func JsonTest (args []*string) (_err error) {
  m := map[string]interface{}{
    "key1": "test1",
    "key2": "test2",
    "key3": 3,
    "key4": map[string]interface{}{
      "key5": 123,
      "key6": "321",
    },
  }
  m1 := &M{
    A: dara.String("test"),
  }
  ms := darajson.Stringify(m)
  m1s := darajson.Stringify(m1)
  ma := darajson.ParseJSON(ms)
  arrStr := "[1,2,3,4]"
  arr := darajson.ParseJSON(arrStr)
  return _err
}

func LogerTest (args []*string) (_err error) {
  fmt.Printf("[LOG] %s\n", "test")
  fmt.Printf("[INFO] %s\n", "test")
  fmt.Printf("[WARNING] %s\n", "test")
  fmt.Printf("[DEBUG] %s\n", "test")
  fmt.Printf("[ERROR] %s\n", "test")
  return _err
}

func MapTestCase (args []*string) (_err error) {
  mapTest := map[string]*string{
    "key1": dara.String("value1"),
    "key2": dara.String("value2"),
    "key3": dara.String("value3"),
  }
  length := DaraMap.Size(mapTest)
  num := length + 3
  keys := DaraMap.KeySet(mapTest)
  allKey := ""
  for _, key := range keys {
    allKey = allKey + key
  }
  entries := DaraMap.Entries(mapTest)
  newKey := ""
  newValue := ""
  for _, e := range entries {
    newKey = newKey + e.Key
    newValue = newValue + e.Value.(string)
  }
  json := darajson.Stringify(mapTest)
  mapTest2 := map[string]*string{
    "key1": dara.String("value4"),
    "key4": dara.String("value5"),
  }
  mapTest3 := dara.Merge(mapTest , mapTest2)
  if dara.ToString(mapTest3["key1"]) == "value4" {
    return
  }

  return _err
}

func NumberTest (args []*string) (_err error) {
  num := 3.2
  inum := int(num)
  lnum := int64(num)
  fnum := float32(num)
  dnum := float64(num)
  inum = int(inum)
  lnum = int64(inum)
  fnum = float32(inum)
  dnum = float64(inum)
  inum = int(lnum)
  lnum = int64(lnum)
  fnum = float32(lnum)
  dnum = float64(lnum)
  inum = int(fnum)
  lnum = int64(fnum)
  fnum = float32(fnum)
  dnum = float64(fnum)
  inum = int(dnum)
  lnum = int64(dnum)
  fnum = float32(dnum)
  dnum = float64(dnum)
  lnum = int64(inum)
  inum = int(lnum)
  randomNum := daramath.Random()
  inum = daramath.Floor(inum)
  inum = daramath.Round(inum)
  min := daramath.Min(inum, fnum)
  max := daramath.Max(inum, fnum)
  return _err
}

func StreamTest (args []*string) (_err error) {
  if darafile.Exists("/tmp/test") {
    rs, _err := darafile.CreateReadStream("/tmp/test")
    if _err != nil {
      return _err
    }

    ws, _err := darafile.CreateWriteStream("/tmp/test")
    if _err != nil {
      return _err
    }

    data, _err := rs.Read(30)
    if _err != nil {
      return _err
    }

    _err = ws.Write(data)
    if _err != nil {
      return _err
    }
    _err = rs.Pipe(ws)
    if _err != nil {
      return _err
    }
    data, _err = darastream.ReadAsBytes(rs)
    if _err != nil {
      return _err
    }

    obj, _err := darastream.ReadAsJSON(rs)
    if _err != nil {
      return _err
    }

    jsonStr, _err := darastream.ReadAsString(rs)
    if _err != nil {
      return _err
    }

  }

  return _err
}

func StringTest (args []*string) (_err error) {
  fullStr := daraarray.Join(args, ",")
  args = fullStr.Split(",")
  if (strings.Length(fullStr) > 0) && strings.Contains(fullStr, "hangzhou") {
    newStr1, _err := regexp.MustCompile(`hangzhou`).ReplaceAllString(fullStr, "beijing")
    if _err != nil {
      return _err
    }

  }

  if strings.HasPrefix(fullStr, "cn") {
    newStr2, _err := regexp.MustCompile(`(?i)cn`).ReplaceAllString(fullStr, "zh")
    if _err != nil {
      return _err
    }

  }

  if strings.HasSuffix(fullStr, "beijing") {
    newStr3, _err := regexp.MustCompile(`beijing`).ReplaceAllString(fullStr, "chengdu")
    if _err != nil {
      return _err
    }

  }

  start := strings.Index(fullStr, "beijing")
  end := start + 7
  region := fullStr[start: end]
  lowerRegion := strings.ToLower(region)
  upperRegion := strings.ToUpper(region)
  if region == "beijing" {
    region = region + " "
    region = strings.TrimSpace(region)
  }

  tb := dara.ToBytes(fullStr, "utf8")
  em := "xxx"
  if len(em) == 0 {
    return
  }

  num := "32.0a"
  inum := strconv.Atoi(num) + 3
  lnum := strconv.ParseInt(num, 10, 64)
  fnum := strconv.ParseFloat(num, 32) + 1
  dnum := strconv.ParseFloat(num, 64) + 1
  return _err
}

func UrlTest (args []*string) (_err error) {
  url, _err := *$URL.NewURL(dara.StringValue(args[0]))
  if _err != nil {
    return _err
  }

  path := url.Path()
  pathname := url.Pathname()
  protocol := url.Protocol()
  hostname := url.Hostname()
  port := url.Port()
  host := url.Host()
  hash := url.Hash()
  search := url.Search()
  href := url.Href()
  auth := url.Auth()
  url2, _err := daraurl.Parse(dara.StringValue(args[1]))
  if _err != nil {
    return _err
  }

  path = url2.Path()
  newUrl := daraurl.UrlEncode(dara.StringValue(args[2]))
  newSearch := daraurl.PercentEncode(search)
  newPath := daraurl.PathEncode(pathname)
  all := "test" + path + protocol + hostname + hash + search + href + auth + newUrl + newSearch + newPath
  return _err
}

func XmlTest (args []*string) (_err error) {
  m := map[string]interface{}{
    "key1": "test1",
    "key2": "test2",
    "key3": 3,
    "key4": map[string]interface{}{
      "key5": 123,
      "key6": "321",
    },
  }
  xml := daraxml.ToXML(m)
  xml = xml + "<key7>132</key7>"
  respMap := daraxml.ParseXml(xml, nil)
  return _err
}

func ReturnAny () (_result interface{}) {
  panic("No Support!")
}

func Main (args []*string) (_err error) {
  _err = ArrayTest(args)
  if _err != nil {
    return _err
  }
  _err = BytesTest(args)
  if _err != nil {
    return _err
  }
  _err = DateTest(args)
  if _err != nil {
    return _err
  }
  _err = EnvTest(args)
  if _err != nil {
    return _err
  }
  _err = FileTest(args)
  if _err != nil {
    return _err
  }
  _err = FormTest(args)
  if _err != nil {
    return _err
  }
  _err = LogerTest(args)
  if _err != nil {
    return _err
  }
  _err = MapTestCase(args)
  if _err != nil {
    return _err
  }
  _err = NumberTest(args)
  if _err != nil {
    return _err
  }
  _err = StreamTest(args)
  if _err != nil {
    return _err
  }
  _err = StringTest(args)
  if _err != nil {
    return _err
  }
  _err = UrlTest(args)
  if _err != nil {
    return _err
  }
  _err = XmlTest(args)
  if _err != nil {
    return _err
  }
  a := dara.ToInt(dara.StringValue(args[0])) + 10
  b := dara.ToString(a) + dara.StringValue(args[1]) + dara.ToString(ReturnAny())
  c := dara.ToInt(b) + dara.ToInt(a) + dara.ToInt(ReturnAny())
  d := dara.ToInt8(b) + dara.ToInt8(a) + dara.ToInt8(ReturnAny())
  e := dara.ToInt16(b) + dara.ToInt16(a) + dara.ToInt16(ReturnAny())
  f := dara.ToInt32(b) + dara.ToInt32(a) + dara.ToInt32(ReturnAny())
  g := dara.ToInt64(b) + dara.ToInt64(a) + dara.ToInt64(ReturnAny())
  h := dara.ToInt64(b) + dara.ToInt64(a) + dara.ToInt64(ReturnAny())
  i := dara.ToUint64(b) + dara.ToUint64(a) + dara.ToUint64(ReturnAny())
  j := dara.ToUint8(b) + dara.ToUint8(a) + dara.ToUint8(ReturnAny())
  k := dara.ToUint16(b) + dara.ToUint16(a) + dara.ToUint16(ReturnAny())
  l := dara.ToUint32(b) + dara.ToUint32(a) + dara.ToUint32(ReturnAny())
  m := dara.ToUint64(b) + dara.ToUint64(a) + dara.ToUint64(ReturnAny())
  n := dara.ToFloat32(b) + dara.ToFloat32(a) + dara.ToFloat32(ReturnAny())
  o := dara.ToFloat64(b) + dara.ToFloat64(a) + dara.ToFloat64(ReturnAny())
  if dara.ToBoolean(dara.StringValue(args[2])) {
    data := []byte(ReturnAny())
    length := len(data)
    test := data
    maps := map[string]*string{
      "key": dara.String("value"),
    }
    obj := maps
    ws := dara.ToWritable(obj)
    rs := dara.ToReadable(maps)
    data, _err = rs.Read(30)
    if _err != nil {
      return _err
    }

    if !darautils.isNull(data) {
      _err = ws.Write(data)
      if _err != nil {
        return _err
      }
    }

  }

  darautils.sleep(a)
  defaultVal := dara.ToString(darautils.default(dara.StringValue(args[0]), dara.StringValue(args[1])))
  if defaultVal == b {
    return
  }

  return _err
}

