"use strict";

const MiB = 1 << 20;
const HTTP_STATUS_OK = 200;
const HTTP_STATUS_REQUEST_ENTITY_TOO_LARGE = 413;
const config = fetch("/config")
    .then(resp => resp.json())
    .then(data => data);

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
 * Upload one file and return its URL on the file server
 * @param {File} up 
 * @returns {Promise<string>}
 */
async function upload(up) {
    const fd = new FormData();
    fd.set("upload", up)
    const options = {
        method: "POST",
        body: fd,
    };

    // TODO Handle all possible HTTP codes
    const resp = await fetch("/upload", options);
    const httpStatus = resp.status;
    if (httpStatus != HTTP_STATUS_OK) {
        throw httpStatus;
    }
    const url = await resp.text();
    return url;
}

/**
 * Upload multiple files by calling upload() on each file of input,
 * then grab the file name and its URL on the file server and add 
 * it to the DOM
 */
async function uploadMultiple() {
    const input = document.getElementById("upload");
    const uploadLimit = (await config).uploadLimit;

    let ups = input.files;
    // TODO Maybe uploading concurrently?
    for (let i = 0; i < ups.length; i++) {
        let up = ups[i];

        if (up.size / MiB > uploadLimit) { // Convert size from byte to MiB
            displayUploadedURL(null, up.name, "File size exceeds limit!");
            continue;
        }

        let url;
        try {
            url = await upload(up);
        } catch (err) {
            switch (err) {
                case HTTP_STATUS_REQUEST_ENTITY_TOO_LARGE:
                    displayUploadedURL(null, up.name, "File size exceeds limit!");
                    break;
                default:
                    displayUploadedURL(null, up.name, `Unexpected error. HTTP code: ${err}`);
            }
            continue;
        }

        displayUploadedURL(url, up.name, null);
    }

    input.value = null;
}

/**
 * Add file name and its URL (if exists) on file server to the DOM,
 * and show a small info message (if exists)
 * @param {?string} url
 * @param {string} fileName 
 * @param {?string} msg
 */
function displayUploadedURL(url, fileName, msg) {
    const linksLst = document.getElementById("linksList");

    let listItem = document.createElement("li");
    listItem.classList.add("list-group-item");
    listItem.textContent = fileName;

    if (url !== null) {
        let fileURL = document.createElement("a");
        fileURL.setAttribute("href", url);
        fileURL.setAttribute("target", "_blank");
        fileURL.textContent = fileName;

        listItem.textContent = null;
        listItem.appendChild(fileURL);
    }

    if (msg !== null) {
        let alert = document.createElement("div");
        alert.classList.add("text-danger");
        alert.classList.add("fw-bold");
        alert.textContent = msg;

        listItem.appendChild(alert);
    }

    linksLst.appendChild(listItem);
}