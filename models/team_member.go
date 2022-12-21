package models

import (
    "fmt"
)

// レスポンスデータ：チームメンバー情報
type ResponseTeamMembers struct {
    Data struct {
        Organization struct {
            Team struct {
                Members struct {
                    Nodes []struct {
                        Login string `json:"login"`
                    } `json:"nodes"`
                    PageInfo struct {
                        HasNextPage bool   `json:"hasNextPage"`
                        EndCursor   string `json:"endCursor"`
                    } `json:"pageInfo"`
                } `json:"members"`
            } `json:"team"`
        } `json:"organization"`
    } `json:"data"`
}

// メンバー情報を検索するクエリー
func (requestParam RequestParam) GetTeamMembersQuery(teamName string) string {
    return fmt.Sprintf(`
    {
        organization(login: "Fablic") {
            team(slug: "%s") {
                members(first: 100) {
                    nodes {
                        login
                    }
                    pageInfo {
                        hasNextPage
                        endCursor
                    }
                }
            }
        }
    }
    `, teamName)
}
