# DL4CSR Chaincode

# APIs

## Create

| Name           | Necessary arguments                                                    |
|:---------------|:-----------------------------------------------------------------------|
| ping           | -                                                                      |
| CreateUser     | email, username, password                                              |
| CreateReviewer | email, username, password                                              |
| CreateAdmin    | email, username, password                                              |
| CreatePaper    | uploader email, title, abstract, authors in JSON, and keywords in JSON |

## Update

| Name                               | Necessary arguments                           |
|:-----------------------------------|:----------------------------------------------|
| UpdateUserByEmail                  | email                                         |
| UpdateUserName                     | email and new username                        |
| UpdateUserPassword                 | email and new password                        |
| UpdateUserIsReviewer               | email                                         |
| UpdateUserIsNotReviewer            | email                                         |
| UpdateUserIsAdmin                  | email                                         |
| UpdateUserIsNotAdmin               | email                                         |
| UpdatePaperById                    | paper ID                                      |
| UpdatePeerReviewByPaperAndReviewer | paper ID, reviewer email, acceptance, comment |

## Retrieve

| Name                                                   | Necessary arguments |
|:-------------------------------------------------------|:--------------------|
| RetrieveUsersByQuery                                   | query               |
| RetrieveUsers                                          | -                   |
| RetrieveUsersSortByEmail                               | -                   |
| RetrieveUsersSortByName                                | -                   |
| RetrieveUsersByNameSortByEmail                         | username            |
| RetrieveReviewersSortByEmail                           | -                   |
| RetrieveReviewersSortByName                            | -                   |
| RetrieveReviewersByPaperIdSortByEmail                  | paper ID            |
| RetrieveReviewersByPaperIdSortByName                   | paper ID            |
| RetrieveAdminsSortByEmail                              | -                   |
| RetrieveAdminsSortByName                               | -                   |
| RetrieveUserByEmail                                    | email               |
| RetrievePapersByQuery                                  | query               |
| RetrievePapers                                         | -                   |
| RetrievePapersSortByTitle                              | -                   |
| RetrievePapersSortByUploadTime                         | -                   |
| RetrieveAcceptedPapersSortByTitle                      | -                   |
| RetrieveAcceptedPapersSortByUploadTime                 | -                   |
| RetrieveAcceptedPapersSortByPublishTime                | -                   |
| RetrieveRejectedPapersSortByTitle                      | -                   |
| RetrieveRejectedPapersSortByUploadTime                 | -                   |
| RetrieveReviewingPapersSortByTitle                     | -                   |
| RetrieveReviewingPapersSortByUploadTime                | -                   |
| RetrieveAcceptedPapersByTitleSortByTitle               | title               |
| RetrieveAcceptedPapersByTitleSortByPublishTime         | title               |
| RetrieveAcceptedPapersByAuthorSortByTitle              | author              |
| RetrieveAcceptedPapersByAuthorSortByPublishTime        | author              |
| RetrieveAcceptedPapersByKeywordSortByTitle             | keyword             |
| RetrieveAcceptedPapersByKeywordSortByPublishTime       | keyword             |
| RetrieveAcceptedPapersByUploaderSortByTitle            | uploader email      |
| RetrieveAcceptedPapersByUploaderSortByUploadTime       | uploader email      |
| RetrieveAcceptedPapersByUploaderSortByPublishTime      | uploader email      |
| RetrieveRejectedPapersByUploaderSortByTitle            | uploader email      |
| RetrieveRejectedPapersByUploaderSortByUploadTime       | uploader email      |
| RetrieveRejectedPapersByUploaderSortByPublishTime      | uploader email      |
| RetrieveReviewingPapersByUploaderSortByTitle           | uploader email      |
| RetrieveReviewingPapersByUploaderSortByUploadTime      | uploader email      |
| RetrievePaperById                                      | paper ID            |
| RetrievePeerReviewsByQuery                             | query               |
| RetrievePeerReviewsByReviewerSortByCreateTime          | -                   |
| RetrieveAcceptedPeerReviewsByReviewerSortByCreateTime  | -                   |
| RetrieveRejectedPeerReviewsByReviewerSortByCreateTime  | -                   |
| RetrieveReviewingPeerReviewsByReviewerSortByCreateTime | -                   |

# Indexes

- paper_publish_time
- paper_title    
- paper_upload_time     
- peer_review_create_time
- peer_review_time
- user_email      
- user_name      
- user_reviewing    
