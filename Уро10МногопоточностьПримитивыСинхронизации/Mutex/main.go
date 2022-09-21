package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	//mu  sync.Mutex //желательно ЧТОБ МЮТЕКС ДОЛЖЕН БЫТЬ ОБСЯВЛЕН ВЫШЕ ВСЕХ ЗНАЧЕНИИ КОТОРЫЕ ОН ЗАЩИЩАЕТ
	mu  sync.RWMutex
	cou map[string]int
}

func (c *Counter) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cou
}

func (c *Counter) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock() //блокировка и разблокировка только на чтение так как здесь он только читает таким образо пройзводительность повышается
	return c.cou
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()

	c.cou[key]++ //здес куча горутин обращяюся к мапе а она небезопастна поэтому этот кусок данных обрабатывает большое количество горутин
	//критическая секчия
	c.mu.Unlock()
	//НЕНАДО ДОЛГО ДЕРЖАТЬ БЛОК ТОЛЬКО ЕСЛИ ЭТО ТРЕБУЕТ ЭКСКЛЮЗИВНОГО ДОСТУПА

}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() //после того как код критеческой секции выполнется
	//ТАМ ГДЕ РЕТУРН ЛУТШЕ ВСЕГО НАПИСАТЬ DEFER
	return c.cou[key]
}

func main() {
	key := "test"
	c := Counter{cou: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)

	}
	time.Sleep(time.Second * 3)
	fmt.Println(c.Value(key))

}
