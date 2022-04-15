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
    new AJAXRequest("/users/login", `email=${email}&password=${password}`, succeed, failed).post()
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
    new AJAXRequest("/users/logout", "", succeed, failed).post()
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
    new AJAXRequest("/users/register", `email=${email}&username=${username}&password=${password}`, succeed, failed).post()
}