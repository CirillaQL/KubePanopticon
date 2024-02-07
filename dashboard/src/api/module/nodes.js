import http from "../http.js"

export const getNodeList = () => {
    return http({
        method: "get",
        url: `/nodes/list`
    });
};