function login() {
    let content = document.getElementById("login-content")
    let spinner = document.getElementById("login-spinner")
    let loginEmail = document.getElementById("login-email").value
    let loginPassword = document.getElementById("login-password").value
    if (loginEmail === "") {
        mdui.snackbar("Please enter email")
    } else if (loginPassword === "") {
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
                    mdui.snackbar(`Failed to login: wrong email or password`)
                } else {
                    mdui.snackbar(`Failed to login: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            content.hidden = false
            spinner.hidden = true
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/users/login", "application/json", JSON.stringify({
            "email": loginEmail,
            "password": md5(loginPassword)
        }), succeed, failed).post()
    }
}

function logout() {
    let content = document.getElementById("logout-content")
    let spinner = document.getElementById("logout-spinner")
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
                mdui.snackbar(`Failed to logout`)
            } else {
                mdui.snackbar(`Failed to login: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        content.hidden = false
        spinner.hidden = true
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    new AJAXRequest("/users/logout", "application/json", JSON.stringify({}), succeed, failed).post()
}

function register() {
    let content = document.getElementById("register-content")
    let spinner = document.getElementById("register-spinner")
    let resisterEmail = document.getElementById("register-email").value
    let registerUsername = document.getElementById("register-username").value
    let registerPassword = document.getElementById("register-password").value
    if (resisterEmail === "") {
        mdui.snackbar("Please enter email")
    } else if (registerUsername === "") {
        mdui.snackbar("Please enter username")
    } else if (registerPassword === "") {
        mdui.snackbar("Please enter password")
    } else {
        content.hidden = true
        spinner.hidden = false
        let succeed = (resp) => {
            let respObj = JSON.parse(resp)
            if (respObj.result === true) {
                content.hidden = false
                spinner.hidden = true
                mdui.snackbar(`User registered successfully`)
            } else {
                content.hidden = false
                spinner.hidden = true
                if (respObj.error === "" || respObj.error === undefined) {
                    mdui.snackbar(`Failed to register`)
                } else {
                    mdui.snackbar(`Failed to register: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            content.hidden = false
            spinner.hidden = true
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/users/register", "application/json", JSON.stringify({
            "email": resisterEmail,
            "username": registerUsername,
            "password": md5(registerPassword)
        }), succeed, failed).post()
    }
}

function updateUsername() {
    let content = document.getElementById("update-username-content")
    let spinner = document.getElementById("update-username-spinner")
    let username = document.getElementById("update-username-username").value
    let password = document.getElementById("update-username-password").value
    if (username === "") {
        mdui.snackbar("Please enter new username")
    } else if (password === "") {
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
                    mdui.snackbar(`Failed to update username: wrong password`)
                } else {
                    mdui.snackbar(`Failed to update username: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            content.hidden = false
            spinner.hidden = true
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/users/update", "application/json", JSON.stringify({
            "field": "username",
            "value": username,
            "password": md5(password)
        }), succeed, failed).post()
    }
}

function updatePassword() {
    let content = document.getElementById("update-password-content")
    let spinner = document.getElementById("update-password-spinner")
    let oldPassword = document.getElementById("update-password-old").value
    let newPassword = document.getElementById("update-password-new").value
    if (oldPassword === "") {
        mdui.snackbar("Please enter password")
    } else if (newPassword === "") {
        mdui.snackbar("Please enter new password")
    } else {
        content.hidden = true
        spinner.hidden = false
        let succeed = (resp) => {
            let respObj = JSON.parse(resp)
            if (respObj.result === true) {
                content.hidden = false
                spinner.hidden = true
                mdui.snackbar(`Password has been updated successfully`)
            } else {
                content.hidden = false
                spinner.hidden = true
                if (respObj.error === "" || respObj.error === undefined) {
                    mdui.snackbar(`Failed to update password: wrong password`)
                } else {
                    mdui.snackbar(`Failed to update password: ${respObj.error}`)
                }
            }
        }
        let failed = (code, text) => {
            content.hidden = false
            spinner.hidden = true
            mdui.snackbar(`Received status ${code}: ${text}`)
        }
        new AJAXRequest("/users/update", "application/json", JSON.stringify({
            "field": "password",
            "value": md5(newPassword),
            "password": md5(oldPassword)
        }), succeed, failed).post()
    }
}

function search() {
    let keyword = document.getElementById("search-keyword").value
    if (keyword !== "") {
        window.location.href = `/papers?t=${keyword}`
    }
}
