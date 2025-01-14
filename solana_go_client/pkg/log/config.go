package log

import "time"

type Config struct {
	Name             string           // 모듈명
	FileLogging      bool             // 파일로깅
	LumberJackConfig LumberJackConfig // 파일로깅설정
	Caller           bool             // 파일위치출력
	PrettyLog        bool             // Pretty로깅
	LogLevel         Level            // 로그레벨
	TimeFormat       string           // 시간표기형식
}

type LumberJackConfig struct {
	Filepath   string        // 파일경로
	Filename   string        // 파일명
	MaxSize    int           // 파일최대크기(MB)
	MaxAge     int           // 파일저장기간설정, 미기입시 삭제 비활성화
	MaxBackups int           // 보관파일개수
	LocalTime  bool          // 로컬시간기준, 기본값 UTC
	Compress   bool          // 압축여부
	Rotate     time.Duration // 로그주기 (초)
}
