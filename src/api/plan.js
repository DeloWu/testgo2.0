import requests from '@/utils/requests'

export function getPlanTotal() {
    return requests({
        url: "/api/v1/plan-total",
        method: "get"
    })
}

export function getPlanById(id) {
    return requests({
        url: "/api/v1/plan/" + id,
        method: "get"
    })
}

export function getPlansByPagination(params) {
    return requests({
        url: "/api/v1/plans",
        method: "get",
        params
    })
}

export function addPlan(data) {
    return requests({
        url: "/api/v1/plan",
        method: "post",
        data
    })
}

export function editPlan(data) {
    return requests({
        url: "/api/v1/plan",
        method: "put",
        data
    })
}

export function deletePlanById(id) {
    return requests({
        url: "/api/v1/plan/" + id,
        method: "delete"
    })
}