let paperId

function setPaperId(id) {
    paperId = id
}

function peerReview() {
    let content = document.getElementById("peer-review-content")
    let spinner = document.getElementById("peer-review-spinner")
    let comment = document.getElementById("peer-review-comment").value
    if (comment === "") {
        mdui.snackbar("Please enter comments")
    } else {
        content.hidden = true
        spinner.hidden = false
        let paperStatus
        let ratios = document.getElementsByName("paper-status")
        for (const s of ratios) {
            if (s.checked) {
                paperStatus = s.value
                break
            }
        }
        let succeed = (resp) => {
            let respObj = JSON.parse(resp)
            if (respObj.result === true) {
                window.location.reload()
            } else {
                content.hidden = false
                spinner.hidden = true
                if (respObj.error === "" || respObj.error === undefined) {
                    mdui.snackbar(`Failed to peer review`)
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
        new AJAXRequest("/papers/peer_review", "application/json", JSON.stringify({
            "id": paperId,
            "status": paperStatus,
            "comment": comment
        }), succeed, failed).post()
    }
}
