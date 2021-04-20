export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/fetchKelpErrors",{
        method: "GET"
        }).then(resp => {
        return resp.json();
    });
};