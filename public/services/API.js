
const API = {
    baseURL: "/api/v1",
    fetch: async (serviceUrl, _args) => {
        try {
            const res = await fetch(`${API.baseURL}/${serviceUrl}`)
            const data = await res.json()
            return data
        } catch (e) {
            console.error(e)
        }

    },
    getCases: async () => {
        return API.fetch("cases")
    }
}

export default API;