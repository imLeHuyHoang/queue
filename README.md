# Queue

Thư viện Queue cung cấp cấu trúc dữ liệu hàng đợi (FIFO - First In First Out) với kiểu generic cho Go.

## Tổng quan

Queue là một cấu trúc dữ liệu hoạt động theo nguyên tắc **FIFO** (First In First Out - vào trước ra trước). Phần tử được thêm vào đầu tiên sẽ được lấy ra đầu tiên, giống như hàng người xếp hàng chờ đợi - người đến trước sẽ được phục vụ trước.

**Thuật ngữ:**
- **Front/Head (Đầu)**: Vị trí lấy phần tử ra
- **Rear/Tail (Đuôi)**: Vị trí thêm phần tử vào
- **Enqueue**: Thao tác thêm phần tử vào đuôi queue
- **Dequeue**: Thao tác lấy và xóa phần tử ở đầu queue

**Ứng dụng thực tế:**
- Xử lý hàng đợi in ấn (print queue)
- Quản lý yêu cầu trong web server
- Breadth-First Search (BFS) trong thuật toán đồ thị
- Hệ thống xử lý task/job scheduling
- Message queue trong hệ thống phân tán

## Cài đặt

```bash
go get github.com/imLeHuyHoang/queue
```

## Các phương thức

### `NewQueue[T any]() *Queue[T]`
Tạo một queue mới với kiểu dữ liệu tùy chỉnh.

```go
q := queue.NewQueue[int]()        // Queue chứa số nguyên
q2 := queue.NewQueue[string]()    // Queue chứa chuỗi
```

### `Enqueue(v T)`
Thêm một phần tử vào đuôi queue.

```go
q := queue.NewQueue[int]()
q.Enqueue(10)  // Queue: [10]
q.Enqueue(20)  // Queue: [10, 20]
q.Enqueue(30)  // Queue: [10, 20, 30]
//              Front ^            ^ Rear
```

### `Dequeue() (v T, ok bool)`
Lấy và xóa phần tử ở đầu queue. Trả về giá trị và `ok=true` nếu thành công, `ok=false` nếu queue rỗng.

```go
v, ok := q.Dequeue()  // v = 10, ok = true, Queue: [20, 30]
if ok {
    fmt.Println(v)  // Output: 10
}

// Nếu queue rỗng
empty := queue.NewQueue[int]()
v, ok := empty.Dequeue()  // v = 0 (giá trị zero), ok = false
```

### `Front() (v T, ok bool)`
Xem phần tử ở đầu queue mà không xóa nó. Trả về giá trị và `ok=true` nếu thành công, `ok=false` nếu queue rỗng.

```go
v, ok := q.Front()  // v = 20, ok = true, Queue vẫn: [20, 30]
if ok {
    fmt.Println(v)  // Output: 20
}
```

### `Rear() (v T, ok bool)`
Xem phần tử ở đuôi queue mà không xóa nó. Trả về giá trị và `ok=true` nếu thành công, `ok=false` nếu queue rỗng.

```go
v, ok := q.Rear()  // v = 30, ok = true, Queue vẫn: [20, 30]
if ok {
    fmt.Println(v)  // Output: 30
}
```

### `Len() int`
Trả về số lượng phần tử trong queue.

```go
count := q.Len()  // count = 2
```

### `IsEmpty() bool`
Kiểm tra queue có rỗng hay không.

```go
if q.IsEmpty() {
    fmt.Println("Queue rỗng")
} else {
    fmt.Println("Queue có", q.Len(), "phần tử")
}
```

### `Clear()`
Xóa tất cả phần tử trong queue.

```go
q.Clear()
fmt.Println(q.IsEmpty())  // true
```

### `ToSlice() []T`
Trả về slice chứa tất cả phần tử trong queue (từ đầu đến đuôi). Slice được trả về là bản sao, không ảnh hưởng đến queue gốc.

```go
q := queue.NewQueue[int]()
q.Enqueue(10)
q.Enqueue(20)
q.Enqueue(30)

slice := q.ToSlice()  // [10, 20, 30]
```

## Ví dụ đầy đủ

