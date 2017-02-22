package MaxHeap

var capacity, size int32
var data []int32

// 初始化容器数组
func Init(cap int32) {
	data = make([]int32, cap, cap*2)
	size = 0
}

// 获取数组当前元素个数
func GetSize() int32 {
	return size
}

// 获取数据
func GetData() []int32 {
	return data
}

// 插入元素
func Insert(item int32) {
	data[size+1] = item
	size++
	shiftUp(size)
}

// 推出元素，只能是第一个元素
func Extract() (int32, int32) {
	if size < 1 {
		return 0, size
	}
	item := data[1]
	data[1] = data[size]
	size--
	shiftDown()
	return item, size
}

// 添加元素时，添加至数组末尾，然后调用shiftUp调整元素到合适的位置
func shiftUp(k int32) {
	for k > 1 && data[k] > data[k/2] {
		tmp := data[k]
		data[k] = data[k/2]
		data[k/2] = tmp
		k /= 2
	}
}

// 弹出顶部元素时，需要把最后一个元素换至堆顶，然后调用shiftDown调整到合适位置
func shiftDown() {
	var k int32 = 1
	for 2*k <= size {
		j := 2 * k
		if j+1 <= size && data[j+1] > data[j] {
			j += 1
		}

		if data[k] >= data[j] {
			break
		}

		// 不使用临时变量交换slice的值
		data[k], data[j] = data[j], data[k]
		k = j
	}
}
