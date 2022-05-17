function upload() {
    const formData = new FormData();
    const fs = document.getElementById("uploadfile")

    for (let i = 0; i < fs.files.length; i++) {
        formData.append(`files${i}`, fs.files[i])
    }

    fetch('localhost:8080/', {
        method: 'POST',
        body: formData,
    })
        .then(response => response.json())
        .then(result => {
            console.log('Success:', result);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}