```go
package main

import (
	"fmt"
	"github.com/imLeHuyHoang/queue"
)

func main() {
	// Tạo queue chứa string
	q := queue.NewQueue[string]()
	
	// Kiểm tra queue rỗng
	fmt.Println("Queue rỗng:", q.IsEmpty())  // true
	
	// Thêm phần tử
	q.Enqueue("Alice")
	q.Enqueue("Bob")
	q.Enqueue("Charlie")
	fmt.Println("Số người trong hàng:", q.Len())  // 3
	
	// Xem người đầu tiên
	first, ok := q.Front()
	if ok {
		fmt.Println("Người đầu tiên:", first)  // Alice
	}
	
	// Xem người cuối cùng
	last, ok := q.Rear()
	if ok {
		fmt.Println("Người cuối cùng:", last)  // Charlie
	}
	
	// Phục vụ từng người (FIFO)
	fmt.Println("\nĐang phục vụ:")
	for !q.IsEmpty() {
		person, _ := q.Dequeue()
		fmt.Println("-", person)
	}
	// Output:
	// - Alice
	// - Bob
	// - Charlie
	
	// Thử dequeue khi queue rỗng
	_, ok = q.Dequeue()
	if !ok {
		fmt.Println("\nQueue đã rỗng, không còn ai để phục vụ")
	}
}
```

## Ví dụ thực tế: Print Queue

```go
package main

import (
	"fmt"
	"github.com/imLeHuyHoang/queue"
)

type PrintJob struct {
	ID       int
	Document string
	Pages    int
}

func main() {
	printQueue := queue.NewQueue[PrintJob]()
	
	// Thêm các công việc in
	printQueue.Enqueue(PrintJob{1, "Report.pdf", 10})
	printQueue.Enqueue(PrintJob{2, "Invoice.docx", 2})
	printQueue.Enqueue(PrintJob{3, "Presentation.pptx", 25})
	
	fmt.Println("Số công việc đang chờ:", printQueue.Len())
	
	// Xem công việc tiếp theo
	next, ok := printQueue.Front()
	if ok {
		fmt.Printf("Công việc tiếp theo: #%d - %s (%d trang)\n", 
			next.ID, next.Document, next.Pages)
	}
	
	// Xử lý các công việc in
	fmt.Println("\nĐang in:")
	for !printQueue.IsEmpty() {
		job, _ := printQueue.Dequeue()
		fmt.Printf("In #%d: %s (%d trang)... Hoàn thành!\n", 
			job.ID, job.Document, job.Pages)
	}
}
```

## So sánh Queue và Stack

| Đặc điểm | Queue (FIFO) | Stack (LIFO) |
|----------|--------------|--------------|
| Nguyên tắc | Vào trước - Ra trước | Vào sau - Ra trước |
| Thêm phần tử | Enqueue (vào đuôi) | Push (vào đỉnh) |
| Lấy phần tử | Dequeue (từ đầu) | Pop (từ đỉnh) |
| Ví dụ thực tế | Hàng xếp hàng | Chồng đĩa |
| Ứng dụng | BFS, Task scheduling | DFS, Undo/Redo |

## Độ phức tạp thời gian

| Phương thức | Độ phức tạp | Ghi chú |
|-------------|-------------|---------|
| Enqueue     | O(1)*       | Amortized khi cần mở rộng slice |
| Dequeue     | O(n)        | Do phải dịch chuyển các phần tử |
| Front       | O(1)        | |
| Rear        | O(1)        | |
| Len         | O(1)        | |
| IsEmpty     | O(1)        | |
| Clear       | O(1)        | |
| ToSlice     | O(n)        | Tạo bản sao |

**Lưu ý về hiệu suất:** Implementation hiện tại sử dụng slice, dẫn đến Dequeue có độ phức tạp O(n). Để đạt O(1) cho cả Enqueue và Dequeue, có thể sử dụng circular buffer hoặc linked list. Implementation này tối ưu cho các use case đơn giản và dễ hiểu.

## Cải tiến trong tương lai

- [ ] Circular buffer implementation cho Dequeue O(1)
- [ ] Linked list implementation
- [ ] Bounded queue với capacity giới hạn
- [ ] Priority queue
- [ ] Thread-safe queue với mutex

## License

MIT
