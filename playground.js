// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

function init() {
    var query = new URLSearchParams(window.location.search);

    if (query.has("gist")) {
        var path = new URL(query.get("gist")).pathname.split("/")
        var url = "https://gist.github.com/" + path[1] + "/" + path[2] + ".js"
        var code = "code:import(string " + JSON.stringify(query.get("gist")) + ")"

        console.log(query.get("gist"))
        console.log(url)
        console.log(code)
        
        createIFrame(url)
        document.querySelector(".examples").classList.add("hide")
        document.querySelector(".gist").classList.remove("hide")
        register(
            {innerText: code},
            document.querySelector(".gist button"),
            document.querySelector(".gist .output")
        )
        return
    }
    
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
        Funnel.Eval(text, function(result) {
            output.textContent = JSON.stringify(result)
        })
    });
}

function createIFrame(url) {
    // Create an iframe, append it to this document where specified
    var gistFrame = document.createElement("iframe");
    gistFrame.setAttribute("width", "100%");
    gistFrame.id = "gistFrame";

    document.body.appendChild(gistFrame)
    
    // Create the iframe's document
    var gistFrameHTML = '<html><base target="_parent" /><body onload="parent.adjustIframeSize(document.body.scrollHeight)"><scr' +
        'ipt type="text/javascript" src="' + url + '"></sc'+'ript></body></html>';
			
    // Set iframe's document with a trigger for this document to adjust the height
    var gistFrameDoc = gistFrame.document;
    
    if (gistFrame.contentDocument) {
	gistFrameDoc = gistFrame.contentDocument;
    } else if (gistFrame.contentWindow) {
	gistFrameDoc = gistFrame.contentWindow.document;
    }
    
    gistFrameDoc.open();
    gistFrameDoc.writeln(gistFrameHTML);
    gistFrameDoc.close();		
}

function adjustIframeSize(newHeight) {
    var i = document.getElementById("gistFrame");
    i.style.height = parseInt(newHeight) + "px";
    console.log("size adjusted", newHeight);
}

init();
