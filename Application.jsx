const React = require("react");
const ReactDOM = require("react-dom/client");

function Application() {
    return <div id = "EPIC">Epic Application</div>
}

const root = ReactDOM.createRoot(document.querySelector("#Application"));
root.render(<Application/>)