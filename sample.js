var response = await fetch('http://localhost:3000/rpc', {
    method: 'POST',
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    },
    body: JSON.stringify({
        jsonrpc: '2.0',
        method: 'Auth.Login',
        params: {User:"user_name",Pass:"user_pass"},
        id: '0',
    })
});
var data = await response.json();
console.log(data);