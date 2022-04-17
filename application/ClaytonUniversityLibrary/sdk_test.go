package main

import (
	"ClaytonUniversityLibrary/blockchain"
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func formatJSON(input []byte) string {
	var output bytes.Buffer
	_ = json.Indent(&output, input, "", "  ")
	return output.String()
}

func testExecute(t *testing.T, fcn string, args ...string) {
	var argsBytes [][]byte
	for _, arg := range args {
		argsBytes = append(argsBytes, []byte(arg))
	}
	resp, err := blockchain.Execute(fcn, argsBytes...)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("Received:", formatJSON(resp.Payload))
}

func testQuery(t *testing.T, fcn string, args ...string) {
	var argsBytes [][]byte
	for _, arg := range args {
		argsBytes = append(argsBytes, []byte(arg))
	}
	resp, err := blockchain.Query(fcn, argsBytes...)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println("Received:", formatJSON(resp.Payload))
}

func TestPing(t *testing.T) {
	testQuery(t, blockchain.FuncPing)
}

// Create

func TestCreateUser(t *testing.T) {
	testExecute(t, blockchain.FuncCreateUser, "fang.hj@dl4csr.org", "Hongjian Fang", "25d55ad283aa400af464c76d713c07ad")
	testExecute(t, blockchain.FuncCreateUser, "han.xy@dl4csr.org", "Xueyu Han", "25d55ad283aa400af464c76d713c07ad")
}

func TestCreateReviewer(t *testing.T) {
	testExecute(t, blockchain.FuncCreateReviewer, "alpha@dl4csr.org", "Alpha", "25d55ad283aa400af464c76d713c07ad")
	testExecute(t, blockchain.FuncCreateReviewer, "beta@dl4csr.org", "Beta", "25d55ad283aa400af464c76d713c07ad")
	testExecute(t, blockchain.FuncCreateReviewer, "gamma@dl4csr.org", "Gamma", "25d55ad283aa400af464c76d713c07ad")
	testExecute(t, blockchain.FuncCreateReviewer, "delta@dl4csr.org", "Delta", "25d55ad283aa400af464c76d713c07ad")
}

func TestCreateAdmin(t *testing.T) {
	testExecute(t, blockchain.FuncCreateAdmin, "areszera@dl4csr.org", "areszera", "25d55ad283aa400af464c76d713c07ad")
}

func TestCreatePaper(t *testing.T) {
	testExecute(t, blockchain.FuncCreatePaper,
		"areszera@dl4csr.org",
		"21st Century Schizoid Man",
		"Cat's foot iron claw. Neuro-surgeons scream for more. At paranoia's poison door. Twenty first century schizoid man. Blood rack barbed wire. Politicians' funeral pyre. Innocents raped with napalm fire. Twenty first century schizoid man. Death seed blind man's greed. Poets' starving children bleed. Nothing he's got he really needs. Twenty first century schizoid man.",
		`["Greg Lake","Robert Fripp","Ian McDonald","Michael Rex Giles","Peter John Sinfield"]`,
		`["Progressive","King Crimson"]`,
	)
	testExecute(t, blockchain.FuncCreatePaper,
		"areszera@dl4csr.org",
		"The Court Of The Crimson King",
		"The rusted chains of prison moons are shattered by the sun. I walk a road, horizons change. The tournament's begun. The purple piper plays his tune. The choir softly sing. Three lullabies in an ancient tongue. For the court of the crimson king. The keeper of the city keys put shutters on the dreams. I wait outside the pilgrim's door with insufficient schemes. The black queen chants the funeral march. The cracked brass bells will ring to summon back the fire witch to the court of the crimson king. The gardener plants an evergreen. Whilst trampling on a flower. I chase the wind of a prism ship to taste the sweet and sour. The pattern juggler lifts his hand. The orchestra begin as slowly turns the grinding wheel in the court of the crimson king. On soft gray mornings widows cry. The wise men share a joke. I run to grasp divining signs to satisfy the hoax. The yellow jester does not play but gently pulls the strings. And smiles as the puppets dance in the court of the crimson king.",
		`["Robert Fripp","Giles Michael Rex"]`,
		`["Progressive","King Crimson"]`,
	)
	testExecute(t, blockchain.FuncCreatePaper,
		"areszera@dl4csr.org",
		"Blockchain-Based Digital Library for Computer Science Research",
		"The paper introduces the implementation of blockchain-based digital library for computer science research using Hyperledger Fabric. Hyper means extraordinary. Ledger means account book for transactions. Hyperledger, extraordinary account book for transactions.",
		`["AreSZerA"]`,
		`["Blockchain","Digital Library","Hyperledger Fabric"]`,
	)
}

// Update

func TestUpdateUserName(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserName, "areszera@dl4csr.org", "AreSZerA")
}

func TestUpdateUserPassword(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserPassword, "areszera@dl4csr.org", "e807f1fcf82d132f9bb018ca6738a19f")
}

func TestUpdateUserIsReviewer(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserIsReviewer, "areszera@dl4csr.org")
}

func TestUpdateUserIsNotReviewer(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserIsNotReviewer, "areszera@dl4csr.org")
}

func TestUpdateUserIsNotAdmin(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserIsNotAdmin, "areszera@dl4csr.org")
}

func TestUpdateUserIsAdmin(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserIsAdmin, "areszera@dl4csr.org")
}

func TestUpdateUserReviewing(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserReviewing, "areszera@dl4csr.org", "10000")
}

func TestUpdateUserByEmail(t *testing.T) {
	testExecute(t, blockchain.FuncUpdateUserByEmail, "areszera@dl4csr.org")
}

