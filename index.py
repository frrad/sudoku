def cells(a,b,x,y):
    ans = []
    for i in xrange(a,b):
        acc = []
        for j in xrange(x,y):
            acc.append('<input type="text" size="1" id="ans-%d-%d">' % (i,j))
        ans.append(acc)
    return ans

def table(rows, border=False):
    b = ""
    if border:
        b = ' border="1"'

    ans =""
    ans += "<table%s>\n" %b
    for row in rows:
        ans += "  <tr>\n"
        for col in row:
            ans+=  "    <td>%s</td>\n" % col
        ans += "  </tr>\n"
    ans+= "</table>\n"
    return ans 



prepasta = '''
<html>
    <head>
    </head>
    <body>
		<meta charset="utf-8">
		<script src="wasm_exec.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
				go.run(result.instance);
			});
		</script>
'''

postpasta = '''
    </body>
</html>
'''

tbl = []
for x in xrange(3):
    acc = []
    for y in xrange(3):
        acc.append(table(cells(3*x,3*(x+1),3*y,3*(y+1))))
    tbl.append(acc)

print prepasta
print table(tbl, True)
print postpasta
