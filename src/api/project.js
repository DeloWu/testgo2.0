import requests from '@/utils/requests'

export function getProjectTotal() {
    return requests({
        url: "/api/v1/project-total",
        method: "get"
    })
}

export function getProjectById(id) {
    return requests({
        url: "/api/v1/project/" + id,
        method: "get"
    })
}

export function getProjectsByPagination(params) {
    return requests({
        url: "/api/v1/projects",
        method: "get",
        params
    })
}

export function addProject(data) {
    return requests({
        url: "/api/v1/project",
        method: "post",
        data
    })
}

export function editProject(data) {
    return requests({
        url: "/api/v1/project",
        method: "put",
        data
    })
}

export function deleteProjectById(id) {
    return requests({
        url: "/api/v1/project/" + id,
        method: "delete"
    })
}