package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	enc = binary.BigEndian // 레코드 크기와 인덱스 항목을 저장할 때의 인코딩 정의
)

const (
	lenWidth = 8 // 레코드 길이를 저장하는 바이트. 여기서 레코드는 로그에 저장한 데이터를 의미함
)

// 파일의 단순한 래퍼
type store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	size := uint64(fi.Size())
	return &store{
		File: f,
		buf:  bufio.NewWriter(f),
		size: size,
	}, nil
}

// Append : 실제로 쓴 바이트 수, 파일의 위치를 리턴한다. 세그먼트는 레코드의 인덱스 항목을 생성할 때 이 위치 정보를 사용.
func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	pos = s.size
	// 로그 버퍼에 담을 때, 바이트 길이를 미리 담아두면 버퍼의 크기를 미리 알 수 있어 리소스 관리 나 네트워크 통신 최적화에 도움이 된다고 함.
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}
	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}
	w += lenWidth
	s.size += uint64(w)
	return uint64(w), pos, nil
}

// Read: pos 위치에 저장된 레코드를 리턴하는 메서드.
func (s *store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 읽으려는 레코드가 아직 버퍼에 있을 때를 대비해서 우선은 쓰기 버퍼에서 Flush.
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}
	size := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}
	b := make([]byte, enc.Uint64(size))
	if _, err := s.File.ReadAt(b, int64(pos+lenWidth)); err != nil {
		return nil, err
	}
	return b, nil
}

// ReadAt: off 오프셋부터 len(p) 바이트만큼 p에 넣어준다.
func (s *store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.buf.Flush(); err != nil {
		return 0, err
	}
	return s.File.ReadAt(p, off)
}

// Close: 파일을 닫기 전 버퍼의 데이터를 파일에 쓴다.
func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.buf.Flush(); err != nil {
		return err
	}
	return s.File.Close()
}
