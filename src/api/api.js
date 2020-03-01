import requests from '@/utils/requests'

export function getApiTotal() {
    return requests({
        url: "/api/v1/api-total",
        method: "get"
    })
}

export function getApiById(id) {
    return requests({
        url: "/api/v1/api/" + id,
        method: "get"
    })
}

export function getApisByPagination(params) {
    return requests({
        url: "/api/v1/apis",
        method: "get",
        params
    })
}

export function addApi(data) {
    return requests({
        url: "/api/v1/api",
        method: "post",
        data
    })
}

export function editApi(data) {
    return requests({
        url: "/api/v1/api",
        method: "put",
        data
    })
}

export function deleteApiById(id) {
    return requests({
        url: "/api/v1/api/" + id,
        method: "delete"
    })
}