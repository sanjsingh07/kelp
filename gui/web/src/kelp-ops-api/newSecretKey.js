export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/newSecretKey",{
        method: "GET"
        }).then(resp => {
        return resp.text();
    });
};