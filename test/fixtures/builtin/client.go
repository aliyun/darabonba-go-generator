// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  "github.com/alibabacloud-go/tea/dara"
  "encoding/hex"
  "encoding/base64"
  "fmt"
  "regexp"
  "strconv"
  "time"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetA(v string) *M
  GetA() *string 
  SetB(v int) *M
  GetB() *int 
}

type M struct {
  A *string `json:"a,omitempty" xml:"a,omitempty" require:"true"`
  B *int `json:"b,omitempty" xml:"b,omitempty" require:"true"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetA() *string  {
  return s.A
}

func (s *M) GetB() *int  {
  return s.B
}

func (s *M) SetA(v string) *M {
  s.A = &v
  return s
}

func (s *M) SetB(v int) *M {
  s.B = &v
  return s
}

func (s *M) Validate() error {
  return dara.Validate(s)
}


func ArrayTest (args []*string) (_err error) {
  if (len(args) > 0) && dara.ArrContains(args, "cn-hanghzou") {
    index := dara.ArrIndex(args, "cn-hanghzou")
    regionId := args[index]
    all := dara.ArrJoin(args, ",")
    first := dara.ArrShift(&args)
    last := dara.ArrPop(&args)
    length1 := dara.ArrUnshift(&args, first)
    length2 := dara.ArrPush(&args, last)
    length3 := length1 + length2
    longStr := "long" + first + last
    fullStr := dara.ArrJoin(args, ",")
    newArr := []*string{dara.String("asc"), dara.String("test1"), dara.String("test2")}
    cArr := dara.ConcatArr(newArr, args).([]*string)
    nArr := []*int{dara.Int(1), dara.Int(3), dara.Int(4)}
    nnArr := dara.ConcatArr(nArr, []*int{dara.Int(4), dara.Int(5), dara.Int(6)}).([]*int)
    m1 := &M{
      A: dara.String("test"),
    }
    m2 := &M{
      A: dara.String("test2"),
    }
    m3 := &M{
      A: dara.String("test3"),
    }
    sArr := []*M{m1, m2}
    nsArr := dara.ConcatArr(sArr, []*M{m3}).([]*M)
    acsArr := dara.SortArr(newArr, "asc").([]*string)
    acsArr1 := dara.SortArr(newArr, dara.StringValue(newArr[0])).([]*string)
    descArr := dara.SortArr(newArr, "desc").([]*string)
    tmpStr := dara.StringValue(newArr[1])
    llArr := dara.ConcatArr(acsArr, descArr).([]*string)
    dara.ArrAppend(&llArr, "test")
    dara.ArrRemove(&llArr, "test")
    if dara.BoolValue(CheckStr(newArr[3])) {
      // TODO
    } else if dara.BoolValue(CheckStr(dara.String(all))) {
      // TODO
    }

  }

  return _err
}

func CheckStr (str *string) (_result *bool) {
  _result = dara.Bool(dara.Length(str) > 2)
  return _result
}

func BytesTest (args []*string) (_err error) {
  fullStr := dara.ArrJoin(args, ",")
  data := dara.ToBytes(fullStr, "utf8")
  newFullStr := dara.ToString(data)
  if fullStr != newFullStr {
    return
  }

  hexStr := hex.EncodeToString(data)
  base64Str := base64.StdEncoding.EncodeToString(data)
  length := len(data)
  obj := dara.ToString(data)
  data2 := dara.BytesFromString(fullStr, "base64")
  return _err
}

func DateTest (args []*string) (_err error) {
  date, _err := *dara.Date.NewDate("2023-09-12 17:47:31.916000 +0800 UTC")
  if _err != nil {
    return _err
  }

  dateStr := date.Format("YYYY-MM-DD HH:mm:ss")
  if dara.BoolValue(CheckStr(dara.String(dateStr))) {
    m1 := &M{
      A: dara.String(date.Format("YYYY-MM-DD HH:mm:ss")),
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
  if dara.Exists("/tmp/test") {
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

    _err = file.Write(dara.BytesFromString("test", "utf8"))
    if _err != nil {
      return _err
    }
    rs, _err := dara.CreateReadStream("/tmp/test")
    if _err != nil {
      return _err
    }

    ws, _err := dara.CreateWriteStream("/tmp/test")
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
  form := dara.ToFormString(m)
  form = form + "&key7=23233&key8=" + dara.GetBoundary()
  r := dara.ToFileForm(m, dara.GetBoundary())
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
  ms := dara.Stringify(m)
  m1s := dara.Stringify(m1)
  ma := dara.ParseJSON(ms)
  arrStr := "[1,2,3,4]"
  arr := dara.ParseJSON(arrStr)
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
  length := len(mapTest)
  num := length + 3
  keys := dara.KeySet(mapTest)
  allKey := ""
  for _, key := range keys {
    allKey = allKey + key
  }
  entries := dara.Entries(mapTest)
  newKey := ""
  newValue := ""
  for _, e := range entries {
    newKey = newKey + e.Key
    newValue = newValue + dara.StringValue(e.Value.(*string))
  }
  json := dara.Stringify(mapTest)
  mapTest2 := map[string]*string{
    "key1": dara.String("value4"),
    "key4": dara.String("value5"),
  }
  mapTest3 := dara.ToMap(mapTest , mapTest2)
  if dara.ToString(mapTest3["key1"]) == "value4" {
    return
  }

  mapTest4 := map[string]interface{}{
    "key1": "value4",
    "key2": 2,
    "key3": true,
  }
  entries2 := dara.Entries(mapTest4)
  for _, e := range entries2 {
    newKey = newKey + e.Key
    newValue = newValue + dara.ToString(e.Value)
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
  randomNum := dara.Random()
  inum = dara.Floor(inum)
  inum = dara.Round(inum)
  return _err
}

func StreamTest (args []*string) (_err error) {
  if dara.Exists("/tmp/test") {
    rs, _err := dara.CreateReadStream("/tmp/test")
    if _err != nil {
      return _err
    }

    ws, _err := dara.CreateWriteStream("/tmp/test")
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
    data, _err = dara.ReadAsBytes(rs)
    if _err != nil {
      return _err
    }

    obj, _err := dara.ReadAsJSON(rs)
    if _err != nil {
      return _err
    }

    jsonStr, _err := dara.ReadAsString(rs)
    if _err != nil {
      return _err
    }

  }

  return _err
}

func StringTest (args []*string) (_err error) {
  fullStr := dara.ArrJoin(args, ",")
  args = fullStr.Split(",")
  if (dara.Length(fullStr) > 0) && dara.Contains(fullStr, "hangzhou") {
    newStr1, _err := regexp.MustCompile(`hangzhou`).ReplaceAllString(fullStr, "beijing")
    if _err != nil {
      return _err
    }

  }

  if dara.HasPrefix(fullStr, "cn") {
    newStr2, _err := regexp.MustCompile(`(?i)cn`).ReplaceAllString(fullStr, "zh")
    if _err != nil {
      return _err
    }

  }

  if dara.HasSuffix(fullStr, "beijing") {
    newStr3, _err := regexp.MustCompile(`beijing`).ReplaceAllString(fullStr, "chengdu")
    if _err != nil {
      return _err
    }

  }

  start := dara.Index(fullStr, "beijing")
  end := start + 7
  region := fullStr[start: end]
  lowerRegion := dara.ToLower(region)
  upperRegion := dara.ToUpper(region)
  if region == "beijing" {
    region = region + " "
    region = dara.TrimSpace(region)
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
  url, _err := *dara.URL.NewURL(dara.StringValue(args[0]))
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
  url2, _err := dara.Parse(dara.StringValue(args[1]))
  if _err != nil {
    return _err
  }

  path = url2.Path()
  newUrl := dara.UrlEncode(dara.StringValue(args[2]))
  newSearch := dara.PercentEncode(search)
  newPath := dara.PathEncode(pathname)
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
  xml := dara.ToXML(m)
  xml = xml + "<key7>132</key7>"
  respMap := dara.ParseXml(xml, nil)
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
  a := dara.ForceInt(dara.StringValue(args[0])) + 10
  md := &M{
    A: dara.String("test"),
    B: dara.Int(10),
  }
  b := dara.ToString(a) + dara.StringValue(args[1]) + dara.ToString(dara.IntValue(md.B)) + dara.ToString(ReturnAny())
  c := dara.ForceInt(b) + dara.ForceInt(a) + dara.ForceInt(ReturnAny())
  d := dara.ForceInt8(b) + dara.ForceInt8(a) + dara.ForceInt8(ReturnAny())
  e := dara.ForceInt16(b) + dara.ForceInt16(a) + dara.ForceInt16(ReturnAny())
  f := dara.ForceInt32(b) + dara.ForceInt32(a) + dara.ForceInt32(ReturnAny())
  g := dara.ForceInt64(b) + dara.ForceInt64(a) + dara.ForceInt64(ReturnAny())
  h := dara.ForceInt64(b) + dara.ForceInt64(a) + dara.ForceInt64(ReturnAny())
  i := dara.ForceUint64(b) + dara.ForceUint64(a) + dara.ForceUint64(ReturnAny())
  j := dara.ForceUint8(b) + dara.ForceUint8(a) + dara.ForceUint8(ReturnAny())
  k := dara.ForceUint16(b) + dara.ForceUint16(a) + dara.ForceUint16(ReturnAny())
  l := dara.ForceUint32(b) + dara.ForceUint32(a) + dara.ForceUint32(ReturnAny())
  m := dara.ForceUint64(b) + dara.ForceUint64(a) + dara.ForceUint64(ReturnAny())
  n := dara.ForceFloat32(b) + dara.ForceFloat32(a) + dara.ForceFloat32(ReturnAny())
  o := dara.ForceFloat64(b) + dara.ForceFloat64(a) + dara.ForceFloat64(ReturnAny())
  if dara.ForceBoolean(dara.StringValue(args[2])) {
    data := []byte(dara.ToString(ReturnAny()))
    length := len(data)
    test := data
    maps := map[string]*string{
      "key": dara.String("value"),
    }
    obj := dara.ToMap(maps)
    ws := dara.ToWritable(obj)
    rs := dara.ToReadable(maps)
    data, _err = rs.Read(30)
    if _err != nil {
      return _err
    }

    if !dara.IsNil(data) {
      _err = ws.Write(data)
      if _err != nil {
        return _err
      }
    }

  }

  time.Sleep(time.Duration(a) * time.Millisecond)
  defaultVal := dara.ToString(dara.Default(dara.StringValue(args[0]), dara.StringValue(args[1])))
  if defaultVal == b {
    return
  }

  return _err
}

