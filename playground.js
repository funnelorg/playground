function init() {
    var snippets = document.querySelectorAll(".snippet");
    for (kk = 0; kk < snippets.length; kk ++) {
        var code = snippets[kk].querySelector(".code");
        var button = snippets[kk].querySelector("button");
        var output = snippets[kk].querySelector(".output");
        register(code, button, output)
    }
}

function register(code, button, output) {
    button.addEventListener("click", function() {
        var text
        if (code.tagName == "TEXTAREA") {
            text = code.value
        } else {
            text = code.innerText
        }
        output.textContent = JSON.stringify(Funnel.Eval(text));
    });
}

init();
