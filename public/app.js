import API from "./services/API.js";
import { Test } from "./components/Test.js";

window.app = {
    search: (event) => {
        event.preventDefault();
        const query = document.querySelector("input[type=search]").value;
    },
    api: API
}

window.addEventListener("DOMContentLoaded", () => {
    document.querySelector("main").appendChild(new Test)
})