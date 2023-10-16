# React Server Component Parser (v0.2.0 beta)

### Install
```
go install github.com/horita-yuya/rscq@0.2.0-beta
```

If you have used Next.js with AppRouter, you may have noticed that browser get React Server Component from Next server.
You can check this request in Network tab of Chrome DevTools.

Sometimes, it doesn't appear in the tab. Then you can copy curl command and run it in terminal as the following.

```
curl 'http://localhost:3000/page3?_rsc=8nwf1' \
  -H 'Accept: */*' \
  -H 'Accept-Language: en-US,en;q=0.9,ja;q=0.8,pl;q=0.7' \
  -H 'Connection: keep-alive' \
  -H 'Next-Router-State-Tree: %5B%22%22%2C%7B%22children%22%3A%5B%5B%22id%22%2C%22page3%22%2C%22d%22%5D%2C%7B%22children%22%3A%5B%22__PAGE__%22%2C%7B%7D%5D%7D%2Cnull%2C%22refetch%22%5D%7D%5D' \
  -H 'Next-Url: /page3' \
  -H 'RSC: 1' \
  -H 'Referer: http://localhost:3000/' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36' \
  -H 'sec-ch-ua: "Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --compressed
```

This returns a RSC payload.

```
0:["N157NNKhIC8b0k4urA243",[["children",["id","page3","d"],[["id","page3","d"],{"children":["__PAGE__",{}]}],"$L1",[[],"$L2"]]]]
3:I{"id":1443,"chunks":["272:static/chunks/webpack-bf6fccfdfe35d157.js","971:static/chunks/fd9d1056-2581dce591ac5cde.js","864:static/chunks/864-b23738e09c185fe9.js"],"name":"","async":false}
4:I{"id":8639,"chunks":["272:static/chunks/webpack-bf6fccfdfe35d157.js","971:static/chunks/fd9d1056-2581dce591ac5cde.js","864:static/chunks/864-b23738e09c185fe9.js"],"name":"","async":false}
1:["$","$L3",null,{"parallelRouterKey":"children","segmentPath":["children",["id","page3","d"],"children"],"loading":"$undefined","loadingStyles":"$undefined","hasLoading":false,"error":"$undefined","errorStyles":"$undefined","template":["$","$L4",null,{}],"templateStyles":"$undefined","notFound":"$undefined","notFoundStyles":"$undefined","childProp":{"current":["$L5","$L6",null],"segment":"__PAGE__"},"styles":[]}]
2:[["$","meta","0",{"charSet":"utf-8"}],["$","meta","1",{"name":"viewport","content":"width=device-width, initial-scale=1"}],["$","link","2",{"rel":"icon","href":"/favicon.ico","type":"image/x-icon","sizes":"16x16"}]]
8:"$Sreact.suspense"
9:I{"id":7772,"chunks":["233:static/chunks/233-f9c6a8ef9ae4b8c3.js","531:static/chunks/app/[id]/page-d4e2ada29cefa5fc.js"],"name":"Client4","async":false}
6:["$","div",null,{"children":["$L7",["$","$8",null,{"fallback":["$","div",null,{"children":"Preparing"}],"children":["$","$L9",null,{"children":"$La"}]}]]}]
5:null
b:I{"id":2233,"chunks":["233:static/chunks/233-f9c6a8ef9ae4b8c3.js","531:static/chunks/app/[id]/page-d4e2ada29cefa5fc.js"],"name":"Client1","async":false}
7:["$","div",null,{"children":["This is Light Server Component",["$","div",null,{"children":["$","$Lb",null,{"title":"Light Component"}]}]]}]
a:["$","div",null,{"children":"Heavy Server Component"}]
```

I can read it if I try, but I want it to be easier to read.

rscq is a tool to parse this payload. You can try `cat examples/sample2.rsc | rscq`

```
0--------------------
  N157NNKhIC8b0k4urA243
  children
  id
  page3
  d
  id
  page3
  d
  map[children:[__PAGE__ map[]]]
  $L1
  $L2
--------------------0
3--------------------
  272:static/chunks/webpack-bf6fccfdfe35d157.js
  971:static/chunks/fd9d1056-2581dce591ac5cde.js
  864:static/chunks/864-b23738e09c185fe9.js
--------------------3
4--------------------
  272:static/chunks/webpack-bf6fccfdfe35d157.js
  971:static/chunks/fd9d1056-2581dce591ac5cde.js
  864:static/chunks/864-b23738e09c185fe9.js
--------------------4
1--------------------
  <$L3
    childProp={map[current:[$L5 $L6 <nil>] segment:__PAGE__]}
    segmentPath={[children [id page3 d] children]}
    template={
      <$L4>
      </$L4>
    }

    parallelRouterKey={children}
  >
  </$L3>
--------------------1
2--------------------
  <meta key=0
    charSet={utf-8}
  >
  </meta>
  <meta key=1
    name={viewport}
    content={width=device-width, initial-scale=1}
  >
  </meta>
  <link key=2
    rel={icon}
    href={/favicon.ico}
    type={image/x-icon}
    sizes={16x16}
  >
  </link>
--------------------2
8--------------------
  "$Sreact.suspense"
--------------------8
9--------------------
  233:static/chunks/233-f9c6a8ef9ae4b8c3.js
  531:static/chunks/app/[id]/page-d4e2ada29cefa5fc.js
--------------------9
6--------------------
  <div>
    $L7

    <$8
      fallback={
        <div>
          Preparing
        </div>
      }

    >
      <$L9>
        $La
      </$L9>
    </$8>
  </div>
--------------------6
5--------------------
  <nil>
--------------------5
b--------------------
  233:static/chunks/233-f9c6a8ef9ae4b8c3.js
  531:static/chunks/app/[id]/page-d4e2ada29cefa5fc.js
--------------------b
7--------------------
  <div>
    This is Light Server Component

    <div>
      <$Lb
        title={Light Component}
      >
      </$Lb>
    </div>
  </div>
--------------------7
a--------------------
  <div>
    Heavy Server Component
  </div>
--------------------a
```
