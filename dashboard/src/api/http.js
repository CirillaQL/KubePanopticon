import axios from "axios";

const url = import.meta.env.VITE_URL

const http = axios.create({
    baseURL: url,
    timeout: 6000
})

export default http