package main

import (
	"ClaytonUniversityLibrary/blockchain"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func init() {
	_, _ = blockchain.Execute(blockchain.FuncCreateUser, []byte("user@email"), []byte("user_name"), []byte("password"))
	_, _ = blockchain.Execute(blockchain.FuncCreateUser, []byte("reviewer@email"), []byte("reviewer_name"), []byte("password"))
}

func doTest(t *testing.T, fcn string, args ...string) {
	var argsBytes [][]byte
	for _, arg := range args {
		argsBytes = append(argsBytes, []byte(arg))
	}
	resp, err := blockchain.Execute(fcn, argsBytes...)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("Received:", string(resp.Payload))
}

func TestPing(t *testing.T) {
	doTest(t, blockchain.FuncPing)
}

func TestCreateUser(t *testing.T) {
	doTest(t, blockchain.FuncCreateUser, uuid.New().String(), "name", "password")
	doTest(t, blockchain.FuncCreateUser, uuid.New().String(), "name", "password")
	doTest(t, blockchain.FuncCreateUser, uuid.New().String(), "name", "password")
}

func TestCreateReviewer(t *testing.T) {
	doTest(t, blockchain.FuncCreateReviewer, uuid.New().String(), "name", "password")
	doTest(t, blockchain.FuncCreateReviewer, uuid.New().String(), "name", "password")
	doTest(t, blockchain.FuncCreateReviewer, uuid.New().String(), "name", "password")
}

func TestUpdateUserName(t *testing.T) {
	doTest(t, blockchain.FuncUpdateUserName, "user@email", "new_user_name")
}

func TestUpdateUserPasswd(t *testing.T) {
	doTest(t, blockchain.FuncUpdateUserPasswd, "user@email", "new_password")
}

func TestUpdateUserIsReviewer(t *testing.T) {
	doTest(t, blockchain.FuncUpdateUserIsReviewer, "user@email", "true")
	doTest(t, blockchain.FuncUpdateUserIsReviewer, "user@email", "false")
}

func TestUpdateUserIsAdmin(t *testing.T) {
	doTest(t, blockchain.FuncUpdateUserIsAdmin, "user@email", "true")
	doTest(t, blockchain.FuncUpdateUserIsAdmin, "user@email", "false")
}

func TestRetrieveUserByEmail(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveUserByEmail, "user@email")
}

func TestRetrieveAllUsers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveAllUsers)
}

func TestRetrieveAllReviewers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveAllReviewers)
}

func TestRetrieveCountAllUsers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveCountAllUsers)
}

func TestRetrieveCountAllReviewers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveCountAllReviewers)
}

func TestCreatePaper(t *testing.T) {
	doTest(t, blockchain.FuncCreatePaper, "user@email", "title", "abstract", "[]", "[]")
}

func TestRetrieveAllPapers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveAllPapers)
}

func TestRetrieveAcceptedPapers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveAcceptedPapers)
}

func TestRetrieveRejectedPapers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveRejectedPapers)
}

func TestRetrieveReviewingPapers(t *testing.T) {
	doTest(t, blockchain.FuncRetrieveReviewingPapers)
}

func TestRetrievePapersByEmail(t *testing.T) {
	doTest(t, blockchain.FuncRetrievePapersByEmail, "user@email")
}

func TestRetrievePapersByTitle(t *testing.T) {
	doTest(t, blockchain.FuncRetrievePapersByTitle, "title")
}

func TestRetrievePaperById(t *testing.T) {
	doTest(t, blockchain.FuncRetrievePaperById, "52515735b4516da145ee35e25ea6d0e5035978771d012837a184180c360637e7")
}

func TestUpdatePaperStatus(t *testing.T) {
	doTest(t, blockchain.FuncUpdatePaperStatus, "52515735b4516da145ee35e25ea6d0e5035978771d012837a184180c360637e7")
}

func TestUpdatePeerReview(t *testing.T) {
	doTest(t, blockchain.FuncUpdatePeerReview, "52515735b4516da145ee35e25ea6d0e5035978771d012837a184180c360637e7", "c61a0058-08d2-41cb-a182-a5f99e63b243", "accepted", "comment")
}

func TestRetrievePeerReviewsByPaperId(t *testing.T) {
	doTest(t, blockchain.FuncRetrievePeerReviewsByPaperId, "52515735b4516da145ee35e25ea6d0e5035978771d012837a184180c360637e7")
}

func TestRetrievePeerReviewByIds(t *testing.T) {
	doTest(t, blockchain.FuncRetrievePeerReviewByIds, "52515735b4516da145ee35e25ea6d0e5035978771d012837a184180c360637e7", "c61a0058-08d2-41cb-a182-a5f99e63b243")
}
