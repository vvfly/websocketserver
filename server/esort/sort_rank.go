package esort

import "github.com/luckyweiwei/websocketserver/model/do"

/*
排序规则：
贵族类型(10000权重) + 守护类型(1000权重) + 经验值 的结果从高到低进行排序
*/
func Compare(i, j do.OnlineUser) bool {
	var (
		iGuardAndNobilityWeight = 0
		jGuardAndNobilityWeight = 0
		iNobilityType           = 0
		jNobilityType           = 0
		iGuardType              = 0
		jGuardType              = 0
	)

	//周时开通守护和贵族的权重最高
	if i.IsNobility() && i.IsGuard() {
		iGuardAndNobilityWeight = 100000
	}
	if j.IsNobility() && j.IsGuard() {
		jGuardAndNobilityWeight = 100000
	}

	// 贵族
	if i.IsNobility() {
		iNobilityType = 10000
	}
	if j.IsNobility() {
		jNobilityType = 10000
	}

	// 守护
	if i.IsGuard() {
		iGuardType = 1000
	}
	if j.IsGuard() {
		iGuardType = 1000
	}

	iWeight := iGuardAndNobilityWeight + iNobilityType + iGuardType + i.ExpGrade
	jWeight := jGuardAndNobilityWeight + jNobilityType + jGuardType + j.ExpGrade

	return jWeight >= iWeight
}

func SortRank(rank []do.OnlineUser, count int) []do.OnlineUser {
	lenRank := len(rank)
	if count <= 0 || lenRank <= count {
		return rank
	}

	newRank := make([]do.OnlineUser, 0)

	for i := 0; i <= lenRank-1; i++ {
		for j := i; j < lenRank-1; j++ {
			if Compare(rank[i], rank[j]) {
				t := rank[i]
				rank[i] = rank[j]
				rank[j] = t
			}
		}

		newRank = append(newRank, rank[i])
		if i >= count-1 {
			break
		}
	}

	return newRank
}
