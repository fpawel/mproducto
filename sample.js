async function fetchProducto(method, params, id ) {
    id = id || 0;
    const response = await fetch('http://localhost:3000/rpc', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            jsonrpc: '2.0',
            method: method,
            params: params,
            id: id,
        })
    });
    return await response.json();
}