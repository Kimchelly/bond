// +build !darwin,!linux,!freebsd,!openbsd,!windows

package process

import (
	"syscall"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/internal/common"
	"github.com/shirou/gopsutil/net"
)

type MemoryMapsStat struct {
	Path         string `json:"path" bson:"path,omitempty"`
	Rss          uint64 `json:"rss" bson:"rss,omitempty"`
	Size         uint64 `json:"size" bson:"size,omitempty"`
	Pss          uint64 `json:"pss" bson:"pss,omitempty"`
	SharedClean  uint64 `json:"sharedClean" bson:"sharedClean,omitempty"`
	SharedDirty  uint64 `json:"sharedDirty" bson:"sharedDirty,omitempty"`
	PrivateClean uint64 `json:"privateClean" bson:"privateClean,omitempty"`
	PrivateDirty uint64 `json:"privateDirty" bson:"privateDirty,omitempty"`
	Referenced   uint64 `json:"referenced" bson:"referenced,omitempty"`
	Anonymous    uint64 `json:"anonymous" bson:"anonymous,omitempty"`
	Swap         uint64 `json:"swap" bson:"swap,omitempty"`
}

type MemoryInfoExStat struct {
}

func Pids() ([]int32, error) {
	return []int32{}, common.ErrNotImplementedError
}

func NewProcess(pid int32) (*Process, error) {
	return nil, common.ErrNotImplementedError
}

func (p *Process) Ppid() (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Name() (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Exe() (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Cmdline() (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) CmdlineSlice() ([]string, error) {
	return []string{}, common.ErrNotImplementedError
}
func (p *Process) CreateTime() (int64, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Cwd() (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Parent() (*Process, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) Status() (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Uids() ([]int32, error) {
	return []int32{}, common.ErrNotImplementedError
}
func (p *Process) Gids() ([]int32, error) {
	return []int32{}, common.ErrNotImplementedError
}
func (p *Process) Terminal() (string, error) {
	return "", common.ErrNotImplementedError
}
func (p *Process) Nice() (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) IOnice() (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Rlimit() ([]RlimitStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) RlimitUsage(_ bool) ([]RlimitStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) IOCounters() (*IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) NumCtxSwitches() (*NumCtxSwitchesStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) NumFDs() (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) NumThreads() (int32, error) {
	return 0, common.ErrNotImplementedError
}
func (p *Process) Threads() (map[int32]*cpu.TimesStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) Times() (*cpu.TimesStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) CPUAffinity() ([]int32, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) MemoryInfo() (*MemoryInfoStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) MemoryInfoEx() (*MemoryInfoExStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) Children() ([]*Process, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) OpenFiles() ([]OpenFilesStat, error) {
	return []OpenFilesStat{}, common.ErrNotImplementedError
}
func (p *Process) Connections() ([]net.ConnectionStat, error) {
	return []net.ConnectionStat{}, common.ErrNotImplementedError
}
func (p *Process) NetIOCounters(pernic bool) ([]net.IOCountersStat, error) {
	return []net.IOCountersStat{}, common.ErrNotImplementedError
}
func (p *Process) IsRunning() (bool, error) {
	return true, common.ErrNotImplementedError
}
func (p *Process) MemoryMaps(grouped bool) (*[]MemoryMapsStat, error) {
	return nil, common.ErrNotImplementedError
}
func (p *Process) SendSignal(sig syscall.Signal) error {
	return common.ErrNotImplementedError
}
func (p *Process) Suspend() error {
	return common.ErrNotImplementedError
}
func (p *Process) Resume() error {
	return common.ErrNotImplementedError
}
func (p *Process) Terminate() error {
	return common.ErrNotImplementedError
}
func (p *Process) Kill() error {
	return common.ErrNotImplementedError
}
func (p *Process) Username() (string, error) {
	return "", common.ErrNotImplementedError
}
