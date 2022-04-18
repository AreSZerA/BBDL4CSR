let paperId

function setPaperId(id) {
    paperId = id
}

function peerReview() {
    let paperStatus
    let ratios = document.getElementsByName("paper-status")
    for (const s of ratios) {
        if (s.checked) {
            paperStatus = s.value
            break
        }
    }
    console.log(paperStatus)
    let comment = document.getElementById("peer-review-comment").value
    let succeed = (resp) => {
        let respObj = JSON.parse(resp)
        if (respObj.result === true) {
            setTimeout( () => window.location.reload(), 1000)
        } else {
            if (respObj.error === "" || respObj.error === undefined) {
                mdui.snackbar(`Failed to peer review`)
            } else {
                mdui.snackbar(`Failed to register: ${respObj.error}`)
            }
        }
    }
    let failed = (code, text) => {
        mdui.snackbar(`Received status ${code}: ${text}`)
    }
    if (comment !== "") {
        new AJAXRequest("/papers/peer_review", "application/json", JSON.stringify({
            "id": paperId,
            "status": paperStatus,
            "comment": comment
        }), succeed, failed).post()
    }
}