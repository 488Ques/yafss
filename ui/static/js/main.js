var fileInput = document.getElementById("fileInput");
var dropZone = document.getElementById("dropZone");

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

    var dt = new DataTransfer();
    for (var i = 0; i < ev.dataTransfer.files.length; i++) {
        dt.items.add(ev.dataTransfer.files[i]);
    }

    fileInput.files = dt.files;
    ev.target.classList.remove("dragging");
}  