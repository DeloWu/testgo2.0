import requests from '@/utils/requests'

export function getSuiteTotal() {
    return requests({
        url: "/api/v1/suite-total",
        method: "get"
    })
}

export function getSuiteById(id) {
    return requests({
        url: "/api/v1/suite/" + id,
        method: "get"
    })
}

export function getSuitesByPagination(params) {
    return requests({
        url: "/api/v1/suites",
        method: "get",
        params
    })
}

export function addSuite(data) {
    return requests({
        url: "/api/v1/suite",
        method: "post",
        data
    })
}

export function editSuite(data) {
    return requests({
        url: "/api/v1/suite",
        method: "put",
        data
    })
}

export function deleteSuiteById(id) {
    return requests({
        url: "/api/v1/suite/" + id,
        method: "delete"
    })
}