let targetEmail

function setReviewer(e) {
    targetEmail = e
}

function grantReviewer() {
    let content = document.getElementById("grant-content")
    let spinner = document.getElementById("grant-spinner")
    let password = document.getElementById("grant-password").value
    if (password === "") {
        mdui.snackbar("Please enter password")
    } else {
        content.hidden = true
        spinner.hidden = false
        let succeed = (resp) => {
            let respObj = JSON.parse(resp)
            if (respObj.result === true) {
                window.location.reload()
            } else {
                content.hidden = false
                spinner.hidden = true
                if (respObj.error === "" || respObj.error === undefined) {
                    mdui.snackbar(`Failed to grant reviewer privilege: wrong password`)
                } else {
                    mdui.snackbar(`Failed to grant reviewer privilege: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            content.hidden = false
            spinner.hidden = true
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/admin", "application/json", JSON.stringify({
            "password": md5(password),
            "email": targetEmail,
        }), succeed, failed).post()
    }
}
