package models

import (
	"errors"
	"strconv"
	"strings"
	"sync"
)

var ShopName string = "Тестовый магазин"

// объекты и методы для хранения паролей
type Shop struct {
	mx     sync.RWMutex
	goods  Goods
	orders Orders
}

func NewShop() *Shop {
	return &Shop{goods: Goods{list: make(map[uint16]Good, 100), entryID: 0}}
}

func (s *Shop) AddGood(name, unit string, price uint64) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if name == "" || unit == "" {
		return errors.New("incorrect input data")
	}

	if price == 0 {
		return errors.New("incorrect price")
	}

	s.goods.list[s.goods.entryID] = Good{
		name:  name,
		price: price,
		unit:  unit,
	}

	s.goods.entryID = s.goods.entryID + 1

	return nil
}

func (s *Shop) DeleteGood(goodID uint16) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	_, ok := s.goods.list[goodID]
	if ok {
		delete(s.goods.list, goodID)
		s.goods.entryID = goodID
		return nil
	}

	return errors.New("there is no such a good in the db")
}

func (s *Shop) GetGoods() (goods string, err error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	for key, value := range s.goods.list {
		goods = strings.Join([]string{goods, "<p>", value.name, ", ", value.unit," <a href=/order?order=", strconv.Itoa(int(key)), ">Купить</a></p><br>"}, "")
	}
	return
}

type Good struct {
	name  string
	price uint64
	unit  string
}

type Goods struct {
	list    map[uint16]Good
	entryID uint16
}

type Client struct {
	ID    uint16
	Name  string
	Email string
}

type Order struct {
	ClientID uint16
	Basket   map[uint16]Good
}

type Orders struct {
	List []Order
}
