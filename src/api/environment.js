import requests from '@/utils/requests'

export function getEnvironmentTotal() {
    return requests({
        url: "/api/v1/environment-total",
        method: "get"
    })
}

export function getEnvironmentById(id) {
    return requests({
        url: "/api/v1/environment/" + id,
        method: "get"
    })
}

export function getEnvironmentsByPagination(params) {
    return requests({
        url: "/api/v1/environments",
        method: "get",
        params
    })
}

export function addEnvironment(data) {
    return requests({
        url: "/api/v1/environment",
        method: "post",
        data
    })
}

export function editEnvironment(data) {
    return requests({
        url: "/api/v1/environment",
        method: "put",
        data
    })
}

export function deleteEnvironmentById(id) {
    return requests({
        url: "/api/v1/environment/" + id,
        method: "delete"
    })
}