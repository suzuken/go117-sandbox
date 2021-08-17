# May be output like below:
# 
# 2021/08/18 08:41:44 http: URL query contains semicolon, which is no longer a supported separator; parts of the query may be stripped when parsed; see golang.org/issue/25192
#
# see also https://golang.org/doc/go1.17#semicolons
test:
	curl -v "localhost:8080/?a=1;b=2&c=3"
