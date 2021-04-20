export default (baseUrl) => {
    return fetch(baseUrl + "/api/v1/getNewBotConfig",{
        method: "GET"
        }).then(resp => {
        return resp.json();
    });
};