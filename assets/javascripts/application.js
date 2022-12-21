Vue.createApp({
    data() {
        return {
            // GitHubのチーム情報
            teamMembers: [],
            // GitHubのPR情報
            pullRequests: [],
            // レビュー状態
            reviewState: -1,
            // レビュー状態一覧
            reviewLabels: [
                { value: -1, label: 'ALL' },
                { value:  0, label: 'Approved' },
                { value:  1, label: 'None' }
            ],
            // チームの初期値（選択したチームはローカルストレージに保存）
            selectedTeam: { id: 'team_a', name: 'チームA' },
            // チーム一覧
            teamOptions: [
                // id:GitHubのチーム名
                // name:画面に表示するチーム名
                { id: 'team_a', name: 'チームA' },
                { id: 'team_b', name: 'チームB' },
            ],
            // 作成開始日
            createdFrom: luxon.DateTime.now().minus({days: 14}).toFormat('yyyy-MM-dd'),
            // 作成終了日
            createdTo: luxon.DateTime.now().toFormat('yyyy-MM-dd'),
            // true：処理中・false：処理中以外
            processing: false,
        }
    },
    // データの変更を監視
    watch: {
        pullRequests: {
            handler: function () {
                if (!this.teamMembers[0]) return
                if (this.pullRequests.length >= this.teamMembers[0].length) this.processing = false
            },
            deep: true,
        }
    },
    // 算出プロパティ
    computed: {
        // レビュー状態一覧を表示する
        labels() {
            return this.reviewLabels.reduce(function (a, b) {
                return Object.assign(a, { [b.value]: b.label })
            }, {})
        },
        // 合計変更数で降順ソートする
        sortedByTotalChanges() {
            return this.pullRequests.sort((a, b) => {
              return (b.count + b.sumAdditions + b.sumDeletions) - (a.count + a.sumAdditions + a.sumDeletions)
            })
        },
        // 本日の日付を取得する
        todayYmd() {
            return luxon.DateTime.now().toFormat('yyyy-MM-dd')
        }
    },
    // インスタンス作成時の処理
    created: function() {
        if (localStorage.hasOwnProperty('selectedTeam')) {
            this.selectedTeam = JSON.parse(localStorage.getItem('selectedTeam'))
        }
    },
    // DOM作成後の処理
    mounted: function() {
        this.onSearch()
    },
    methods: {
        // メンバー情報を取得した後、PR情報を検索する
        async onSearch(){
            if (this.processing) return
            this.processing = true

            this.teamMembers.splice(0)
            this.pullRequests.splice(0)

            this.teamMembers.push(await this.doGetTeamMembers())
            for (const member of this.teamMembers[0]) {
                this.doPostPullRequests(member.login)
            }

            localStorage.setItem('selectedTeam', JSON.stringify(this.selectedTeam))
        },
        // メンバー情報を取得する
        async doGetTeamMembers() {
            let responseData = []

            await axios.get('/v1/member', {
                params: {
                  teamName: this.selectedTeam.id
                }
            })
            .then(response => {
                responseData = response.data
            })
            .catch(err => {
                console.log(err)
            })

            return responseData
        },
        // PR情報を取得する
        doPostPullRequests(githubId) {
            const params = new URLSearchParams()
            params.append('githubId', githubId)
            params.append('state', this.reviewState)
            params.append('createdFrom', this.createdFrom)
            params.append('createdTo', this.createdTo)

            axios.post('/v1/search', params)
            .then(response => {
                this.pullRequests.push(response.data.pullRequest)
            })
            .catch(err => {
                console.log(err)
            })
        },
    }
}).mount('#app')
