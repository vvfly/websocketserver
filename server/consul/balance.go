package consul

import (
	"fmt"
	"sync"

	"github.com/luckyweiwei/websocketserver/proto"
)

type ringNode struct {
	n *ringNode
	v interface{}
}

type Balance struct {
	lock *sync.Mutex

	hnode *ringNode
}

var (
	BalanceSysParam *Balance = nil
	BalanceUser     *Balance = nil
	BalanceNobility *Balance = nil
	BalanceGift     *Balance = nil
	BalanceShop     *Balance = nil
	BalanceLive     *Balance = nil
)

func (b *Balance) Make(size int) {
	if size < 1 {
		return
	}
	b.hnode = &ringNode{}
	m := b.hnode
	for i := 0; i < size-1; i++ {
		m.n = &ringNode{}
		m = m.n
	}
	m.n = b.hnode
}

// 添加节点数据
func (b *Balance) Add(v interface{}) {
	b.hnode.v = v
	b.hnode = b.hnode.n
}

// 轮询
func (b *Balance) Roll() interface{} {
	b.lock.Lock()
	defer b.lock.Unlock()

	v := b.hnode.v
	b.hnode = b.hnode.n
	return v
}

func BalanceInit() {
	RegisterBalance(proto.SysParamServerName)
	RegisterBalance(proto.UserServerName)
	RegisterBalance(proto.NobilityServerName)
	RegisterBalance(proto.GiftServerName)
	RegisterBalance(proto.ShopServerName)
	RegisterBalance(proto.LiveServerName)
}

func RegisterBalance(name string) {
	balance := &Balance{
		lock: new(sync.Mutex),
	}

	if name == proto.SysParamServerName {
		BalanceSysParam = balance
	} else if name == proto.UserServerName {
		BalanceUser = balance
	} else if name == proto.NobilityServerName {
		BalanceNobility = balance
	} else if name == proto.GiftServerName {
		BalanceGift = balance
	} else if name == proto.ShopServerName {
		BalanceShop = balance
	} else if name == proto.LiveServerName {
		BalanceLive = balance
	} else {
		return
	}

	srvs := GetConsulSrv(name)
	balance.Make(len(srvs))
	for _, v := range srvs {
		balance.Add(v)
	}

}

func GetBalanceAddr(name string) string {
	var balance *Balance

	if name == proto.SysParamServerName {
		balance = BalanceSysParam
	} else if name == proto.UserServerName {
		balance = BalanceUser
	} else if name == proto.NobilityServerName {
		balance = BalanceNobility
	} else if name == proto.GiftServerName {
		balance = BalanceGift
	} else if name == proto.ShopServerName {
		balance = BalanceShop
	} else if name == proto.LiveServerName {
		balance = BalanceLive
	} else {
		return ""
	}

	b := balance.Roll().(*ConsulManager)
	return fmt.Sprintf("http://%s:%d", b.Addr, b.Port)
}
