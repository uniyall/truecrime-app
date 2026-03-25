export class Test extends HTMLElement {
    connectedCallback() {
        const template = document.getElementById("template-test");
        const content = template.content.cloneNode(true);
        this.appendChild(content)
    }
}

customElements.define("app-test", Test);