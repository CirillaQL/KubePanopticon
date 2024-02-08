import http from "../http.js"

export const getClusterAllPodCPUUsage = () => {
    return http({
        method: "get",
        url: `/cluster/cluster-all-pods-cpu-usage`
    });
};