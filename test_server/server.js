var http = require('http');
var _c = 0;
function reset () { _c = 0; console.log('\n----------') }
var to = setTimeout(reset, 1000)
var connections = 0

function responseResult (res, block) {
  res.write('i\'ll write some data\n')
  for (var i = 0; i < _c; i++) {
    res.write('...and some more data\n')
  }
  setTimeout(function () {
    connections--
    res.end('Hello World\n')
  }, block)
}

http.createServer(function (req, res) {
  connections++

  clearTimeout(to)
  to = setTimeout(reset, 1000)

  res.writeHead(200, {'Content-Type': 'text/plain'});
  var b = Math.random()*1000*(1+_c/10)
  responseResult(res, b)
  console.log(++_c)
}).listen(1337, '127.0.0.1');


setInterval(function () {
  if (connections > 0) { 
    console.log(new Date() + ' : Concurrent Connections : ' + connections)
  }
}, 500)
console.log('Server running at http://127.0.0.1:1337/');
