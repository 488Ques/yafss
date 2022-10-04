"use strict";

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

    let dt = new DataTransfer();
    for (let i = 0; i < ev.dataTransfer.files.length; i++) {
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

/**
 * Upload one file, grab the file name and its URL on the file server
 * and add it to the DOM
 * @param {File} up 
 */
function upload(up) {
    const fd = new FormData();
    fd.set("upload", up)
    const options = {
        method: "POST",
        body: fd,
    };

    fetch("/upload", options)
        .then(resp => resp.text())
        .then(uri => {
            let linksLst = document.getElementById("linksLst");

            let a = document.createElement("a");
            a.setAttribute("href", uri);
            a.setAttribute("target", "_blank");
            a.textContent = up.name;

            let li = document.createElement("li");
            li.className = "list-group-item";

            li.appendChild(a);
            linksLst.appendChild(li);
        })
}

/**
 * Upload multiple files by calling upload() on each file of input
 */
function uploadMultiple() {
    const input = document.getElementById('upload');
    let ups = input.files;
    for (let i = 0; i < ups.length; i++) {
        upload(ups[i]);
    }

    input.value = null;
}