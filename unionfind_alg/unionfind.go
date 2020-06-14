package unionfindalg

import "errors"

// UnionFind1 :
type UnionFind1 map[int]*ufGroup1

// ufGroup :
type ufGroup1 struct {
	label    int
	groupMap map[int]struct{}
}

// NewUnionFind1 :
func NewUnionFind1() *UnionFind1 {
	res := make(UnionFind1)
	return &res
}

// Union :
func (u *UnionFind1) Union(dst, src int) {
	if u == nil {
		return
	}
	dstGroup, dstExist := (*u)[dst]
	srcGroup, srcExist := (*u)[src]
	if !dstExist && !srcExist {
		(*u)[dst] = newUnionFindGroup1(dst)
		(*u)[dst].groupMap[src] = struct{}{}
		(*u)[src] = (*u)[dst]
	} else if !dstExist && srcExist {
		(*u)[dst] = (*u)[src]
		(*u)[dst].groupMap[dst] = struct{}{}
		u.SetNewGroupLabel(src, dst)
	} else if dstExist && !srcExist {
		(*u)[dst].groupMap[src] = struct{}{}
		(*u)[src] = (*u)[dst]
	} else {
		if dst == src {
			return
		}
		for k := range srcGroup.groupMap {
			dstGroup.groupMap[k] = struct{}{}
			(*u)[k] = dstGroup
		}
	}
}

// IsInSameGroup :
func (u *UnionFind1) IsInSameGroup(a, b int) bool {
	if u == nil {
		return false
	}
	aLabel, e1 := u.Find(a)
	bLabel, e2 := u.Find(b)
	if e1 != nil || e2 != nil {
		return false
	}
	return aLabel == bLabel
}

// Find : Return the label in the group
func (u *UnionFind1) Find(l int) (int, error) {
	if u == nil {
		return 0, errors.New("Receiver is nil")
	}
	group, exist := (*u)[l]
	if !exist {
		return 0, errors.New("Invalid lable")
	}
	return group.label, nil
}

// SetNewGroupLabel :
func (u *UnionFind1) SetNewGroupLabel(old, new int) bool {
	if u == nil {
		return false
	}
	// 查找旧组是否存在
	oldGroup, exist := (*u)[old]
	if oldGroup == nil || !exist {
		return false
	}
	// 新Label必须是组中的一员
	if _, exist := oldGroup.groupMap[new]; !exist {
		return false
	}
	oldGroup.label = new
	return true
}

// newUnionFindGroup1 :
func newUnionFindGroup1(label int) *ufGroup1 {
	res := &ufGroup1{
		label:    label,
		groupMap: make(map[int]struct{}),
	}
	res.groupMap[label] = struct{}{}
	return res
}
