rdwayback
---------

Take a domain and find unique urls or subdomains from the Wayback Machine.

### Usage 

|Arg|Type|Description|
|---|----|---|
|`-d`|string|Input domain|
|`-sub`|bool|Get subdomains (default false)|
|`-o`|string|Output file path, optional|

#### Examples
Find urls for a domain and write to output file:
```
> go run rdwayback.go -d example.com -o output.txt 
http://example.com:80/how-to-format-hard-drive-windows-xp
http://www.example.com/thecode.html
http://example.com/19667/ceredigion-joins-lights-out-event/
http://www.example.com:80/articles-json/
...
```

Find subdomains for a domain:
```
> go run rdwayback.go -d example.com -sub
www.example.com.:80
1.question.api.example.com
example.com:80
www.example.com:80
...
```