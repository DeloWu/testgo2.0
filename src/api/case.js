import requests from '@/utils/requests'

export function getCaseTotal() {
    return requests({
        url: "/api/v1/case-total",
        method: "get"
    })
}

export function getCaseById(id) {
    return requests({
        url: "/api/v1/case/" + id,
        method: "get"
    })
}

export function getCasesByPagination(params) {
    return requests({
        url: "/api/v1/cases",
        method: "get",
        params
    })
}

export function addCase(data) {
    return requests({
        url: "/api/v1/case",
        method: "post",
        data
    })
}

export function editCase(data) {
    return requests({
        url: "/api/v1/case",
        method: "put",
        data
    })
}

export function deleteCaseById(id) {
    return requests({
        url: "/api/v1/case/" + id,
        method: "delete"
    })
}