function login() {
    let email = document.getElementById("login-email").value
    let password = document.getElementById("login-password").value
    let succeed = (resp) => {
        let respObj = JSON.parse(resp)
        if (respObj.result === true) {
            window.location.reload()
        } else {
            if (respObj.error === "" || respObj.error === undefined) {
                mdui.snackbar(`Failed to login: wrong email or password`)
            } else {
                mdui.snackbar(`Failed to login: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    if (email !== "" && password !== "") {
        new AJAXRequest("/users/login", JSON.stringify({
            "email": email,
            "password": md5(password)
        }), succeed, failed).post()
    }
}

function logout() {
    let succeed = (resp) => {
        let respObj = JSON.parse(resp)
        if (respObj.result === true) {
            setTimeout(() => window.location.reload(), 1000)
        } else {
            if (respObj.error === "" || respObj.error === undefined) {
                mdui.snackbar(`Failed to logout`)
            } else {
                mdui.snackbar(`Failed to login: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    new AJAXRequest("/users/logout", JSON.stringify({}), succeed, failed).post()
}

function register() {
    let email = document.getElementById("register-email").value
    let username = document.getElementById("register-username").value
    let password = document.getElementById("register-password").value
    let succeed = (resp) => {
        let respObj = JSON.parse(resp)
        if (respObj.result === true) {
            mdui.snackbar(`User registered successfully`)
        } else {
            if (respObj.error === "" || respObj.error === undefined) {
                mdui.snackbar(`Failed to register`)
            } else {
                mdui.snackbar(`Failed to register: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    if (email !== "" && username !== "" && password !== "") {
        new AJAXRequest("/users/register", JSON.stringify({
            "email": email,
            "username": username,
            "password": md5(password)
        }), succeed, failed).post()
    }
}

function updateUsername() {
    let username = document.getElementById("update-username-username").value
    let password = document.getElementById("update-username-password").value
    let succeed = (resp) => {
        let respObj = JSON.parse(resp)
        if (respObj.result === true) {
            setTimeout(() => window.location.reload(), 1000)
        } else {
            if (respObj.error === "" || respObj.error === undefined) {
                mdui.snackbar(`Failed to update username: wrong password`)
            } else {
                mdui.snackbar(`Failed to update username: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    if (username !== "" && password !== "") {
        new AJAXRequest("/users/update", JSON.stringify({
            "field": "username",
            "value": username,
            "password": md5(password)
        }), succeed, failed).post()
    }
}

function updatePassword() {
    let oldPassword = document.getElementById("update-password-old").value
    let newPassword = document.getElementById("update-password-new").value
    let succeed = (resp) => {
        let respObj = JSON.parse(resp)
        if (respObj.result === true) {
            mdui.snackbar(`Password has been updated successfully`)
        } else {
            if (respObj.error === "" || respObj.error === undefined) {
                mdui.snackbar(`Failed to update password: wrong password`)
            } else {
                mdui.snackbar(`Failed to update password: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    if (oldPassword !== "" && newPassword !== "") {
        new AJAXRequest("/users/update", JSON.stringify({
            "field": "password",
            "value": md5(newPassword),
            "password": md5(oldPassword)
        }), succeed, failed).post()
    }
}
