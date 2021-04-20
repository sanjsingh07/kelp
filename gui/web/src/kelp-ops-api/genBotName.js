export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/genBotName",{
        method: "GET"
        }).then(resp => {
        return resp.text();
    });
};