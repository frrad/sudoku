def cells(a,b,x,y,template):
    ans = []
    for i in xrange(a,b):
        acc = []
        for j in xrange(x,y):
            acc.append(template % (i,j))
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

def sudoku(template):
    tbl = []
    for x in xrange(3):
        acc = []
        for y in xrange(3):
            acc.append(table(cells(3*x,3*(x+1),3*y,3*(y+1),template)))
        tbl.append(acc)
    return table(tbl, True)

prepasta = '''
<html>
    <head><title>Sudoku Solver</title></head>
    <body>
		<meta charset="utf-8">
		<script src="wasm_exec.js"></script>
		<script>
                const go = new Go();
                let mod, inst;
                WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(async (result) => {
                        mod = result.module;
                        inst = result.instance;
                        await go.run(inst)
                });
		</script>
'''

postpasta = '''
    </body>
</html>
'''


print prepasta
print sudoku('<input type="text" size="1" id="input-%d-%d">')
print '<button onClick="solveSudoku();">Solve</button>'
print '<br><br>'
print '<p id="messageP"></p>'
print sudoku('<p id="ans-%d-%d">&nbsp;&nbsp;&nbsp;</p>')
print postpasta
