/**
 *
 * @param method HTTP method, e.g. GET, POST, PUT, DELETE
 * @param data Object to be sent in the request body
 * @returns {Promise<Response>}
 */
export const doReq = async (method, data)  => await fetch(location.pathname, {
        headers: {
            'Content-Type': 'application/json'
        },
        method,
        body: JSON.stringify(data)
    })