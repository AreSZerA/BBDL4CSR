let email

function setReviewer(e) {
    email = e
}

function grantReviewer() {
    let password = document.getElementById("grant-password").value
    if (password === "") {
        mdui.snackbar("Please enter password")
    } else {
        let succeed = (resp) => {
            let respObj = JSON.parse(resp)
            if (respObj.result === true) {
                setTimeout(() => window.location.reload(), 1000)
            } else {
                if (respObj.error === "" || respObj.error === undefined) {
                    mdui.snackbar(`Failed to grant reviewer privilege: wrong password`)
                } else {
                    mdui.snackbar(`Failed to grant reviewer privilege: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/admin", "application/json", JSON.stringify({
            "password": md5(password),
            "email": email,
        }), succeed, failed).post()
    }
}