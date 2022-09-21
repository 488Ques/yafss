function dragOverHandler(ev) {
    ev.preventDefault();
    ev.target.classList.add("dragging");
}

function dragLeaveHandler(ev) {
    ev.preventDefault();
    ev.target.classList.remove("dragging");
}

function dropHandler(ev) {
    ev.preventDefault();

    const input = document.getElementById('upload');

    var dt = new DataTransfer();
    for (var i = 0; i < ev.dataTransfer.files.length; i++) {
        dt.items.add(ev.dataTransfer.files[i]);
    }

    input.files = dt.files;
    ev.target.classList.remove("dragging");

    uploadMultiple();
}

function browseClick() {
    const input = document.getElementById('upload');
    input.click();
}

function upload(up) {
    const fd = new FormData();
    fd.set("upload", up)
    const options = {
        method: "POST",
        body: fd,
    };

    fetch("/upload", options)
        .then(resp => resp.json())
        .then(data => {
            var linksLst = document.getElementById("linksLst");
            for (var fileName in data) {
                var a = document.createElement("a");
                a.setAttribute("href", data[fileName]);
                a.setAttribute("target", "_blank");
                a.textContent = fileName

                var li = document.createElement("li");
                li.className = "list-group-item";

                li.appendChild(a);
                linksLst.appendChild(li);
            }
        })
}

function uploadMultiple() {
    const input = document.getElementById('upload');
    var ups = input.files;
    for (let i = 0; i < ups.length; i++) {
        upload(ups[i]);
    }

    input.value = null;
}