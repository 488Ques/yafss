function dragOverHandler(ev) {
    ev.preventDefault();
}

var fileInput = document.getElementById("fileInput")

function dropHandler(ev) {
    ev.preventDefault();

    var dt = new DataTransfer();
    for (var i = 0; i < ev.dataTransfer.files.length; i++) {
        dt.items.add(ev.dataTransfer.files[i]);
    }

    fileInput.files = dt.files;
}  