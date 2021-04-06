export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/listBots",{
        method: "GET"
        }).then(resp => {
        return resp.json();
    });
};