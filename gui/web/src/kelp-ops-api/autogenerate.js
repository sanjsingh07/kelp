export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/autogenerate", {
        method: "GET"
        }).then(resp => {
        return resp.json();
    });
};