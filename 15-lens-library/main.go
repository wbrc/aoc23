package main

import (
	"bufio"
	"container/list"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func hash(s []byte) int {
	h := 0
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}

	return h
}

type box struct {
	elems map[string]*list.Element
	l     *list.List
}

type entry struct {
	key   string
	value int
}

func process(s []byte, boxes [256]*box) {
	if s[len(s)-1] == '-' {
		key := string(s[:len(s)-1])
		box := boxes[hash(s[:len(s)-1])]
		if elem, ok := box.elems[key]; ok {
			box.l.Remove(elem)
			delete(box.elems, key)
			return
		}
	} else {
		key := strings.Split(string(s), "=")[0]
		num := strings.Split(string(s), "=")[1]
		val, _ := strconv.Atoi(num)
		box := boxes[hash([]byte(key))]
		if elem, ok := box.elems[key]; ok {
			elem.Value.(*entry).value = val
		} else {
			box.elems[key] = box.l.PushBack(&entry{key, val})
		}
	}
}

func main() {
	boxes := [256]*box{}
	for i := range boxes {
		boxes[i] = &box{
			elems: map[string]*list.Element{},
			l:     list.New(),
		}
	}

	sl := []byte{}

	for r := bufio.NewReader(os.Stdin); ; {
		b, err := r.ReadByte()
		if errors.Is(err, io.EOF) {
			if len(sl) > 0 {
				process(sl, boxes)
			}
			break
		}

		if b == '\n' {
			continue
		}

		if b == ',' {
			process(sl, boxes)
			sl = sl[:0]
			continue
		}

		sl = append(sl, b)
	}

	sum := 0
	for boxNum, box := range boxes {
		itemNr := 0
		for e := box.l.Front(); e != nil; e, itemNr = e.Next(), itemNr+1 {
			sum += e.Value.(*entry).value * (boxNum + 1) * (itemNr + 1)
		}
	}

	println(sum)
}
