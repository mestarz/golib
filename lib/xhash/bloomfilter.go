package xhash

type SimpleBloomFilter struct {
	// 256个bit的bloomFilter
	bits  [4]uint64
	HFunc func(interface{}, ...interface{}) []byte
}

func NewSimpleBloomFilter(f func(interface{}, ...interface{}) []byte) *SimpleBloomFilter {
	return &SimpleBloomFilter{HFunc: f}
}

func (bf *SimpleBloomFilter) onBit(index int) {
	block := (index % 256) / 64
	offset := index % 64
	bf.bits[block] = bf.bits[block] | 1<<offset
}

func (bf *SimpleBloomFilter) offBit(index int) {
	block := (index % 256) / 64
	offset := index % 64
	bf.bits[block] = bf.bits[block] & (0 << offset & uint64((1<<64)-1))
}

func (bf *SimpleBloomFilter) Insert(value interface{}, salt ...interface{}) {
	vb := bf.HFunc(value, salt)
	for _, b := range vb {
		bf.onBit(int(b))
	}
}

func (bf *SimpleBloomFilter) Find(value interface{}, salt ...interface{}) bool {
	vb := bf.HFunc(value, salt)
	for _, b := range vb {
		index := int(b)
		block := (index % 256) / 64
		offset := index % 64
		if bf.bits[block]&(1<<offset) == 0 {
			return false
		}
	}
	return true
}


// 布隆过滤器
type BloomFilter struct {
	// 256个累加器，最大值 1 << 64 - 1
	bits  []uint64
	Len   int
	HFunc func(interface{}, ...interface{}) []int
}

func NewBloomFilter(len int, f func(interface{}, ...interface{}) []int) *BloomFilter {
	bfx := BloomFilter{Len: len, HFunc: f}
	bfx.bits = make([]uint64, len)
	return &bfx
}

func (bfx *BloomFilter) mapIndex(i int) int {
	return i % bfx.Len
}

func (bfx *BloomFilter) Insert(value interface{}, salt ...interface{}) {
	vb := bfx.HFunc(value, salt)
	for _, b := range vb {
		bfx.bits[bfx.mapIndex(b)]++
	}
}

func (bfx *BloomFilter) Find(value interface{}, salt ...interface{}) bool {
	vb := bfx.HFunc(value, salt)
	for _, b := range vb {
		if bfx.bits[bfx.mapIndex(b)] == 0 {
			return false
		}
	}
	return true
}

func (bfx *BloomFilter) Number(value interface{}, salt ...interface{}) uint64 {
	num := uint64((1 << 64) - 1)
	vb := bfx.HFunc(value, salt)
	for _, b := range vb {
		if bfx.bits[bfx.mapIndex(b)] < num {
			num = bfx.bits[bfx.mapIndex(b)]
		}
		if num == 0 {
			return 0
		}
	}
	return num
}

func (bfx *BloomFilter) Remove(value interface{}, salt ...interface{}) {
	vb := bfx.HFunc(value, salt)
	// 先判断是否存在
	for _, b := range vb {
		if bfx.bits[bfx.mapIndex(b)] == 0 {
			return
		}
	}
	// 然后减去
	for _, b := range vb {
		bfx.bits[bfx.mapIndex(b)]--
	}
}

// 将返回的byte按batch组合起来
func HFuncByBatch(batch int, f func(interface{}, ...interface{}) []byte) (int, func(interface{}, ...interface{}) []int) {
	return (1 << (8 * batch)) - 1, func(value interface{}, salt ...interface{}) []int {
		var result []int
		bytes := f(value, salt)
		for i := batch - 1; i < len(bytes); i += batch {
			con_data := 0
			for j := 0; j < batch; j++ {
				con_data = (con_data << 8) + int(bytes[i-j])
			}
			result = append(result, con_data)
		}
		return result
	}
}

// 使用MD5构造的布隆过滤器
// xhash.NewBloomFilterEX(xhash.HFuncByBatch(3, xhash.GenMD5))
func NewBloomFilterByMD5Level(level int) *BloomFilter {
	return NewBloomFilter(HFuncByBatch(level, GenMD5))
}

// 使用SHA1构造的布隆过滤器
// xhash.NewBloomFilterEX(xhash.HFuncByBatch(3, xhash.GenMD5))
func NewBloomFilterBySHA1Level(level int) *BloomFilter {
	return NewBloomFilter(HFuncByBatch(level, GenSHA1))
}