func TestUpdatePaperById(t *testing.T) {
	testExecute(t, blockchain.FuncUpdatePaperById, "da253ca163027d21860425cdd773ad42")
	testExecute(t, blockchain.FuncUpdatePaperById, "cb0220bc59dcdcfb0c3c86ac36b427bf")
}

func TestUpdatePeerReviewByPaperAndReviewer(t *testing.T) {
	testExecute(t, blockchain.FuncUpdatePeerReviewByPaperAndReviewer, "da253ca163027d21860425cdd773ad42", "alpha@dl4csr.org", "accepted", "Awesome!")
	time.Sleep(time.Second)
	testExecute(t, blockchain.FuncUpdatePeerReviewByPaperAndReviewer, "da253ca163027d21860425cdd773ad42", "beta@dl4csr.org", "accepted", "Unbelievable!")
	time.Sleep(time.Second)
	testExecute(t, blockchain.FuncUpdatePeerReviewByPaperAndReviewer, "da253ca163027d21860425cdd773ad42", "delta@dl4csr.org", "accepted", "More please.")
	time.Sleep(time.Second)
	testExecute(t, blockchain.FuncUpdatePeerReviewByPaperAndReviewer, "cb0220bc59dcdcfb0c3c86ac36b427bf", "alpha@dl4csr.org", "rejected", "What?")
	time.Sleep(time.Second)
	testExecute(t, blockchain.FuncUpdatePeerReviewByPaperAndReviewer, "cb0220bc59dcdcfb0c3c86ac36b427bf", "gamma@dl4csr.org", "rejected", "Stopping watching Chubbyemu's videos when coding!")
	time.Sleep(time.Second)
	testExecute(t, blockchain.FuncUpdatePeerReviewByPaperAndReviewer, "cb0220bc59dcdcfb0c3c86ac36b427bf", "delta@dl4csr.org", "rejected", "The abstract is so abstract.")
}

// Retrieve

func TestRetrieveUsers(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveUsers)
	testQuery(t, blockchain.FuncRetrieveUsersSortByEmail)
	testQuery(t, blockchain.FuncRetrieveUsersSortByName)
}

func TestRetrieveUsersByName(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveUsersByNameSortByEmail, "AreSZerA")
}

func TestRetrieveReviewers(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveReviewersSortByEmail)
	testQuery(t, blockchain.FuncRetrieveReviewersSortByName)
}

func TestRetrieveAdmins(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveAdminsSortByEmail)
	testQuery(t, blockchain.FuncRetrieveAdminsSortByName)
}

func TestRetrieveReviewersByPaperId(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveReviewersByPaperIdSortByEmail, "da253ca163027d21860425cdd773ad42")
	testQuery(t, blockchain.FuncRetrieveReviewersByPaperIdSortByName, "da253ca163027d21860425cdd773ad42")
}

func TestRetrieveUserByEmail(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveUserByEmail, "areszera@dl4csr.org")
}

func TestRetrievePapers(t *testing.T) {
	testQuery(t, blockchain.FuncRetrievePapers)
	testQuery(t, blockchain.FuncRetrievePapersSortByTitle)
	testQuery(t, blockchain.FuncRetrievePapersSortByUploadTime)
}

func TestRetrieveAcceptedPapers(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersSortByTitle)
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersSortByUploadTime)
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersSortByPublishTime)
}

func TestRetrieveRejectedPapers(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveRejectedPapersSortByTitle)
	testQuery(t, blockchain.FuncRetrieveRejectedPapersSortByUploadTime)
}

func TestRetrieveReviewingPapers(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveReviewingPapersSortByTitle)
	testQuery(t, blockchain.FuncRetrieveReviewingPapersSortByUploadTime)
}

func TestRetrieveAcceptedPapersByTitle(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByTitleSortByTitle, "Schizoid")
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByTitleSortByPublishTime, "Schizoid")
}

func TestRetrieveAcceptedPapersByAuthor(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByAuthorSortByTitle, "Robert Fripp")
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByAuthorSortByPublishTime, "Robert Fripp")
}

func TestRetrieveAcceptedPapersByKeyword(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByKeywordSortByTitle, "Progressive")
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByKeywordSortByPublishTime, "Progressive")
}

func TestRetrieveAcceptedPapersByUploader(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByUploaderSortByTitle, "areszera@dl4csr.org")
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByUploaderSortByUploadTime, "areszera@dl4csr.org")
	testQuery(t, blockchain.FuncRetrieveAcceptedPapersByUploaderSortByPublishTime, "areszera@dl4csr.org")
}

func TestRetrieveRejectedPapersByUploader(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveRejectedPapersByUploaderSortByTitle, "areszera@dl4csr.org")
	testQuery(t, blockchain.FuncRetrieveRejectedPapersByUploaderSortByUploadTime, "areszera@dl4csr.org")
	testQuery(t, blockchain.FuncRetrieveRejectedPapersByUploaderSortByPublishTime, "areszera@dl4csr.org")
}

func TestRetrieveReviewingPapersByUploader(t *testing.T) {
	testQuery(t, blockchain.FuncRetrieveReviewingPapersByUploaderSortByTitle, "areszera@dl4csr.org")
	testQuery(t, blockchain.FuncRetrieveReviewingPapersByUploaderSortByUploadTime, "areszera@dl4csr.org")
}

func TestRetrievePaperById(t *testing.T) {
	testQuery(t, blockchain.FuncRetrievePaperById, "da253ca163027d21860425cdd773ad42")
	testQuery(t, blockchain.FuncRetrievePaperById, "cb0220bc59dcdcfb0c3c86ac36b427bf")
}

func TestRetrievePeerReviewsByReviewer(t *testing.T) {
	testQuery(t, blockchain.FuncRetrievePeerReviewsByReviewerSortByCreateTime, "alpha@dl4csr.org")
}
