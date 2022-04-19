let countAuthorId = 0
let countKeywordId = 0

let authors = []
let keywords = []

function uploadPaper() {
    let content = document.getElementById("upload-content")
    let spinner = document.getElementById("upload-spinner")
    let title = document.getElementById("paper-title").value
    let abstract = document.getElementById("paper-abstract").value
    let file = document.getElementById("paper-file").files[0]
    if (title === "") {
        mdui.snackbar('Please enter title')
    } else if (abstract === "") {
        mdui.snackbar("Please enter abstract")
    } else if (file === undefined) {
        mdui.snackbar("Please select the paper file")
    } else if (authors.length === 0) {
        mdui.snackbar("Please add at least one author")
    } else if (keywords.length === 0) {
        mdui.snackbar("Please add at least one keyword")
    } else {
        content.hidden = true
        spinner.hidden = false
        let formData = new FormData()
        formData.append("title", title)
        formData.append("abstract", abstract)
        formData.append("authors", JSON.stringify(authors))
        formData.append("keywords", JSON.stringify(keywords))
        formData.append("file", file)
        let succeed = (resp) => {
            let respObj = JSON.parse(resp)
            if (respObj.result === true) {
                content.hidden = false
                spinner.hidden = true
                mdui.snackbar("Paper has been uploaded successfully")
            } else {
                content.hidden = false
                spinner.hidden = true
                if (respObj.error === "" || respObj.error === undefined) {
                    mdui.snackbar(`Failed to upload paper`)
                } else {
                    mdui.snackbar(`Failed to upload paper: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            content.hidden = false
            spinner.hidden = true
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/papers", undefined, formData, succeed, failed).post()
    }
}

function addAuthor() {
    let id = countAuthorId
    let author = document.getElementById("paper-author").value
    if (author !== "") {
        document.getElementById("paper-authors").innerHTML += `<div class="mdui-chip mdui-m-r-1" id="author-${id}"><span class="mdui-chip-title">${author}</span><span class="mdui-chip-delete" onclick="removeAuthor(${id})"><span class="mdui-icon material-icons">cancel</span></span></div>`
        document.getElementById("paper-author").value = ""
        authors[id] = author
        countAuthorId++
    }
}

function removeAuthor(id) {
    document.getElementById(`author-${id}`).remove()
    authors.splice(id, 1)
}

function addKeyword() {
    let id = countKeywordId
    let keyword = document.getElementById("paper-keyword").value
    if (keyword !== "") {
        document.getElementById("paper-keywords").innerHTML += `<div class="mdui-chip mdui-m-r-1" id="keyword-${id}"><span class="mdui-chip-title">${keyword}</span><span class="mdui-chip-delete" onclick="removeKeyword(${id})"><span class="mdui-icon material-icons">cancel</span></span></div>`
        document.getElementById("paper-keyword").value = ""
        keywords[id] = keyword
        countKeywordId++
    }
}

function removeKeyword(id) {
    document.getElementById(`keyword-${id}`).remove()
    keywords.splice(id, 1)
}

function updateFilename() {
    let file = document.getElementById("paper-file").files[0]
    let btn = document.getElementById("upload-file-btn")
    let chip = document.getElementById("uploaded-file")
    let filename = document.getElementById("uploaded-filename")
    btn.setAttribute("hidden", "")
    chip.removeAttribute("hidden")
    filename.innerText = file.name
}

function removeFile() {
    document.getElementById("paper-file").value = ""
    document.getElementById("upload-file-btn").removeAttribute("hidden")
    document.getElementById("uploaded-file").setAttribute("hidden", "")
    document.getElementById("uploaded-filename").innerText = ""
}
