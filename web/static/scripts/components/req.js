/**
 *
 * @param method HTTP method, e.g. GET, POST, PUT, DELETE
 * @param data Object to be sent in the request body
 * @returns {Promise<Response>}
 */
export async function doReq(method, data)  {
    const headers = data ? {
        'Content-Type': 'application/json'
    }:null,
        body = data ? JSON.stringify(data):null

    await fetch(location.pathname, {
        headers,
        method,
        body
    })

}