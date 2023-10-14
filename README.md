# React Server Component Parser (v0.0.1 alpha)

If you have used Next.js with AppRouter, you may have noticed that browser get React Server Component from Next server.
You can check this request in Network tab of Chrome DevTools.

Sometimes, it doesn't appear in the tab. Then you can copy curl command and run it in terminal as the following.

```
curl 'http://localhost:3000/page3?_rsc=1pmm9' \
  -H 'Accept: */*' \
  -H 'Accept-Language: en-US,en;q=0.9,ja;q=0.8,pl;q=0.7' \
  -H 'Connection: keep-alive' \
  -H 'Cookie: Idea-bf3098b7=c656c584-9e67-4a35-b4c3-11248ca02880; next-auth.csrf-token=02f1ed3bf17904c2bc7917b337af6af072b2bd26a69cbababf8211bc3d3f97de%7C85fa62240c44c41e20205d1e8465520bb7173c70ff9d953562d6f17606a8b4a0; _ga=GA1.1.307397531.1696336373; next-auth.callback-url=http%3A%2F%2Flocalhost%3A3000%2F; _ga_FJFF8ZDYCT=GS1.1.1696945530.5.0.1696945530.0.0.0; io=rj8KxW0REocIUSRTAAAz' \
  -H 'Next-Router-State-Tree: %5B%22%22%2C%7B%22children%22%3A%5B%22__PAGE__%22%2C%7B%7D%5D%7D%2Cnull%2Cnull%2Ctrue%5D' \
  -H 'Next-Url: /' \
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

rscq is a tool to parse this payload. You can try `cat examples/sample2.rsc | go run main.go`

```
0 [N157NNKhIC8b0k4urA243 [[children [id page3 d] [[id page3 d] map[children:[__PAGE__ map[]]]] $L1 [[] $L2]]]]
3 Import id: name: chunks[ 3 elements ]
4 Import id: name: chunks[ 3 elements ]
1 <$L3 segmentPath=[children [id page3 d] children] childProp=map[current:[$L5 $L6 <nil>] segment:__PAGE__] parallelRouterKey=children template=<$L4></$L4>></$L3>
2 [<meta key=0 charSet=utf-8></meta> <meta key=1 name=viewport content=width=device-width, initial-scale=1></meta> <link key=2 rel=icon href=/favicon.ico type=image/x-icon sizes=16x16></link>]
8 "$Sreact.suspense"
9 Import id: name:Client4 chunks[ 2 elements ]
6 <div>
    [$L7 <$8 fallback=<div>
    Preparing
</div>>
    <$L9>
    $La
</$L9>
</$8>]
</div>
5 <nil>
b Import id: name:Client1 chunks[ 2 elements ]
7 <div>
    [This is Light Server Component <div>
    <$Lb title=Light Component></$Lb>
</div>]
</div>
a <div>
    Heavy Server Component
</div>

```
