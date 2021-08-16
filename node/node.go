package node

type Node struct {
	Name            string
	IP              string
	Memory          int
	MemoryAllocated int
	Cores           int
	Disk            int
	DiskAllocated   int
	TaskCount       int
	Role            string
}
