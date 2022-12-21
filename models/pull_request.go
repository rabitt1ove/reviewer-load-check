package models

import (
    "fmt"
    "reviewer-load-check/constants"
)

// レスポンスデータ：PR情報
type ResponsePullRequests struct {
    Data struct {
        Search struct {
            Nodes []struct {
                Additions int `json:"additions"`
                Deletions int `json:"deletions"`
            } `json:"nodes"`
            PageInfo struct {
                HasNextPage bool   `json:"hasNextPage"`
                EndCursor   string `json:"endCursor"`
            } `json:"pageInfo"`
        } `json:"search"`
    } `json:"data"`
}

// レスポンスデータ：集計したPR情報
type ResponsePullRequest struct {
    PullRequest struct {
        ID           string `json:"id"`
        Count        int    `json:"count"`
        SumAdditions int    `json:"sumAdditions"`
        SumDeletions int    `json:"sumDeletions"`
        State        string `json:"state"`
    } `json:"pullRequest"`
}

// PR情報を検索するクエリー
func (requestParam RequestParam) GetPullRequestsQuery(githubID string, state string, from string, to string) string {
    var reviewState string

    switch state {
    case constants.REVIEW_APPROVED:
        reviewState = "review:approved"
    case constants.REVIEW_NONE:
        reviewState = "review:none"
    default:
        reviewState = ""
    }

    return fmt.Sprintf(`
    {
        search(type: ISSUE, last: 100, query: "is:pr %s reviewed-by:%s created:%s..%s") {
            nodes {
                ... on PullRequest {
                    additions
                    deletions
                }
            }
            pageInfo {
                hasNextPage
                endCursor
            }
        }
    }
    `, reviewState, githubID, from, to)
}

// PR情報を集計する
func TabulatePullRequest(responsePullRequests *ResponsePullRequests, githubId string, state string) ResponsePullRequest {
    var sumAdditions int = 0
    var sumDeletions int = 0

    pullRequestCount := len(responsePullRequests.Data.Search.Nodes)
    for nodesIndex := 0; nodesIndex < pullRequestCount; nodesIndex++ {
        additions := responsePullRequests.Data.Search.Nodes[nodesIndex].Additions
        deletions := responsePullRequests.Data.Search.Nodes[nodesIndex].Deletions
        sumAdditions += additions
        sumDeletions += deletions
    }

    responsePullRequest := ResponsePullRequest{
        PullRequest: struct {
            ID           string `json:"id"`
            Count        int    `json:"count"`
            SumAdditions int    `json:"sumAdditions"`
            SumDeletions int    `json:"sumDeletions"`
            State        string `json:"state"`
        }{
            ID:           githubId,
            Count:        pullRequestCount,
            SumAdditions: sumAdditions,
            SumDeletions: sumDeletions,
            State:        state,
        },
    }

    return responsePullRequest
}
