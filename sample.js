async function fetchProducto(method, params, id ) {
    id = id || 0;
    const request = JSON.stringify({
        jsonrpc: '2.0',
        method: method,
        params: params,
        id: id,
    });
    console.log(request);
    const response = await fetch('http://localhost:3001/rpc', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: request
    });
    return await response.json();
}

async function fetch1(x) {
    const request = JSON.stringify(x);
    console.log(request);
    const response = await fetch('http://localhost:3001/rpc', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: request
    });
    return await response.json();
}